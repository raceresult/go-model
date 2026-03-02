package decimal

import (
	"math"
	"strings"
	"time"
)

// Decimals is the conversion factor for the fixed decimals number
const Decimals = 10000

// Min is the smallest possible Decimal value
const Min = Decimal(math.MinInt64)

// Max is the biggest possible Decimal value
const Max = Decimal(math.MaxInt64)

// FromFloat creates a new decimal from float
func FromFloat(f float64) Decimal {
	return Decimal(math.RoundToEven(f * Decimals))
}

// FromInt creates a Decimal from an int
func FromInt(i int) Decimal {
	return Decimal(i * Decimals)
}

// FromInt64 creates a Decimal from an int64
func FromInt64(i int64) Decimal {
	return Decimal(i * Decimals)
}

// FromDuration creates a time from Duration
func FromDuration(d time.Duration) Decimal {
	return Decimal(int64(d) / (int64(time.Second) / Decimals))
}

// Val parses a string to Decimal, ignores errors
func Val(s string) Decimal {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	hasDecimals := false
	var end int
loop:
	for end = 0; end < len(s); end++ {
		switch s[end] {
		case '-':
			if end > 0 {
				break loop
			}
		case '+':
			if end > 0 {
				break loop
			}
			return Val(s[end+1:])
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		// ok
		case '.':
			if hasDecimals {
				break loop
			} else {
				hasDecimals = true
			}
		case ',':
			if hasDecimals {
				break loop
			} else {
				b := []byte(s)
				b[end] = '.'
				s = string(b)
				hasDecimals = true
			}
		default:
			break loop
		}
	}
	if end > 1 && s[end-1] == '.' {
		end--
	}
	x, _ := FromString(s[:end])
	return x
}

// Decimal implements a number with 4 Decimals.
type Decimal int64

// ToFloat64 converts the type to float64.
func (s Decimal) ToFloat64() float64 {
	return float64(s) / Decimals
}

// ToInt converts the type to int.
func (s Decimal) ToInt() int {
	return int(int64(s) / Decimals)
}

// ToDuration converts the Decimal to a duration
func (s Decimal) ToDuration() time.Duration {
	return time.Second / Decimals * time.Duration(s)
}

// Round returns the Decimal rounded according to the given number of decimals
func (s Decimal) Round(numberOfDecimals int) Decimal {
	if s < 0 {
		if s == Min { // endless loop without this check
			return s
		}
		return -((-s).Round(numberOfDecimals))
	}
	var x Decimal
	var a Decimal
	switch numberOfDecimals {
	case 0:
		x = Decimals
		a = 5000
	case 1:
		x = Decimals / 10
		a = 500
	case 2:
		x = Decimals / 100
		a = 50
	case 3:
		x = Decimals / 1000
		a = 5
	default:
		x = Decimals / 10000
		a = 0
	}

	c := (s + a) / x
	c *= x
	return c
}

// RoundUp returns the Decimal rounded up
func (s Decimal) RoundUp(numberOfDecimals int) Decimal {
	if s < 0 {
		if s == Max { // endless loop without this check
			return s
		}
		return -((-s).RoundUp(numberOfDecimals))
	}

	d := Decimal(Decimals)
	for i := 0; i < numberOfDecimals; i++ {
		d /= 10
	}
	return s + (d-(s%d))%d
}

// Mult multiplies the Decimal with another Decimal
func (s Decimal) Mult(d Decimal) Decimal {
	if d.IsInt() { // to avoid overflow for large numbers
		return s.MultInt(d.ToInt())
	}
	return s * d / Decimals
}

// MultInt multiplies the decimal with an int and returns a new decimal
func (s Decimal) MultInt(i int) Decimal {
	return s * Decimal(i)
}

// MultFloat multiplies the decimal with a float and returns a new decimal
func (s Decimal) MultFloat(f float64) Decimal {
	if f == 1 {
		return s
	}
	return FromFloat(s.ToFloat64() * f)
}

// Div divides the Decimal by a float and returns a float
func (s Decimal) Div(f float64) float64 {
	return s.ToFloat64() / f
}

// DivDecimal divides the Decimal by another Decimal and returns a Decimal
func (s Decimal) DivDecimal(d Decimal) Decimal {
	// todo: rounding for exact results as old SES, may be removed later again
	x := (10 * s * Decimals / d) % 10
	if x >= 5 {
		return s*Decimals/d + 1
	}

	return s * Decimals / d
}

// EqualInt checks if the number is equal to the given int
func (s Decimal) EqualInt(i int) bool {
	return s == FromInt(i)
}

func (s Decimal) IsInt() bool {
	return s%Decimals == 0
}

func GetMax(d1 Decimal, d2 Decimal) Decimal {
	if d1 > d2 {
		return d1
	}
	return d2
}
