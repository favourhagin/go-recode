//CODECRAFTERS — HACKATHON CHALLENGE
/*  THE CLI CALCULATOR


  Two phases. One calculator. Build it in stages.

  Rules:
  → Write everything in Go. Standard library only.
  → Must compile and run without errors.
  → Push to your the-codecrafters repo in a folder
    named: thecodecrafterthon-day1/
  → Include a README.md explaining how to run it.
  → Complete Phase 1 before opening Phase 2.
  → Deadline: 2 hours from — GO.

   PHASE 1 — GET IT WORKING

  Build a calculator that runs in the terminal,
  accepts user input, and keeps going until
  you decide to stop.

  WHAT TO BUILD

- OPERATIONS
  The user types an expression. The calculator
  responds with the result. Support:

     add <a> <b>   → addition
     sub <a> <b>   → subtraction
     mul <a> <b>   → multiplication
     div <a> <b>   → division

  Example interaction:
     > add 10 4
       ✦ Result: 14

     > mul 3 7
       ✦ Result: 21

       -  KEEP IT RUNNING
  The calculator must not stop after one result.
  It waits for the next input, and the next,
  until the user types:

     quit   → exits with a goodbye message

- HELP
  Typing help prints a clean list of all
  supported commands and how to use them.

  VALIDATION — HANDLE ALL OF THESE CLEANLY

  → Division by zero
    must print a clear error. Not a crash.

  → Letters where numbers are expected
    e.g. "add cat dog"
    must print a clear error. Not a crash.

  → Wrong number of arguments
    e.g. "add 5" or "mul 1 2 3"
    must be caught and explained to the user.

  → Unknown commands
    must suggest typing "help".

  → Negative numbers must work perfectly.
    e.g. "add -10 -5  →  -15"

  The golden rule:
  ✦ Nothing the user types should ever
    crash your program. Ever.

  THINK ABOUT

  → How do you read and split user input cleanly?
  → How do you validate before you calculate?
  → What does a helpful error message look like?
  → How do you structure your operations so the
    code doesn't become one giant if-else chain?
reader := bufio.NewReader(os.Stdin)
input,_ := reader.ReadString('\n')
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("------------------------------------")
	fmt.Println("----------WELCOME TO CHUBIYOJO'S CLI CALCULATOR")
	fmt.Println("-------------------------------------")
	fmt.Println("examples:")
	fmt.Println("add <a> <b>")
	fmt.Println("subtract <a> <b>")
	fmt.Println("divide <a> <b>")
	fmt.Println("multiply <a> <b>")
	fmt.Println("quit")
	fmt.Println("help")
	for {
	start:
		fmt.Println("enter operation expression(add):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "quit" {
			fmt.Println("----------------------------")
			fmt.Println("thanks for your time,goodbye")
			fmt.Println("----------------------------")
			break
		}
		if input == "help" {
			fmt.Println("examples of operations that can run:")
			fmt.Println("add <a> <b>")
			fmt.Println("subtract <a> <b>")
			fmt.Println("divide <a> <b>")
			fmt.Println("multiply <a> <b>")
		}
		slice := strings.Fields(input)
		slice[0] = strings.ToUpper(slice[0])
		if slice[0] != "ADD" && slice[0] != "SUBTRACT" && slice[0] != "MULTIPLY" && slice[0] != "DIVIDE" {
			fmt.Println("invalid operation")
			fmt.Println("examples of valid operation:")
			fmt.Println("add <a> <b>")
			fmt.Println("subtract <a> <b>")
			fmt.Println("divide <a> <b>")
			fmt.Println("multiply <a> <b>")
			goto start
		}
		if len(slice) > 3 {
			fmt.Println("more than 2 numbers")
			goto start
		}
		if len(slice) < 3 {
			fmt.Println("numbers less than 2")
			goto start
		}
		for i := 0; i < len(slice[1]); i++ {
			r := rune(slice[1][i])
			if unicode.IsLetter(r) {
				fmt.Println("only digits needed")
				goto start
			}
		}
		for i := 0; i < len(slice[2]); i++ {
			r := rune(slice[2][i])
			if unicode.IsLetter(r) {
				fmt.Println("only digits needed")
				goto start
			}
		}
		a, _ := strconv.ParseFloat(slice[1], 64)
		b, _ := strconv.ParseFloat(slice[2], 64)

		if slice[0] == "ADD" {
			fmt.Println(a + b)
			goto start
		}
		if slice[0] == "SUBTRACT" {
			fmt.Println(a - b)
			goto start
		}
		if slice[0] == "MULTIPLY" {
			fmt.Println(a * b)
			goto start
		}
		if slice[0] == "DIVIDE" && slice[2] == "0" {
			fmt.Println("error: can't divide a number by 0")
			goto start
		}
		if slice[0] == "DIVIDE" {
			fmt.Println(a / b)
			goto start
		}
	}

}
