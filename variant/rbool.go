package variant

import (
	"io"
	"time"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"golang.org/x/text/collate"
)

// RBool creates a boolean variant type.
func RBool(v bool) Variant {
	return rBool(v)
}

// RBool implements a boolean type.
type rBool bool

// Type returns the type of the variant.
func (s rBool) getType() Type {
	return TypeRBool
}

func (s rBool) equals(v Variant, _ bool) bool {
	if v == nil {
		return bool(s)
	}
	return bool(s) == v.toBool()
}

func (s rBool) less(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return false
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
		return s.toInt() < int(val)
	case rBool:
		return bool(s) && !bool(val)
	}
	return s.toInt() < v.toInt()
}

func (s rBool) greater(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return bool(s)
	}
	switch val := v.(type) {
	case rString:
		return string(val) < s.toString()
	case rDateTime:
		return datetime.DateTime(val).Before(s.toDateTime())
	case rDecimal:
		return decimal.Decimal(val) < s.toDecimal()
	case rFloat:
		return float64(val) < s.toFloat64()
	case rInt:
		return int(val) < s.toInt()
	case rBool:
		return !bool(s) && bool(val)
	}
	return v.toInt() < s.toInt()
}

// ToFloat64 converts the type to float64.
func (s rBool) toFloat64() float64 {
	if s {
		return 1
	}
	return 0
}

func (s rBool) toString() string {
	if s {
		return "1"
	}
	return "0"
}

func (s rBool) toStringWithDateFormat(string) string {
	return s.toString()
}

func (s rBool) toDateTime() datetime.DateTime {
	if s {
		return datetime.ZeroDate().Add(24 * time.Hour)
	}
	return datetime.ZeroDate()
}

func (s rBool) toDate() date.Date {
	if s {
		return date.ZeroDateVB.AddDate(0, 0, 1)
	}
	return date.ZeroDateVB
}

func (s rBool) toBool() bool {
	return bool(s)
}

func (s rBool) toInt() int {
	if s {
		return 1
	}
	return 0
}

func (s rBool) toDecimal() decimal.Decimal {
	return decimal.FromInt(s.toInt())
}

func (s rBool) abs() Variant {
	return s
}

func (s rBool) val() Variant {
	return RInt(0)
}

func (s rBool) plus(p Variant) Variant {
	return rInt(s.toInt()).plus(p)
}

func (s rBool) minus(p Variant) Variant {
	return rInt(s.toInt()).minus(p)
}

func (s rBool) mult(p Variant) Variant {
	return rInt(s.toInt()).mult(p)
}

func (s rBool) div(p Variant) Variant {
	return rInt(s.toInt()).div(p)
}

func (s rBool) divInt(p Variant) Variant {
	return rInt(s.toInt()).divInt(p)
}

func (s rBool) mod(p Variant) Variant {
	return rInt(s.toInt()).mod(p)
}

func (s rBool) exp(p Variant) Variant {
	return rInt(s.toInt()).exp(p)
}

func (s rBool) isNumeric() bool {
	return true
}

func (s rBool) toJSON(_ bool) []byte {
	if s {
		return []byte("true")
	}
	return []byte("false")
}

func (s rBool) writeAsJSON(w io.Writer, hashDates bool) {
	_, _ = w.Write(s.toJSON(hashDates))
}
