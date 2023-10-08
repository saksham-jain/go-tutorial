package main

import "fmt"

func main() {
	var ptr *int
	num := 24
	fmt.Println("Value of ptr is", ptr)
	ptr = &num
	fmt.Println("Value of ptr is", ptr)
	fmt.Println("Value at ptr is", *ptr)
	*ptr = *ptr + 2
	fmt.Println("Value of num is", num)
}
