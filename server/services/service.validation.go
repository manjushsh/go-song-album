package services

import (
	"regexp"

	"github.com/google/uuid"
)

func IsValidUsername(username string) bool {
	// Define a regular expression for a valid username (alphanumeric characters only)
	var validUsername = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return validUsername.MatchString(username)
}

func IsValidPassword(password string) bool {
	// Define a regular expression for a valid password (at least 8 characters long)
	var validPassword = regexp.MustCompile(`^.{8,}$`)
	return validPassword.MatchString(password)
}

func IsValidEmail(email string) bool {
	// Define a regular expression for a valid email address
	var validEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return validEmail.MatchString(email)
}

func IsValidURL(url string) bool {
	// Define a regular expression for a valid URL
	var validURL = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/?$`)
	return validURL.MatchString(url)
}

func IsValidFullName(fullName string) bool {
	// Define a regular expression for a valid full name (letters only)
	var validFullName = regexp.MustCompile(`^[a-zA-Z ]+$`)
	return validFullName.MatchString(fullName)
}

func IsValidTitleOrDescription(text string) bool {
	// Define a regular expression for a valid title or description (letters, numbers, and punctuation)
	var validText = regexp.MustCompile(`^[a-zA-Z0-9 .,!?-]+$`)
	return validText.MatchString(text)
}

func IsValidUUID(uuidStr string) bool {
	_, err := uuid.Parse(uuidStr)
	return err == nil
}
