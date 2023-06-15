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
	if len(slice)%2 == 0 {
		return slice[1]
	} else {
		return slice[count/2]
	}
}
