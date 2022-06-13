package user

import (
	"fmt"
	"regexp"
)

// Name represents value obejct user name
type Name struct {
	value string
}

func NewName(value string) (*Name, error) {
	if isInvalidPatternName(value) {
		return nil, fmt.Errorf("user name must be a alphanumeric")
	}
	if isInvalidLengthName(value) {
		return nil, fmt.Errorf("user name must be no longer than 50 characters")
	}
	return &Name{value: value}, nil
}

func (n Name) Value() string {
	return n.value
}

func isInvalidPatternName(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(value)
}

func isInvalidLengthName(value string) bool {
	return 50 < len(value)
}
