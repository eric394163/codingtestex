package main

import "fmt"

func main() {
	a := []int{1, 5, 2, 9, 8, 3, 7, 10}
	fmt.Println(len(a))
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				tmp := a[j]
				a[j] = a[j-1]
				a[j-1] = tmp
			}
		}
	}
	fmt.Println(a)
}
