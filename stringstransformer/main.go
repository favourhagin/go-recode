// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Akoh Chubiyojo Benjamin]
// Squad:  [The BenchMarkers]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func upper(word []string) string {
	car := ""
	for i := 0; i < len(word); i++ {
		car += strings.ToUpper(word[i]) + " "
	}
	return car
}

func lower(word []string) string {
	car := ""
	for i := 0; i < len(word); i++ {
		car += strings.ToLower(word[i]) + " "
	}
	return car
}

func capfirstchar(word []string) string {
	car := ""
	for i := 0; i < len(word); i++ {
		car += strings.ToLower(word[i]) + " "
		car = strings.Title(car)
	}
	return car
}

func snakecase(word []string) string {
	car := ""
	man := ""

	for i := 0; i < len(word); i++ {
		clean := ""
		for _, r := range word[i] {
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				clean += string(r)
			}
		}
		car += strings.ToLower(clean) + " "
	}

	for j := 0; j < len(car)-1; j++ {
		man += string(car[j])
		man = strings.ReplaceAll(man, " ", "_")
	}

	return man
}

func reverse(word []string) string {
	car := ""
	for i := 0; i < len(word); i++ {
		word[i] = word[i] + " "
		for j := len(word[i]) - 1; j >= 0; j-- {
			car += string(word[i][j])
		}
	}
	return car
}

func countchar(word string) int {
	count := 0
	for i := 0; i < len(word)-1; i++ {
		count++
	}
	return count
}
func countletter(word string) int {
	count := 0
	for i := 0; i < len(word)-1; i++ {
		if (string(word[i]) >= "a" && string(word[i]) <= "z") || (string(word[i]) >= "A" && string(word[i]) <= "Z") {
			count++
		}
	}
	return count
}
func totalword(word string) int {
	t := strings.Fields(word)
	return len(t)
}
func space(word string) int {
	count := 0
	for i := 0; i < len(word)-1; i++ {
		if word[i] == ' ' {
			count++
		}
	}
	return count
}
func palindrome(word string) string {
	d := ""
	c := ""
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		if word[i] != ' ' {
			d += string(word[i])
		}
	}
	for i := len(d) - 1; i >= 0; i-- {
		c += string(d[i])
	}
	//return c
	if d == c {
		return "it is a palindrome"
	}
	return "it is not a palindrome"
}

