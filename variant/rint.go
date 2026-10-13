package variant

import (
	"io"
	"math"
	"strconv"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"golang.org/x/text/collate"
)

// RInt creates an integer variant type.
func RInt(v int) Variant {
	return rInt(v)
}

// RInt implements a integer type.
type rInt int

// Type returns the type of the variant.
func (s rInt) getType() Type {
	return TypeRInt
}

func (s rInt) equals(v Variant, _ bool) bool {
	if v == nil {
		return s == 0
	}
	switch val := v.(type) {
	case rBool:
		return s.toBool() == bool(val)
	case rString:
		return s.toString() == string(val)
	case rDateTime:
		return s.toDateTime() == datetime.DateTime(val)
	case rDecimal:
		return s.toDecimal() == decimal.Decimal(val)
	case rFloat:
		return s.toFloat64() == float64(val)
	}
	return int(s) == v.toInt()
}

func (s rInt) less(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return s < 0
	}
	switch val := v.(type) {
	case rString:
		return s.toString() < string(val)
	case rDateTime:
		return s.toDateTime().Before(datetime.DateTime(val))
	case rDecimal:
		return s.toDecimal() < decimal.Decimal(val)
	case rFloat:
		return s.toFloat64() < float64(val)
	case rInt:
		return int(s) < int(val)
	}
	return int(s) < v.toInt()
}

func (s rInt) greater(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return 0 < s
	}
	switch val := v.(type) {
	case rString:
		return s.toString() > string(val)
	case rDateTime:
		return datetime.DateTime(val).Before(s.toDateTime())
	case rDecimal:
		return decimal.Decimal(val) < s.toDecimal()
	case rFloat:
		return float64(val) < s.toFloat64()
	case rInt:
		return int(val) < int(s)
	}
	return v.toInt() < int(s)
}

// ToFloat64 converts the type to float64.
func (s rInt) toFloat64() float64 {
	return float64(s)
}

func (s rInt) toString() string {
	return strconv.Itoa(int(s))
}
func (s rInt) toStringWithDateFormat(string) string {
	return s.toString()
}

func (s rInt) toDateTime() datetime.DateTime {
	return datetime.ZeroDate().AddDate(0, 0, int(s))
}

func (s rInt) toDate() date.Date {
	return date.ZeroDateVB.AddDate(0, 0, s.toInt())
}

func (s rInt) toBool() bool {
	return s != 0
}

func (s rInt) toInt() int {
	return int(s)
}

func (s rInt) toDecimal() decimal.Decimal {
	return decimal.FromInt(int(s))
}

func (s rInt) abs() Variant {
	if s < 0 {
		return -s
	}
	return s
}
func (s rInt) val() Variant {
	return s
}

func (s rInt) plus(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt, TypeRBool:
		return rInt(int(s) + p.toInt())
	case TypeRDecimal:
		return rDecimal(s.toDecimal()).plus(p)
	case TypeRDateTime, TypeRDate, TypeRFloat, TypeRString:
		return rFloat(s.toFloat64()).plus(p)
	default:
		return nil
	}
}

func (s rInt) minus(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt, TypeRBool:
		return rInt(int(s) - p.toInt())
	case TypeRDecimal:
		return rDecimal(s.toDecimal()).minus(p)
	case TypeRDateTime, TypeRDate, TypeRFloat, TypeRString:
		return rFloat(s.toFloat64()).minus(p)
	default:
		return nil
	}
}

func (s rInt) mult(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt, TypeRBool:
		return rInt(int(s) * p.toInt())
	case TypeRDecimal:
		return rDecimal(s.toDecimal()).mult(p)
	case TypeRDateTime, TypeRDate, TypeRFloat, TypeRString:
		return rFloat(s.toFloat64()).mult(p)
	default:
		return nil
	}
}

func (s rInt) div(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt, TypeRBool:
		x := p.toInt()
		if x == 0 {
			return nil
		}
		return rFloat(float64(s) / float64(x))
	case TypeRDecimal:
		return rDecimal(s.toDecimal()).div(p)
	case TypeRDateTime, TypeRDate, TypeRFloat, TypeRString:
		return rFloat(s.toFloat64()).div(p)
	default:
		return nil
	}
}

func (s rInt) divInt(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt, TypeRBool:
		x := p.toInt()
		if x == 0 {
			return nil
		}
		return rInt(int(s) / x)
	case TypeRDecimal:
		return rDecimal(s.toDecimal()).divInt(p)
	case TypeRDateTime, TypeRDate, TypeRFloat, TypeRString:
		return rFloat(s.toFloat64()).divInt(p)
	default:
		return nil
	}
}

func (s rInt) mod(p Variant) Variant {
	x := p.toInt()
	if x == 0 {
		return nil
	}
	return rInt(int(s) % x)
}

func (s rInt) exp(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt, TypeRBool:
		return RInt(int(math.Pow(s.toFloat64(), p.toFloat64())))
	case TypeRDecimal:
		return RFloat(s.toFloat64()).exp(p)
	case TypeRDateTime, TypeRDate, TypeRFloat:
		return rFloat(s.toFloat64()).exp(p)
	case TypeRString:
		n, err := ParseNumber(ToString(p))
		if err != nil {
			return nil
		}
		return s.exp(n)
	default:
		return nil
	}
}

func (s rInt) isNumeric() bool {
	return true
}

func (s rInt) toJSON(_ bool) []byte {
	return []byte(s.toString())
}

func (s rInt) writeAsJSON(w io.Writer, hashDates bool) {
	_, _ = w.Write(s.toJSON(hashDates))
}
