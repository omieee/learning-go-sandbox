package arrays

func Sum(numlist []int) int {
	sum := 0
	for _, v := range numlist {
		sum += v
	}
	return sum
}
