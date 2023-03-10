package leetcode

var cache map[int]int

func initialize() {
	cache = make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	cache[2] = 1
}

func Tribonacci(n int) int {
	if cache == nil {
		initialize()
	}

	val, ok := cache[n]
	if !ok {
		return val
	}
	val = Tribonacci(n-1) + Tribonacci(n-2) + Tribonacci(n-2)
	cache[n] = val
	return cache[n]
}
