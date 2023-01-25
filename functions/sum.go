package functions

// functions with  (arr ...int) are called variadic functions
func Sum(nums ...int) (a int) {
	ans := 0

	for _, i := range nums {
		ans += i
	}

	return ans
}
