package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user unput"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the Pizza: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
}
