package services

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

var (
	validUsername    = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	validPassword    = regexp.MustCompile(`^.{8,}$`)
	validEmail       = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	validURL         = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/?$`)
	validFullName    = regexp.MustCompile(`^[a-zA-Z ]+$`)
	validText        = regexp.MustCompile(`^[a-zA-Z0-9 .,!?-]+$`)
	sanitizeUsername = regexp.MustCompile(`[^a-zA-Z0-9_]`)
	sanitizeUUID     = regexp.MustCompile(`[^a-zA-Z0-9-]`)
)

func IsValidUsername(username string) bool {
	return validUsername.MatchString(username)
}

func IsValidPassword(password string) bool {
	return validPassword.MatchString(password)
}

func IsValidEmail(email string) bool {
	return validEmail.MatchString(email)
}

func IsValidURL(url string) bool {
	return validURL.MatchString(url)
}

func IsValidFullName(fullName string) bool {
	return validFullName.MatchString(fullName)
}

func IsValidTitleOrDescription(text string) bool {
	return validText.MatchString(text)
}

func IsValidUUID(uuidStr string) bool {
	_, err := uuid.Parse(uuidStr)
	return err == nil
}

func SanitizeUsername(input string) (string, error) {
	if !IsValidUsername(input) {
		return "", fmt.Errorf("invalid username")
	}
	return sanitizeUsername.ReplaceAllString(input, ""), nil
}

func SanitizePassword(password string) (string, error) {
	if !IsValidPassword(password) {
		return "", fmt.Errorf("invalid password")
	}
	return password, nil
}

func SanitizeUUID(uuidStr string) (string, error) {
	if !IsValidUUID(uuidStr) {
		return "", fmt.Errorf("invalid UUID")
	}
	return sanitizeUUID.ReplaceAllString(uuidStr, ""), nil
}
