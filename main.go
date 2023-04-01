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

	var charsToAdd []string
	for i := 0; i < minSpecial; i++ {
		charsToAdd = append(charsToAdd, string(specialChars[rand.Intn(len(specialChars))]))
	}
	for i := 0; i < minNumbers; i++ {
		charsToAdd = append(charsToAdd, string(numbers[rand.Intn(len(numbers))]))
	}
	for i := 0; i < minUpper; i++ {
		charsToAdd = append(charsToAdd, string(upperChars[rand.Intn(len(upperChars))]))
	}
	remain := passLength - minSpecial - minNumbers - minUpper
	for i := 0; i < remain; i++ {
		charsToAdd = append(charsToAdd, string(allChars[rand.Intn(len(allChars))]))
	}

	for i := len(charsToAdd) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		charsToAdd[i], charsToAdd[j] = charsToAdd[j], charsToAdd[i]
	}

	for _, c := range charsToAdd {
		password.WriteString(c)
	}

	return password.String()
}
