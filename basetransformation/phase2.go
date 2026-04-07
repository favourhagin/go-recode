/*THE BASE CONVERTER
 Concept: Number Systems & strconv

Rules:
 → DO NOT USE GOOGLE OR ANY AI TOOL; DO IT YOURSELF - you can share ideas with others, but not copy code.
 → Write everything in Go. Standard library only.
 → Must compile and run without errors.
 → Push to your the-codecrafters repo in a folder
   named: thecodecrafterthon-day2/
 → Include a README.md explaining how to run it

 go-reloaded needs to convert hex and binary
 strings into decimal numbers. This project
 teaches you exactly that — and makes you think
 about what happens when the input is garbage.

 Build a CLI tool that converts numbers between
 bases. It runs from the terminal like this:

    > convert 1E hex
      ✦ Decimal: 30

    > convert 10 bin
      ✦ Decimal: 2

    > convert 255 dec
      ✦ Binary:  11111111
      ✦ Hex:     FF

 Requirements:

 → Support three input bases: hex, bin, dec.
 → For dec input, output both binary and hex.
 → For hex and bin input, output only decimal.
 → All hex output must be uppercase.
 → The program keeps running until: quit

 Validation — handle all of these cleanly:
 → "1G" is not valid hex.
 → "10201" is not valid binary.
 → "abc" is not a valid decimal.
 → Negative numbers must be supported for dec.
 → Empty input must not crash the program.

 Think about:
 → Which strconv functions handle base
   conversions for you?
 → How do you detect which characters are
   valid for a given base?
 → What is the difference between ParseInt
   and ParseUint — and does it matter here?

 ✦ Nail this and (hex) / (bin) in go-reloaded
   becomes a 5-minute task.*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("------------------------------------")
	fmt.Println("WELCOME TO THE CHUBIYOJO'S BASE TRANSFORMER")
	fmt.Println("------------------------------------")
	fmt.Println("this are our programs:")
	fmt.Println(" ")
	fmt.Println("bin and hex converts binary and hexadecimal to decimal")
	fmt.Println("dec converts decimal to binary and hexadecimal")
	fmt.Println("hexadecimal")
	fmt.Println("binary")
	fmt.Println("decimal")
	fmt.Println("if you want to quit, type quit")
start:
	for {
		fmt.Println("select a program:")
		fmt.Scan(&input)
		if input == "quit" {
			fmt.Println("it is a pleasure to work with you")
			break
		}
		input = strings.ToLower(input)
		number := ""
		digit := ""
		digits := ""
		if input == "hexadecimal" {
			fmt.Println("input number to be converted:")
			fmt.Scan(&number)
			for _, r := range number {
				if r < '0' || r > 'f' && r <= 'z' || r > 'F' && r <= 'Z' {
					fmt.Println("invalid hexadecimal character")
					goto start
				}
				if r >= '0' && r <= '9' || r >= 'A' && r <= 'F' || r >= 'a' && r <= 'f' {
					h, _ := (strconv.ParseInt(number, 16, 64))
					fmt.Println(h)
					goto start
				}
			}
		}
		if input == "binary" {
			fmt.Println("input number to be converted:")
			fmt.Scan(&digit)
			for _, r := range digit {
				if r < '0' || r > '1' {
					fmt.Println("invalid binary character")
					goto start
				}
				if r >= 0 && r <= '1' {
					man, _ := strconv.ParseInt(digit, 2, 64)
					fmt.Println(man)
					goto start
				}
			}
		}
		if input == "decimal" {
			fmt.Println("input number to be converted:")
			fmt.Scan(&digits)
			if strings.ContainsAny(digits, "abcdefghijklmnopqrstuvwxyz") {
				fmt.Println("invalid decimal charater")
			} else {
				cart, _ := strconv.ParseInt(digits, 10, 64)
				fmt.Println(strconv.FormatInt(cart, 2))
				fmt.Println(strings.ToUpper(strconv.FormatInt(cart, 16)))
				goto start
			}
		}
	}
}