// reader := bufio.NewReader(os.Stdin)
// input,_ := reader.ReadString('\n')
func main() {
	var history []string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("------------------------------")
	fmt.Println("WELCOME TO THE CHUBIYOJO'S STRINGS TRANSFORMER")
	fmt.Println("-----------------------------------")
	fmt.Println("programs you can run:")
	fmt.Println("-----------------------------------")
	fmt.Println("up : capitalize the entire character")
	fmt.Println("low: returns the entire input as lower case")
	fmt.Println("cap: returns the input with first letter of each word capitalize")
	fmt.Println("title: returns the input with first of each non connector words toupper and connectors as they are")
	fmt.Println("reverse : reverses each word in the input")
	fmt.Println("snake: converts input to snake case")
	fmt.Println("exit: ends the program")
	fmt.Println("history: returns last five transformation")
	fmt.Println("count: - Total characters - Total letters only - Total words - Total spaces")
	fmt.Println("palindrome: checks if a string is a palindrome")
start:
	fmt.Println("enter program you want:")
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)
		if input == "" {
			fmt.Println("you have to select a program")
			goto start
		}
		if input != "up" && input != "palindrome" && input != "count" && input != "low" && input != "cap" && input != "history" && input != "title" && input != "reverse" && input != "exit" && input != "snake" {
			fmt.Println(" Unknown command:", input)
			fmt.Println("Valid commands: up, low, cap,title, snake, reverse, exit")
			goto start
		}
		if input == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}
		if input == "up" {
			fmt.Println("word or sentence to be worked on:")
			scans, _ := reader.ReadString('\n')
			scant := strings.Fields(scans)
			if len(scant) == 0 {
				fmt.Println("No text provided. Usage: low <text>")
				goto start
			}
			fmt.Println("result:")
			fmt.Println(upper(scant))
			history = append(history, scans+" up"+" "+upper(scant)+",")
			goto start
		}
		if input == "low" {
			fmt.Println("word or sentence to be worked on:")
			scans, _ := reader.ReadString('\n')
			scant := strings.Fields(scans)
			if len(scant) == 0 {
				fmt.Println("No text provided. Usage: low <text>")
				goto start
			}
			fmt.Println("result:")
			fmt.Println(lower(scant))
			history = append(history, scans+" low"+" "+lower(scant)+",")
			goto start
		}
		if input == "cap" {
			fmt.Println("word or sentence to be worked on:")
			scans, _ := reader.ReadString('\n')
			scant := strings.Fields(scans)
			if len(scant) == 0 {
				fmt.Println("No text provided. Usage: cap <text>")
				goto start
			}
			fmt.Println("result:")
			fmt.Println(capfirstchar(scant))
			history = append(history, scans+" cap"+" "+capfirstchar(scant)+",")
			goto start
		}
		if input == "snake" {
			fmt.Println("word or sentence to be worked on:")
			scans, _ := reader.ReadString('\n')
			scant := strings.Fields(scans)
			if len(scant) == 0 {
				fmt.Println("No text provided. Usage: snake <text>")
				goto start
			}
			fmt.Println("result:")
			fmt.Println(snakecase(scant))
			history = append(history, scans+" snake"+" "+snakecase(scant)+",")
			goto start
		}
		if input == "reverse" {
			fmt.Println("word or sentence to be worked on:")
			scans, _ := reader.ReadString('\n')
			scant := strings.Fields(scans)
			if len(scant) == 0 {
				fmt.Println("No text provided. Usage: reverse <text>")
				goto start
			}
			fmt.Println("result:")
			fmt.Println(reverse(scant))
			history = append(history, scans+" reverse"+" "+upper(scant)+",")
			goto start
		}
		if input == "title" {
			car := ""
			fmt.Println("word or sentence to be worked on:")
			scans, _ := reader.ReadString('\n')
			scant := strings.Fields(scans)
			if len(scant) == 0 {
				fmt.Println("No text provided. Usage: title <text>")
				goto start
			}
			for i := 0; i <= len(scant)-1; i++ {
				if i > 0 && (scant[i] == "a" || scant[i] == "an" || scant[i] == "the" || scant[i] == "and" || scant[i] == "but" || scant[i] == "or" || scant[i] == "for" || scant[i] == "nor" || scant[i] == "on" || scant[i] == "at" || scant[i] == "to" || scant[i] == "by" || scant[i] == "in" || scant[i] == "of" || scant[i] == "up" || scant[i] == "as" || scant[i] == "is" || scant[i] == "it") {
					car += scant[i] + " "
				} else {
					scant[i] = strings.ToLower(scant[i])
					car += strings.Title(scant[i]) + " "
				}
			}
			fmt.Println(car)
			history = append(history, scans+" title"+" "+car+",")
			goto start
		}
		if input == "count" {
			fmt.Println("1 Total characters")
			fmt.Println("2 Total letters only")
			fmt.Println("3 Total words")
			fmt.Println("4 - Total spaces")
			fmt.Println("select a number:")
			count, _ := reader.ReadString('\n')
			count = strings.TrimSpace(count)
			if count == "1" {
				fmt.Println("word or sentence to be worked on:")
				scans, _ := reader.ReadString('\n')
				if len(scans) == 0 {
					fmt.Println("No text provided. Usage: title <text>")
					goto start
				}
				fmt.Println(countchar(scans))
				history = append(history, scans+" totalcharacter"+" ", strconv.Itoa(countchar(scans))+",")
			}
			if count == "2" {
				fmt.Println("word or sentence to be worked on:")
				scans, _ := reader.ReadString('\n')
				if len(scans) == 0 {
					fmt.Println("No text provided. Usage: title <text>")
					goto start
				}
				fmt.Println(countletter(scans))
				history = append(history, scans+" totalletter"+" ", strconv.Itoa(countletter(scans))+",")
			}
			if count == "3" {
				fmt.Println("word or sentence to be worked on:")
				scans, _ := reader.ReadString('\n')
				if len(scans) == 0 {
					fmt.Println("No text provided. Usage: title <text>")
					goto start
				}
				fmt.Println(totalword(scans))
				history = append(history, scans+" totalword"+" ", strconv.Itoa(totalword(scans))+",")
			}
			if count == "4" {
				fmt.Println("word or sentence to be worked on:")
				scans, _ := reader.ReadString('\n')
				if len(scans) == 0 {
					fmt.Println("No text provided. Usage: title <text>")
					goto start
				}
				fmt.Println(space(scans))
				history = append(history, scans+" totalspace"+" ", strconv.Itoa(space(scans))+",")
			}
			goto start
		}
		if input == "palindrome" {
			fmt.Println("word or sentence to be worked on:")
			scans, _ := reader.ReadString('\n')
			scans = strings.TrimSpace(scans)
			if len(scans) == 0 {
				fmt.Println("No text provided. Usage: snake <text>")
				goto start
			}
			fmt.Println("result:")
			fmt.Println(palindrome(scans))
			history = append(history, scans+" palindrome"+" "+palindrome(scans)+",")
			goto start
		}
		if input == "history" {
			if len(history) < 5 {
				fmt.Println(history)
			} else {
				fmt.Println("HISTORY 1:")
				fmt.Println(history[len(history)-1])
				fmt.Println("HISTORY 2:")
				fmt.Println(history[len(history)-2])
				fmt.Println("HISTORY 3:")
				fmt.Println(history[len(history)-3])
				fmt.Println("HISTORY 4:")
				fmt.Println(history[len(history)-4])
				fmt.Println("HISTORY 5:")
				fmt.Println(history[len(history)-5])
			}
			goto start
		}
	}
}
