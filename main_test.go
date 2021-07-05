package main

import (
	"testing"
)

func Benchmark_NextChronon(n *testing.B) {
	initWorld()
	for i := 0; i < 1000; i++ {
		nextChronon()
	}
}
