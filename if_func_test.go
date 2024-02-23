//go:build ignore
// +build ignore

package main

// goos: linux
// goarch: amd64
// pkg: github.com/sceneq/wslpath/bench
// cpu: Intel(R) Core(TM) i7-6700K CPU @ 4.00GHz
// BenchmarkA-8    745224970                1.535 ns/op
// BenchmarkB-8    1000000000               0.3895 ns/op
// PASS
// ok      github.com/sceneq/wslpathgo/bench 1.751s

import (
	"math/rand"
	"testing"
)

func process1() int {
	return 0
}

func process2() int {
	return 1
}

var global int

func BenchmarkA(b *testing.B) {
	var processFunc func() int
	condition := rand.Float32() > 0.5
	if condition {
		processFunc = process1
	} else {
		processFunc = process2
	}

	var temp int
	for i := 0; i < b.N; i++ {
		temp = processFunc()
	}
	global = temp
}

func BenchmarkB(b *testing.B) {
	condition := rand.Float32() > 0.5

	var temp int
	for i := 0; i < b.N; i++ {
		if condition {
			temp = process1()
		} else {
			temp = process2()
		}
	}
	global = temp
}
