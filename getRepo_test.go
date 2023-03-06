package main

import (
	"testing"
)

var (
	res []RepoData
	err error
)

func BenchmarkGetStarredRepo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err = getStarredRepo("nikusaikou", 140)
	}
}

func BenchmarkGetReposGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, err = getReposGoroutine("nikusaikou", 140)
	}
}
