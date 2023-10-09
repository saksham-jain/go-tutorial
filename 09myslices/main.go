package main

import "fmt"

func main() {
	var fruitList = []string{"apple", "banana"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "mango", "grape")
	fmt.Println(fruitList)

	fruitList = fruitList[1:3]
	fmt.Println(fruitList)
}
