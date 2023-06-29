package main

import (
	"fmt"
)

func isValid(s string) bool {
	str := s
	slice := []rune{} // using stack concept
	for _, char := range str {
		if char == '(' || char == '[' || char == '{' {
			slice = append(slice, char)
		} else {

			if len(slice) == 0 {
				return false
			}
			top := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			if (char == ')' && top != '(') || (char == '}' && top != '{') || (char == ']' && top != '[') {
				return false
			}
		}
	}
	return len(slice) == 0
}
func main() {
	str := ")()())"
	fmt.Println(isValid(str))
}
