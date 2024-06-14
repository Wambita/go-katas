// package techplace use this when submitting the exercise
package main

import (
	"fmt"
	"os"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	name := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + name
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	artline := ""
	for i := 0; i < numStarsPerLine; i++ {
		artline += "*"
	}
	return artline + "\n" + welcomeMsg + "\n" + artline
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	newMsg := strings.TrimSpace(oldMsg)
	return newMsg
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	customer := os.Args[1]
	welcome := WelcomeMessage(customer)
	border := AddBorder(welcome, 8)
	fmt.Println(border)
}
