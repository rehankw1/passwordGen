package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/atotto/clipboard"
)

var (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChars = "!@#$%^&*"
	numbers      = "11234567890"
	allChars     = lowerChars + upperChars + specialChars + numbers
)

func main() {

	passLength := 23
	minSpecial := 5
	minUpper := 5
	minNumbers := 6

	password := generatePassword(passLength, minSpecial, minNumbers, minUpper)
	fmt.Println("Your password has been generated and copied to clipboard")

	clipboard.WriteAll(password)

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
