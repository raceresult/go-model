package variant

import "strings"

// RList is an interface that can hold any Value List of this package
type RList interface {

	// ToFloat64 converts the values to float64.
	ToFloat64() Float64List

	// ToString converts the values to string.
	ToString() StringList

	// ToStringWithDateFormat converts the values to string with a given date format
	ToStringWithDateFormat(df string) StringList

	// ToDateTime converts the values to datetime.DateTime.
	ToDateTime() DateTimeList

	// ToDate converts the values to date.Date.
	ToDate() DateList

	// ToBool converts the values to bool.
	ToBool() BoolList

	// ToInt converts the values to int.
	ToInt() IntList

	// ToDecimal converts the values to *RDecimal under the hood.
	ToDecimal() DecimalList

	// ToVariant converts the values to Variants
	ToVariant() VariantList

	// Item returns an item of the list
	Item(index int) Variant

	// Len returns length of underlying slice
	Len() int

	// Abs returns the absolute value
	Abs() RList

	// Val converts the values to number
	Val() RList

	// Plus adds a list of values
	Plus(p RList) RList

	// Minus subtracts a list of values
	Minus(p RList) RList

	// Mult multiplies a list of values
	Mult(p RList) RList

	// Div divides a list of values
	Div(p RList) RList

	// Mod returns modulo values
	Mod(p RList) RList

	// Exp calculates a ^ b
	Exp(p RList) RList

	// DivInt returns result of integer devision
	DivInt(p RList) RList
}

// RListArrayToJSON creates a JSON from an RList array
func RListArrayToJSON(result []RList, hashDates bool) []byte {
	R := len(result)
	if R == 0 {
		return []byte("[]")
	}
	L := result[0].Len()

	// build output
	sb := strings.Builder{}
	sb.Grow(L * R * 30)
	sb.WriteByte('[')
	for i := 0; i < L; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteByte('[')
		for j := 0; j < R; j++ {
			if j > 0 {
				sb.WriteByte(',')
			}
			writeAsJSON(result[j].Item(i), &sb, hashDates)
		}
		sb.WriteByte(']')
	}
	sb.WriteByte(']')

	// output result from string builder
	return []byte(sb.String())
}
