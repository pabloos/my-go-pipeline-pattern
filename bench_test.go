package main

import "testing"

func BenchmarkPipeline(b *testing.B) {
	pip := NewPipeline(start, final, add2, square, add2)

	for n := 0; n < b.N; n++ {
		pip.Exec(1+n, 2+n, 3+n, 4+n, 5+n)
	}
}

func BenchmarkPipeline2(b *testing.B) {
	pip := NewPipeline(start, final, add2, square, add2)

	for n := 0; n < b.N; n++ {
		pip.Exec(1+n, 2+n, 3+n, 4+n, 5+n, 6+n, 7+n, 8+n, 9+n)
	}
}