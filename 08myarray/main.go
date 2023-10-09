package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "mango"
	fruitList[2] = "banana"
	fruitList[3] = "grapes"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	fmt.Printf("Type of fruitList is: %T", fruitList)

	var fruitList2 = [3]string{"apple", "orange"}

	fmt.Println(fruitList2)
	fmt.Println(len(fruitList2))
	fmt.Printf("Type of fruitList2 is: %T", fruitList2)
}
