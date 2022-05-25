// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	sli := strings.Split(data, " ")
	nums := sli[:]
	var intnums []int
	var evens []int
	var odds []int
	var middle []int
	var first2 []int
	var last2 []int
	var evensl1 []int
	var oddsl2 []int
	for _, v := range sli {
		a, _ := strconv.Atoi(v)
		if a%2 == 0 {
			evens = append(evens, a)
		}
		if a%2 == 1 {
			odds = append(odds, a)
		}
		intnums = append(intnums, a)
	}
	middle = intnums[len(intnums)/2-1 : len(intnums)/2+1]
	first2 = intnums[:2]
	last2 = intnums[len(intnums)-2:]
	evensl1 = evens[len(evens)-1:]
	oddsl2 = odds[len(odds)-2:]
	fmt.Println("nums : ", nums)
	fmt.Println("evens : ", evens)
	fmt.Println("odds : ", odds)
	fmt.Println("middle :", middle)
	fmt.Println("first 2 :", first2)
	fmt.Println("last 2 :", last2)
	fmt.Println("evens last 1 :", evensl1)
	fmt.Println("odds last 2 :", oddsl2)
}
