package main

import "strings"

func validateUserInput(userName string, email string) (bool, bool) {
	isValidName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	return isValidEmail, isValidName
}