package reply

import (
	"sync"
	"time"
)

type bufB struct {
	bufsize      int
	modifiedTime time.Time
	buf          []byte
	sync.RWMutex
}

func (t *bufB) size() int {
	t.RLock()
	defer t.RUnlock()

	return t.bufsize
}

func (t *bufB) isEmpty() bool {
	t.RLock()
	defer t.RUnlock()

	return t.bufsize == 0
}

func (t *bufB) reset() {
	t.Lock()
	defer t.Unlock()

	t.bufsize = 0
}

func (t *bufB) append(data []byte) {
	t.Lock()
	defer t.Unlock()

	if len(t.buf) == 0 {
		t.buf = make([]byte, len(data))
	} else {
		diff := len(t.buf) - t.bufsize - len(data)
		if diff < 0 {
			t.buf = append(t.buf, make([]byte, -diff)...)
		} else {
			t.buf = t.buf[:t.bufsize+len(data)]
		}
	}
	t.bufsize += copy(t.buf[t.bufsize:], data)
	t.modifiedTime = time.Now()
}

func (t *bufB) removeFront(n int) {
	if n <= 0 {
		return
	}

	t.Lock()
	defer t.Unlock()

	if t.bufsize == 0 {
		return
	} else if t.bufsize < n {
		panic("尝试移除的数值大于长度")
	} else if t.bufsize == n {
		t.bufsize = 0
	} else {
		t.bufsize = copy(t.buf, t.buf[n:t.bufsize])
	}

	t.modifiedTime = time.Now()
}

func (t *bufB) removeBack(n int) {
	if n <= 0 {
		return
	}

	t.Lock()
	defer t.Unlock()

	if t.bufsize == 0 {
		return
	} else if t.bufsize < n {
		panic("尝试移除的数值大于长度")
	} else if t.bufsize == n {
		t.bufsize = 0
	} else {
		t.bufsize -= n
	}

	t.modifiedTime = time.Now()
}

func (t *bufB) setModifiedTime() {
	t.Lock()
	defer t.Unlock()

	t.modifiedTime = time.Now()
}

func (t *bufB) getModifiedTime() time.Time {
	t.RLock()
	defer t.RUnlock()

	return t.modifiedTime
}

func (t *bufB) hadModified(mt time.Time) bool {
	t.RLock()
	defer t.RUnlock()

	return !t.modifiedTime.Equal(mt)
}

// // 通常情况下使用getCopyBuf替代
func (t *bufB) getPureBuf() (buf []byte) {
	t.RLock()
	defer t.RUnlock()

	return t.buf[:t.bufsize]
}

func (t *bufB) getCopyBuf() (buf []byte) {
	t.RLock()
	defer t.RUnlock()

	buf = make([]byte, t.bufsize)
	copy(buf, t.buf[:t.bufsize])
	return
}
