package arrays

func SumAll(allNums ...[]int) []int {
	numOfSlices := len(allNums)
	retSlices := make([]int, numOfSlices)

	for i, v := range allNums {
		retSlices[i] = Sum(v)
	}
	return retSlices
}
