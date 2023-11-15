package main

import "fmt"

func main() {

	a := []int{1, 4, 2, 5, 10}

	find := -1

	want := 10

	for i := 0; i < len(a); i++ {
		if want == a[i] {
			find = i
			break
		}
	}

	fmt.Println(find)

}
