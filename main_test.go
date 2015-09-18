package main

import (
	"testing"
	"io/ioutil"
	"runtime"
	"github.com/nmerouze/selfjs"
)

func benchmarkRender(i int, b *testing.B) {
	bundle, _ := ioutil.ReadFile("./build/bundle.js")
	fetchFeed()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		selfjs.New(runtime.NumCPU(), string(bundle), rss)
	}
}

func BenchmarkRender1(b *testing.B) { benchmarkRender(1, b) }
