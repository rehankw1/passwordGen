package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChars = "!@#$%^&*"
	numbers      = "11234567890"
	allChars     = lowerChars + upperChars + specialChars + numbers
)

func main() {

	var passLength int
	var minSpecial int
	var minUpper int
	var minNumbers int

	fmt.Println("Enter the password length")
	fmt.Scan(&passLength)
	fmt.Println("Enter the minumum number of special characters")
	fmt.Scan(&minSpecial)
	fmt.Println("Enter the minimum number of uppercase characters")
	fmt.Scan(&minUpper)
	fmt.Println("Enter the minimum number of numbers")
	fmt.Scan(&minNumbers)

	if minSpecial+minUpper+minNumbers > passLength {
		fmt.Println("Password length is too short for the amount of characters specified")
		os.Exit(1)
	}

	password := generatePassword(passLength, minSpecial, minNumbers, minUpper)
	fmt.Println("Your generated password is:", password)
}

func generatePassword(passLength, minSpecial, minNumbers, minUpper int) string {
	var password strings.Builder

	for i := 0; i < minSpecial; i++ {
		r := rand.Intn(len(specialChars))
		password.WriteString(string(specialChars[r]))
	}

	for i := 0; i < minNumbers; i++ {
		r := rand.Intn(len(numbers))
		password.WriteString(string(numbers[r]))
	}

	for i := 0; i < minUpper; i++ {
		r := rand.Intn(len(upperChars))
		password.WriteString(string(upperChars[r]))
	}

	remain := passLength - minSpecial - minNumbers - minUpper
	for i := 0; i < remain; i++ {
		r := rand.Intn(len(allChars))
		password.WriteString(string(allChars[r]))
	}

	return password.String()
}
