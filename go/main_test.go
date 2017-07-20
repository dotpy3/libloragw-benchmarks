package main

import (
	"fmt"
	"testing"
)

func BenchmarkUplinks(b *testing.B) {
	prepare()
	if start() != C.LGW_HAL_SUCCESS {
		fmt.Println("Concentrator start unsuccessful")
		b.FailNow()
	}
	b.ResetTimer()
	run(b.N)
	b.StopTimer()
	stop()
}
