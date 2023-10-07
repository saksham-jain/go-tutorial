package main

import "fmt"

const LoginToken = "sdfndfndf" // Public constant starts with capital letter

func main() {
	var username string = "saksham"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.34545465463453343435
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var largeFloat float64 = 255.34545465463453343435
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	// default values
	var someVariable int
	fmt.Println(someVariable)

	// implicit type (type decided by lexer)
	var website = "www.google.in"
	fmt.Println(website)

	// most common way of initializing

	aVariable := 12
	fmt.Println(aVariable)

	fmt.Println(LoginToken)
}
