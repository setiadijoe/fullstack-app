package helpers

import (
	"errors"
	"fmt"
	"strings"
)

// FormatError ...
func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("nickname_already_taken")
	}
	if strings.Contains(err, "email") {
		return errors.New("email_already_taken")
	}
	if strings.Contains(err, "title") {
		return errors.New("title_already_taken")
	}
	if strings.Contains(err, "password") {
		return errors.New("invalid_credentials")
	}
	fmt.Println(err)
	return errors.New("invalid_details")
}
