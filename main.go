package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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
	minLower := 5

	password := generatePassword(passLength, minSpecial, minNumbers, minUpper, minLower)
	fmt.Println("Your password has been generated, copied to clipboard and stored in password.txt")

	clipboard.WriteAll(password)

	f, err := os.OpenFile("password.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(password + "\n\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func generatePassword(passLength, minSpecial, minNumbers, minUpper, minLower int) string {
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
	for i := 0; i < minLower; i++ {
		charsToAdd = append(charsToAdd, string(lowerChars[rand.Intn(len(lowerChars))]))
	}

	remain := passLength - minSpecial - minNumbers - minUpper - minLower
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
