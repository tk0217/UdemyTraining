package main

import "fmt"

func main() {

	var small int
	var large int
	fmt.Print("Type a large number: ")
	fmt.Scan(&large)
	fmt.Print("Type a smaller number: ")
	fmt.Scan(&small)
	fmt.Println(large, "/", small, " = ", large/small)
	fmt.Println("The remainder of ", large, "/", small, " = ", large%small)
}
