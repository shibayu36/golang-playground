package main

import "fmt"

func main() {
	var s []int
	//	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	s = append(s, 2)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	s = s[:4]
	printSlice(s)
	s = s[:8]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = append(s, 1)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("mem=%p len=%d cap=%d %v\n", &s[0], len(s), cap(s), s)
}
