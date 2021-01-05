package main

import (
	// "os"
	// "runtime/pprof"
	q "github.com/qydysky/bili_danmu"
)

func main() {
	// f, _ := os.OpenFile("cpu.pprof", os.O_RDWR|os.O_CREATE, 0644)
    // pprof.StartCPUProfile(f)

	q.Demo()

	// pprof.StopCPUProfile()
    // f.Close()
}