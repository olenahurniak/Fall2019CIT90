package main

import "testing"

func BenchmarkLaunchRoutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchRoutines()
	}
}
// 2618267 ns/op	   36617 B/op	     264 allocs/op
// 2495035 ns/op	   25047 B/op	     238 allocs/op