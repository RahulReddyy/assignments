package Problem3

//palindrom = "tacocat"

func Palindrome(palindrom string) bool {

	mid := (len(palindrom) / 2)
	maxLength := len(palindrom)
	for index := range palindrom {
		if index != mid && maxLength != mid {
			if palindrom[index] != palindrom[maxLength-1] {
				return false
			}

		}
		maxLength--
	}
	return true
}
