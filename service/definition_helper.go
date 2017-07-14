package service

import "strings"

// IsValid returns a boolean indicating if the Definition's fields meet the
// minimum requirements.
func IsValid(d Definition) bool {
	if InvalidReason(d) == nil {
		return true
	}
	return false
}

// InvalidReason returns an error explaining what makes the Definition invalid,
// or nil if the Definition is valid.
func InvalidReason(d Definition) error {
	name := strings.TrimSpace(d.Name)
	if name == "" {
		return ErrDefEmptyName
	}
	if d.Init == nil {
		return ErrDefNilInit
	}
	return nil
}
