package main

import (
	"fmt"
	"strings"
)

func fixSpace(input string) string {
	result := ""

	for _, item := range input {
		if strings.ContainsAny(string(item), ",.?;") && result[len(result)-1] == ' ' {
			result = result[:len(result)-1]
			result += string(item) + " "
			continue
		}
		if !strings.ContainsAny(string(item), ";,?'") {
			result += string(item)
			continue
		}
		result += string(item)
	}
	return result
}
func main() {
	fmt.Println(fixSpace("this is , word from henceforth, hiring"))
}
