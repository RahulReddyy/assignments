package main

import (
	"fmt"
	"gocons/SlicesMaps/Problem1"
	"gocons/SlicesMaps/Problem2"
	"gocons/SlicesMaps/Problem3"
	"gocons/SlicesMaps/Problem4"
)

func main() {

	//Problem1
	var sliceProblem1 = []int{1, 2, 3, 4, 5}
	eslice, oslice := Problem1.EvenOdd(sliceProblem1)
	fmt.Println("even Slice", eslice)
	fmt.Println("odd Slice", oslice)
	fmt.Println()

	//Problem2
	var sliceProblem2 = []int{1, 2, 2, 3, 4, 5}
	dupMap := Problem2.Duplicates(sliceProblem2)
	fmt.Println("Duplicateamp", dupMap)
	fmt.Println()

	//Problem3
	string1 := "madam"
	value := Problem3.Palindrome(string1)
	fmt.Println("Is Palindrome", value)
	fmt.Println()

	//Problem4
	sampleStirng := "this is a sample string"
	charCheck := 's'
	sliceIndex := Problem4.IndexString(sampleStirng, charCheck)
	fmt.Println("sliceIndex", sliceIndex)
	fmt.Println()
}
