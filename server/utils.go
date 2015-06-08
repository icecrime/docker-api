package server

import (
	"fmt"
	"strings"
)

var allowedBooleanValues = map[string]bool{
	"0":     false,
	"false": false,
	"1":     true,
	"true":  true,
}

const ErrInvalidBooleanValue = "invalid boolean value %q"

func booleanValue(s string, defaultValue bool) (bool, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return defaultValue, nil
	}
	if b, ok := allowedBooleanValues[s]; ok {
		return b, nil
	}
	return false, fmt.Errorf(ErrInvalidBooleanValue, s)
}
