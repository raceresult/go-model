package variant

import (
	"io"
	"time"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"golang.org/x/text/collate"
)

// RDate creates a date variant type.
func RDate(v date.Date) Variant {
	if v.IsZero() {
		v = date.ZeroDateVB
	}
	return rDate(v)
}

// rDate implements a date type.
type rDate date.Date

// Type returns the type of the variant.
func (s rDate) getType() Type {
	return TypeRDate
}

func (s rDate) equals(v Variant, _ bool) bool {
	if v == nil {
		return s.isZero()
	}
	switch val := v.(type) {
	case rBool:
		return s.toBool() == bool(val)
	case rString:
		return s.toString() == string(val)
	}
	return date.Date(s) == v.toDate()
}

func (s rDate) less(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return date.Date(s).Before(date.ZeroDateVB)
	}
	switch val := v.(type) {
	case rString:
		return s.toString() < string(val)
	case rDate:
		return s.toDate().Before(val.toDate())
	}
	return date.Date(s).Before(v.toDate())
}

func (s rDate) greater(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return date.ZeroDateVB.Before(date.Date(s))
	}
	switch val := v.(type) {
	case rString:
		return string(val) < s.toString()
	case rDate:
		return val.toDate().Before(s.toDate())
	}
	return v.toDate().Before(date.Date(s))
}

func (s rDate) toString() string {
	return date.Date(s).String()
}

func (s rDate) toStringWithDateFormat(df string) string {
	if df == "" {
		return date.Date(s).String()
	}
	return date.Date(s).Format(df)
}

func (s rDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.toString() + "\""), nil
}

// ToFloat64 converts the type to float64.
func (s rDate) toFloat64() float64 {
	return float64(date.Date(s).Sub(date.ZeroDateVB))
}

func (s rDate) toDate() date.Date {
	return date.Date(s)
}

func (s rDate) toDateTime() datetime.DateTime {
	return datetime.FromTime(date.Date(s).In(time.UTC), false)
}

func (s rDate) toBool() bool {
	return !s.toDate().IsZero()
}

func (s rDate) toInt() int {
	return int(s.toFloat64())
}

func (s rDate) toDecimal() decimal.Decimal {
	return decimal.FromFloat(s.toFloat64())
}

func (s rDate) isZero() bool {
	return date.Date(s).IsZero()
}

func (s rDate) abs() Variant {
	return s
}
func (s rDate) val() Variant {
	return RFloat(s.toFloat64())
}

func (s rDate) plus(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt:
		return rDate(s.toDate().AddDate(0, 0, p.toInt()))
	case TypeRDecimal:
		return rDateTime(s.toDateTime()).plus(p)
	case TypeRBool:
		return s.plus(rInt(p.toInt()))
	case TypeRDate:
		return nil
	case TypeRDateTime:
		return nil
	case TypeRFloat:
		return rDateTime(s.toDateTime()).plus(p)
	case TypeRString:
		if !p.isNumeric() {
			return nil
		}
		return s.plus(rFloat(p.toFloat64()))
	default:
		return nil
	}
}

func (s rDate) minus(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt:
		return rDate(s.toDate().AddDate(0, 0, -p.toInt()))
	case TypeRDecimal:
		return rDateTime(s.toDateTime()).minus(p)
	case TypeRBool:
		return s.minus(rInt(p.toInt()))
	case TypeRDate:
		return rInt(s.toDate().Sub(p.toDate()))
	case TypeRDateTime:
		return rFloat(s.toDateTime().Sub(p.toDateTime()).Hours() / 24)
	case TypeRFloat:
		return rDateTime(s.toDateTime()).minus(p)
	case TypeRString:
		if !p.isNumeric() {
			return nil
		}
		return s.minus(rFloat(p.toFloat64()))
	default:
		return nil
	}
}

func (s rDate) mult(p Variant) Variant {
	return RFloat(s.toFloat64()).mult(p)
}

func (s rDate) div(p Variant) Variant {
	return RFloat(s.toFloat64()).div(p)
}

func (s rDate) divInt(p Variant) Variant {
	return RFloat(s.toFloat64()).divInt(p)
}

func (s rDate) mod(p Variant) Variant {
	return RFloat(s.toFloat64()).mod(p)
}

func (s rDate) exp(p Variant) Variant {
	return RFloat(s.toFloat64()).exp(p)
}

func (s rDate) isNumeric() bool {
	return true
}

func (s rDate) toJSON(hashDates bool) []byte {
	if s.isZero() {
		return []byte("\"\"")
	}
	if hashDates {
		return []byte("\"#" + s.toDate().String() + "#\"")
	}
	return []byte("\"" + s.toDate().String() + "\"")
}

func (s rDate) writeAsJSON(w io.Writer, hashDates bool) {
	_, _ = w.Write(s.toJSON(hashDates))
}
