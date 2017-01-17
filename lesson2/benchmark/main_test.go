package main

import "testing"

func BenchmarkAverage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		average([]float64{1, 2})
	}
}
