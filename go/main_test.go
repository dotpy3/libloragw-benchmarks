package main

import "testing"

func BenchmarkUplinks(b *testing.B) {
	prepare()
	b.ResetTimer()
	run(b.N)
}
