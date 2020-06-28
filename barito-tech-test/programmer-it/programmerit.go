package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Input : ")
	fmt.Scan(&s)
	programmerIT(s)
}

func programmerIT(s string) {
	dictionary := []string{"pro", "gram", "merit", "program", "it", "programmer"}
	result := []string{}
	for i := 0; i < len(dictionary); i++ {
		if strings.Contains(s, dictionary[i]) {
			result = append(result, dictionary[i])
		}
	}
	if len(result) > 0 {
		fmt.Println(result)
	} else {
		fmt.Println("<no way>")
	}
}
