package modules

import (
	"errors"
	"regexp"
	"strings"
)

func CleanPhoneNumber(p string) (string, error) {
	// Remove all white spaces
	phone := strings.ReplaceAll(p, " ", "")

	// Validate the number contains only digits and is within the expected length
	if !regexp.MustCompile(`^\d{10,15}$`).MatchString(phone) {
		return "", errors.New("invalid phone number")
	}

	// Switch the cases of the number
	switch {
	case len(phone) == 13 && strings.HasPrefix(phone, "9647"):
		return "+" + phone, nil
	case len(phone) == 14 && strings.HasPrefix(phone, "+9647"):
		return phone, nil
	case len(phone) == 15 && strings.HasPrefix(phone, "009647"):
		return "+" + phone[2:], nil
	case len(phone) == 10 && (strings.HasPrefix(phone, "78") || strings.HasPrefix(phone, "79") || strings.HasPrefix(phone, "77") || strings.HasPrefix(phone, "75")):
		return "+964" + phone, nil
	case len(phone) == 11 && strings.HasPrefix(phone, "07"):
		return "+964" + phone[1:], nil
	default:
		return "", errors.New("invalid phone number")
	}
}
