package Problem1

func EvenOdd(testSlice []int) (eSlice, oSlice []int) {

	var evenSlice []int
	var oddSlice []int

	for index := range testSlice {
		if testSlice[index]%2 == 0 {
			evenSlice = append(evenSlice, testSlice[index])
		} else {
			oddSlice = append(oddSlice, testSlice[index])
		}
	}

	return evenSlice, oddSlice
}
