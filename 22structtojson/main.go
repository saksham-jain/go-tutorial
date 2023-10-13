package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string
	Price    float32
	Platform string
	Password string
	Tags     []string
}

func main() {
	var courses = []Course{
		{"Python", 30, "github", "abc123", []string{"Coding", "Programming"}},
		{"Kotlin", 200.10, "github", "bvn3wr3", []string{"Coding", "Programming"}},
		{"Rails", 100, "youtube", "fdsfds353", []string{"Coding", "Programming"}},
	}

	EncodeJson(courses)
}

func EncodeJson(courses []Course) {
	fmt.Println(courses)

	finalJson, err := json.Marshal(courses)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

	finalJsonIndent, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJsonIndent)
}
