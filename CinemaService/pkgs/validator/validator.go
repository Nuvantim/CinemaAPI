package validator

import (
	"fmt"
	"regexp"
)

var hmsRegex = regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d:[0-5]\d$`)

func DurationValid(s string) error {
	if !hmsRegex.MatchString(s) {
		return fmt.Errorf("invalid duration format")
	}
	return nil
}
