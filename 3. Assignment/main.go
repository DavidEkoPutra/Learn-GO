package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, vInts := range ints {
		if vInts%2 == 0 {
			fmt.Println(vInts, "is even")
		} else {
			fmt.Println(vInts, "is odd")
		}
	}
}
