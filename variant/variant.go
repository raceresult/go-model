package variant

import (
	"io"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"golang.org/x/text/collate"
)

// Variant implements a variable type similar to VB Variant but limited to the needed types.
type Variant interface {
	// Type returns the type of the variant.
	getType() Type

	// Equals checks if this value equals the given value.
	// String comparison is used if one of the values is a string,
	// number comparison otherwise using float64.
	equals(v Variant, caseSensitive bool) bool

	// Less checks if this value is less than the given value.
	// If one of the values is a string, this will use a string comparison,
	// otherwise a number comparison using float64.
	less(v Variant, collator *collate.Collator) bool

	// Greater checks if this value is greater than the given value.
	// If one of the values is a string, this will use a string comparison,
	// otherwise a number comparison using float64.
	greater(v Variant, collator *collate.Collator) bool

	// ToFloat64 converts the value to float64.
	toFloat64() float64

	// ToString converts the value to string.
	toString() string

	// ToString converts the value to string with a given date format
	toStringWithDateFormat(df string) string

	// ToDateTime converts the value to DateTime.
	toDateTime() datetime.DateTime

	// ToDate converts the value to Date.
	toDate() date.Date

	// ToBool converts the value to bool.
	toBool() bool

	// ToInt converts the value to int.
	toInt() int

	// FromFloat converts the value to *RDecimal under the hood.
	toDecimal() decimal.Decimal

	// absolute value
	abs() Variant

	// val
	val() Variant

	// plus
	plus(p Variant) Variant

	// minus
	minus(p Variant) Variant

	// mult
	mult(p Variant) Variant

	// div
	div(p Variant) Variant

	// mod
	mod(p Variant) Variant

	// div int
	divInt(p Variant) Variant

	// exp
	exp(p Variant) Variant

	isNumeric() bool

	toJSON(hashDates bool) []byte

	writeAsJSON(w io.Writer, hashDates bool)
}

// ToVariant converts interface{} to Variant. Use with extreme care.
// Most of the time you are doing it wrong if you need this function.
// If the value is not an int, string, bool, or float64, nil will be returned.
func ToVariant(i interface{}) Variant {
	return ToVariant2(i, false)
}

func ToVariant2(i interface{}, datesHashed bool) Variant {
	switch v := i.(type) {
	case int:
		return RInt(v)
	case string:
		l := len(v)
		if datesHashed {
			if (l == 12) && v[5] == '-' && v[8] == '-' {
				if d, err := date.ParseISO(v[1 : l-1]); err == nil {
					return RDate(d)
				}

			} else if (l == 27 || l == 22 || l == 21) && v[5] == '-' && v[8] == '-' {
				if d, ok := datetime.Parse(v[1 : l-1]); ok {
					return RDateTime(d)
				}
			}
		} else {
			if (l == 10) && v[4] == '-' && v[7] == '-' {
				if d, err := date.ParseISO(v); err == nil {
					return RDate(d)
				}
			} else if (l == 25 || l == 20 || l == 19) && v[4] == '-' && v[7] == '-' {
				if d, ok := datetime.Parse(v); ok {
					return RDateTime(d)
				}
			}
		}
		return RString(v)
	case bool:
		return RBool(v)
	case float64:
		return RFloat(v)
	default:
		return nil
	}
}

// ToInterface returns the underlying string/int/bool/datetime/decimal value
func ToInterface(v Variant) interface{} {
	if v == nil {
		return nil
	}
	switch v.getType() {
	case TypeRBool:
		return v.toBool()
	case TypeRString:
		return v.toString()
	case TypeRInt:
		return v.toInt()
	case TypeRFloat:
		return v.toFloat64()
	case TypeRDecimal:
		return v.toDecimal()
	case TypeRDateTime:
		return v.toDateTime()
	case TypeRDate:
		return v.toDate()
	default:
		panic("new variant type not implemented")
	}
}

// ToJSON converts the value to a JSON string
func ToJSON(v Variant, hashDates bool) []byte {
	if v == nil {
		return []byte("\"\"")
	}
	return v.toJSON(hashDates)
}

// ToJSON converts the value to a JSON string
func writeAsJSON(v Variant, w io.Writer, hashDates bool) {
	if v == nil {
		_, _ = w.Write([]byte("\"\""))
	} else {
		v.writeAsJSON(w, hashDates)
	}
}

// Val parses the value to a number
func Val(v Variant) Variant {
	if IsEmpty(v) {
		return RInt(0)
	}
	return v.val()
}
