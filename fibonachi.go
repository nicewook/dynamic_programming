package fibo

func fib(n uint64) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibMemoize(n uint64, memo []uint64) uint64 {
	if memo[n] != 0 {
		return memo[n]
	}
	if n == 1 || n == 2 {
		memo[n] = 1
		return 1
	}
	memo[n] = fibMemoize(n-1, memo) + fibMemoize(n-2, memo)
	return memo[n]
}

func fibBottomUp(n uint64) uint64 {
	if n == 1 || n == 2 {
		return 1
	}

	bottomUp := make([]uint64, n+1)
	bottomUp[1] = 1
	bottomUp[2] = 1
	for i := uint64(3); i < n+1; i++ {
		bottomUp[i] = bottomUp[i-1] + bottomUp[i-2]
	}

	return bottomUp[n]
}
