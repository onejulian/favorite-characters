package util

import (
	"regexp"
)

func IsEmailValid(email string) bool {
	var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if len(email) < 3 || len(email) > 254 {
		return false
	}

	return re.MatchString(email)
}
