package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)

	delete(languages, "JS")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
