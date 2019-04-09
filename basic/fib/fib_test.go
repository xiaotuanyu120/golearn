package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	n := 6
	t.Log(fib(n))
}

func TestFibList(t *testing.T) {
	n := 6
	t.Log(fibList(n))
}

func fib(n int) int {
	if n <= 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fibList(n int) []int {
	fib_list := make([]int, n, n)
	a, temp := 1, 0
	for i := 0; i < n; i++ {
		fib_list[i] = a
		a, temp = temp+a, a
	}
	return fib_list
}
