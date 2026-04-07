// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Akoh Chubiyojo Benjamin]
// Squad:  [Benchmarkers]

package main

import (
	"fmt"
	"os"
	"strings"
)

func pipetrail(word string) string {

	word = strings.TrimSpace(word)
	word = strings.Trim(word, "-")
	if word == "" {
		return strings.TrimSpace(word)
	}
	if strings.Contains(word, "TODO") && word == strings.ToUpper(word) {
		word = strings.ReplaceAll(word, "TODO", "Action")
		word = strings.ToLower(word)
		return strings.Title(word)
	}
	if strings.Contains(word, "TODO") {
		word = strings.ReplaceAll(word, "TODO", "Action")
		return word
	}
	for _, r := range word {
		if string(r) == strings.ToUpper(string(r)) {
			wordss := strings.Fields(word)
			for i := 0; i < len(wordss); i++ {
				wordss[i] = strings.ToUpper(string(wordss[i][0])) + strings.ToLower(string(wordss[i][1:]))
			}
			return strings.Join(wordss, " ")
		}
		if string(r) == strings.ToLower(string(r)) {
			return strings.ToUpper(word)
		}
	}
	return word
}

func main() {
	car := ""
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	words := strings.Split(string(data), "\n")
	fmt.Println("lines read: ", len(words))
	for i := 0; i < len(words); i++ {
		words[i] = pipetrail(words[i])
		words[i] = strings.TrimSpace(words[i])
		car += words[i] + "\n"
	}
	lines := strings.Split(car, "\n")
	filtered := ""
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			if filtered != "" {
				filtered += "\n"
			}
			filtered += line
		}
	}
	car = filtered
	fmt.Println("lines written: ", len(strings.Split(car, "\n")))
	fmt.Println("lines removed: ", len(words)-len(strings.Split(car, "\n")))
	fmt.Println("RULES APPLIED:\n trimspace \n TODO - Action \n remove empty or blank space \n uppercase line to titlecase \n lowercase to uppercase")
	err = os.WriteFile("output.txt", []byte(car), 0755)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
