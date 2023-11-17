package main

import "fmt"

func main() {
	a := []int{1, 5, 3, 2, 9, 7, 10, 4}

	for i := 0; i < len(a); i++ {
		tmp := a[i]
		insert := 0

		for j := i - 1; j >= 0; j-- {
			if a[j] > tmp {
				a[j+1] = a[j]
			} else {
				insert = j + 1
				break
			}
		}
		a[insert] = tmp

	}
	fmt.Println(a)
}
