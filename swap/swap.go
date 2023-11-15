package main

import "fmt"

func main() {
	a := 10
	b := 20
	tmp := 0

	tmp = b
	b = a
	a = tmp

	fmt.Println(a, b, tmp)

}
