package decimal

// MarshalText converts the number to text
func (s Decimal) MarshalText() ([]byte, error) {
	return []byte(s.ToString()), nil
}

// UnmarshalText parses a Decimal from Text
func (s *Decimal) UnmarshalText(data []byte) error {
	x, err := FromString(string(data))
	if err != nil {
		return err
	}
	*s = x
	return nil
}
