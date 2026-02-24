package decimal

import "strings"

// MarshalJSON converts the number to JSON
func (s Decimal) MarshalJSON() ([]byte, error) {
	return []byte(s.ToString()), nil
}

// UnmarshalJSON parses a Decimal from JSON
func (s *Decimal) UnmarshalJSON(data []byte) error {
	// UnmarshalJSON is called with quotes in the byte slice when unmarshalling a map key
	trimmed := strings.Trim(string(data), "\"")
	x, err := FromString(trimmed)
	if err != nil {
		return err
	}
	*s = x
	return nil
}
