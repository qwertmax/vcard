package vcard

import "fmt"

// Birthday type
type Birthday struct {
	Year, Month, Day int
}

func (b Birthday) String() string {
	return fmt.Sprintf("BDAY: %d%d%d\n", b.Year, b.Month, b.Day)
}

func (b Birthday) valid() bool {
	if b.Year < 1900 {
		return false
	}

	if b.Year > 9999 {
		return false
	}

	if b.Month <= 0 {
		return false
	}

	if b.Month > 12 {
		return false
	}

	if b.Day <= 0 {
		return false
	}

	if b.Day > 31 {
		return false
	}

	return true
}
