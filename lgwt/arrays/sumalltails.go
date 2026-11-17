package arrays

func SumAllTails(numlist ...[]int) []int {
	numOfSlices := len(numlist)
	retSlices := make([]int, numOfSlices)

	for i, v := range numlist {
		if len(v) == 0 || len(v) == 1 {
			retSlices[i] = 0
		} else {
			retSlices[i] = Sum(v[1:])
		}

	}
	return retSlices
}
