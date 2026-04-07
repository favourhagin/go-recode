# Base Converter (Go CLI)

This is a Command Line Interface (CLI) project built with Go. It converts numbers between different number systems: binary, decimal, and hexadecimal.

# Features

The program supports the following conversions:

binary → converts binary numbers to decimal
hexadecimal → converts hexadecimal numbers to decimal
decimal → converts decimal numbers to:
binary
hexadecimal (in uppercase)

The program keeps running until the user types quit.

How It Works
The program starts and displays available options.
The user selects a conversion type:
binary
hexadecimal
decimal
The user then inputs a number.
The program validates the input to make sure it matches the selected base.
It uses the strconv package to convert between number systems.
The result is printed to the terminal.

# Validation

The program handles invalid inputs without crashing:

Invalid hexadecimal (e.g. 1G) → shows error
Invalid binary (e.g. 10201) → shows error
Invalid decimal (letters included) → shows error
Empty or wrong input → handled safely

# Commands

binary → convert binary to decimal
hexadecimal → convert hex to decimal
decimal → convert decimal to binary and hex
quit → exit the program