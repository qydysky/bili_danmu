package main

import ( // "runtime"
	// "runtime/pprof"
	// "net/http"
	// _ "net/http/pprof"
	// "github.com/skratchdot/open-golang/open"
	q "github.com/qydysky/bili_danmu"
)

func main() {
	// go func() {
	// 	http.ListenAndServe("0.0.0.0:8899", nil)
	// }()
	// defer func(){
	// 	open.Run("http://127.0.0.1:8899/debug/pprof/goroutine?debug=2")
	// 	time.Sleep(time.Duration(3)*time.Second)
	// }()
	// go func() {
	// for {
	// 	View()
	// 	time.Sleep(time.Duration(60) * time.Second)
	// 	{
	// 		debug.FreeOSMemory()
	// 	}
	// }
	// }()
	// f, err := os.OpenFile("cpu.pprof", os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer f.Close()
	// pprof.StartCPUProfile(f)

	q.Start()

	// pprof.StopCPUProfile()
}

// func View() {
// var memStats runtime.MemStats
// runtime.ReadMemStats(&memStats)
// fmt.Printf("=====\n")
// fmt.Printf("总内存:%v MB\n",memStats.Alloc/1024e2/8)
// fmt.Printf("GC次数:%v \n",memStats.NumGC)
// fmt.Printf("堆 :%v %v MB\n",memStats.HeapInuse/1024e2/8,(memStats.HeapIdle - memStats.HeapReleased)/1024e2/8)
// fmt.Printf("=====\n")
// }
