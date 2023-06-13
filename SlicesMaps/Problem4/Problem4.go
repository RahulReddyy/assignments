package Problem4

// sample string - “this is a sample string”
//      Given character - “S”

func IndexString(data string, check rune) int {
	slice := []int{}
	count := 0
	for index, value := range data {
		if check == value {
			slice = append(slice, index)
			count++
		}
	}
	midValue := count / 2
	return slice[midValue]
}
