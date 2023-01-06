package functions

func Sum(nums ...int) (a int) {
	ans := 0

	for _, i := range nums {
		ans += i
	}

	return ans
}
