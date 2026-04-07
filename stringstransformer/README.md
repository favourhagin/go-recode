# String Transformer (Go CLI)

This is a Command Line Interface (CLI) project built with Go. It allows users to transform and analyze text using different string operations.

# Features

The program supports different string transformations and utilities:

Transformations
up → converts all text to uppercase
low → converts all text to lowercase
cap → capitalizes the first letter of each word
title → capitalizes main words but keeps small connector words lowercase
snake → converts text to snake_case (removes symbols)
reverse → reverses each word in the input
Utilities
count → shows:
total characters
total letters only
total words
total spaces
palindrome → checks if a word or sentence is a palindrome
history → shows the last five transformations
exit → exits the program

# How It Works
The program runs in the terminal.
It uses bufio to read user input.
The user selects a command.
Then the user enters a word or sentence.
The program processes the input using different functions.
The result is printed to the screen.
The program keeps running until the user types exit.


# Error Handling

The program handles errors so it does not crash:

Empty input → shows a message
Unknown command → shows valid commands
Missing text → prompts user again

# History Feature
The program stores previous results.
The history command shows the last five operations performed.