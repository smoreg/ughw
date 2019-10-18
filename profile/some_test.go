package main

import (
	"testing"
)

func BenchmarkInBounds(b *testing.B) {
	var mySlice []int
	for i := 0; i < 99; i++ {
		mySlice = append(mySlice, i)
	}
	ms := newMySliceType(mySlice)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slicerInBounds(ms)
	}
}

func BenchmarkInBoundsChannels(b *testing.B) {
	var mySlice []int
	for i := 0; i < 99; i++ {
		mySlice = append(mySlice, i)
	}
	ms := newMySliceType(mySlice)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slicerInBoundsChannels(ms)
	}
}
