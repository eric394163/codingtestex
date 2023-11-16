package main

import "fmt"

func main() {

	a := []int{1, 5, 3, 9, 2, 8, 10}

	for i := 0; i < len(a)-1; i++ {
		min := a[i]
		k := i
		for j := i + 1; j < len(a); j++ {
			if min > a[j] {
				min = a[j]
				k = j
			}
		}

		tmp := a[i]
		a[i] = a[k]
		a[k] = tmp

	}
	fmt.Println(a)
}
