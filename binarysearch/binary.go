package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	find := -1
	left := 0
	right := len(a) - 1
	want := 8

	for {
		if left <= right {
			middle := int((left + right) / 2)
			if a[middle] == want {
				find = middle
				break
			} else if a[middle] < want {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			break
		}
	}

	fmt.Println(find)
}
