package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	sum := 0

	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}

	fmt.Println(sum)
}
