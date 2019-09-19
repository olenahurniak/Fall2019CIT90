package main

import "testing"

func BenchmarkLaunchRoutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchRoutines()
	}
}
// 2328870 ns/op	   41453 B/op	     273 allocs/op