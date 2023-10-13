package main

import "fmt"

func main() {
	sj := User{"SJ", "SJ@gmail.com", 11, true}
	fmt.Println(sj)
	fmt.Printf("Details are %+v", sj)
	fmt.Printf("Name is %v\n", sj.Name)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
