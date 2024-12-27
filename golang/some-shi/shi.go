package main

import "fmt"

func sum(list ...int) int {

	var total int = 0

	for _, x := range list {
		total += x
	}
	return total

}
func main() {
	fmt.Println("supp g")
	lol := sum(1, 2, 3, 4, 5)
	fmt.Println(lol)
}
