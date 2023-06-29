//Problem 28

package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
func main() {

	haystack := "sadbutsad"
	needle := "sad"
	fmt.Println(strStr(haystack, needle))
}
