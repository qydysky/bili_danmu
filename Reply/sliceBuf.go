package reply

// 线程不安全的[]byte操作
// 需保证append到getSlice的[]byte已被使用之间，对象[]byte未改动
type bufI struct {
	b          []int
	e          []int
	size       int
	buf        []byte
	useDirect  bool //true:直接使用源[]byte，仅适用于连续的append
	useBufPool bool //true：使用内部buf，当getSlice调用时，上次getSlice输出[]byte失效
	//useDirect、useBufPool都为false：每次都返回新创建的[]byte
}

func (t *bufI) reset() {
	t.b = []int{}
	t.e = []int{}
	t.size = 0
}

func (t *bufI) append(b, e int) {
	if len(t.e) > 0 && t.e[len(t.e)-1] == b {
		t.e[len(t.e)-1] = e
	} else {
		t.b = append(t.b, b)
		t.e = append(t.e, e)
	}
	t.size += e - b
}

func (t *bufI) getSlice(buf []byte) []byte {
	if t.useDirect && len(t.b) == 1 {
		return buf[t.b[0]:t.e[0]]
	} else if t.useBufPool {
		if len(t.buf) == 0 {
			t.buf = make([]byte, t.size)
		} else if len(t.buf) < t.size {
			t.buf = append(t.buf, make([]byte, t.size-len(t.buf))...)
		} else if diff := len(t.buf) - t.size; diff > 0 {
			t.buf = t.buf[:t.size+diff/2]
		}
		i := 0
		for k, bi := range t.b {
			i += copy(t.buf[i:], buf[bi:t.e[k]])
		}
		return t.buf[:i]
	} else {
		var b = make([]byte, t.size)
		if len(t.b) == 1 {
			copy(b, buf[t.b[0]:t.e[0]])
		} else {
			i := 0
			for k, bi := range t.b {
				i += copy(b[i:], buf[bi:t.e[k]])
			}
		}
		return b
	}
}
