package main

import "fmt"

func main() {
	var fruitList = []string{"apple", "banana"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "mango", "grape")
	fmt.Println(fruitList)

	fruitList = fruitList[1:3]
	fmt.Println(fruitList)

	// example of how to remove a value from slice
	var fruitList2 = []string{"apple", "banana", "mango", "banana"}
	index := 2
	fruitList2 = append(fruitList2[:index], fruitList2[index+1:]...)
	fmt.Println(fruitList2)
}
