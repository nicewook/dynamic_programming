package fibo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(t *testing.T) {
	// t.Skip()
	number := 50
	expected := uint64(12586269025)
	actual := fib(uint64(number))
	fmt.Printf("fib %d is %d\n", number, actual)
	assert.Equal(t, expected, actual, "they should be equal")
}

func TestFibMemoize(t *testing.T) {
	// t.Skip()
	number := 50
	memo := make([]uint64, 1000)
	expected := uint64(12586269025)
	actual := fibMemoize(uint64(number), memo)
	fmt.Printf("fibMemoize %d is %d\n", number, actual)
	assert.Equal(t, expected, actual, "they should be equal")
}

func TestFibBottomUp(t *testing.T) {
	// t.Skip()
	number := 50
	expected := uint64(12586269025)
	actual := fibBottomUp(uint64(number))
	fmt.Printf("fibBottomUp %d is %d\n", number, actual)
	assert.Equal(t, expected, actual, "they should be equal")
}

func BenchmarkFibo(b *testing.B) {
	// t.Skip()
	for i := 0; i < b.N; i++ {
		number := 50
		fib(uint64(number))
	}
}

func BenchmarkFibMemoize(b *testing.B) {
	// t.Skip()
	for i := 0; i < b.N; i++ {
		number := 1500
		memo := make([]uint64, 2000)
		fibMemoize(uint64(number), memo)
	}
}

func BenchmarkFibBottomUp(b *testing.B) {
	// t.Skip()
	for i := 0; i < b.N; i++ {
		number := 1500
		fibBottomUp(uint64(number))
	}
}
