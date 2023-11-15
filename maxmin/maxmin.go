package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5}

	max := a[0]
	min := a[0]

	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
	}

	fmt.Println(max)
	fmt.Println(min)
}
