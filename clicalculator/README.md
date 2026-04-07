# CLI Calculator (Go)

This is a simple Command Line Interface (CLI) calculator built with Go. It runs in the terminal and allows the user to perform basic arithmetic operations.

# Features

The calculator supports the following operations:

add <a> <b> → addition
subtract <a> <b> → subtraction
multiply <a> <b> → multiplication
divide <a> <b> → division

It keeps running until the user decides to quit.

# How It Works
The program starts and shows a welcome message with examples.
The user types an operation in the terminal.
The input is read using bufio.
The input is cleaned and split using strings.
The program checks:
if the command is valid
if the correct number of arguments is given
if the inputs are numbers
Then it performs the calculation and prints the result.
Commands
help → shows all available operations
quit → exits the program

# Error Handling

The program handles errors properly so it does not crash:

Invalid command → shows correct examples
Wrong number of arguments → shows error
Letters instead of numbers → shows error
Division by zero → shows error message