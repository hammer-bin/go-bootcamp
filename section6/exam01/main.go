package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func main() {
	items := os.Args[1:]
	if len(items) == 0 {
		fmt.Println("Provide a parameter")
		return
	}

	sort.Strings(items)
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}

	var total int
	for _, data := range items {
		total += len(data) + 1
	}
	fmt.Printf("Total required space: %d bytes.\n", total)

	names := make([]byte, 0, total)
	for _, data := range items {
		names = append(names, data...)
		names = append(names, '\n')
	}

	err := ioutil.WriteFile("sorted.txt", names, 0644)
	if err != nil {
		return
	}

}
