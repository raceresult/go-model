package variant

import (
	"io"
	"time"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"golang.org/x/text/collate"
)

// RDateTime creates a date variant type.
func RDateTime(v datetime.DateTime) Variant {
	if v.IsZero() {
		v = datetime.ZeroDate()
	}
	return rDateTime(v)
}

// rDateTime implements a date type.
type rDateTime datetime.DateTime

// Type returns the type of the variant.
func (s rDateTime) getType() Type {
	return TypeRDateTime
}

func (s rDateTime) equals(v Variant, _ bool) bool {
	if v == nil {
		return s.isZero()
	}
	switch val := v.(type) {
	case rBool:
		return s.toBool() == bool(val)
	case rString:
		return s.toString() == string(val)
	}
	return datetime.DateTime(s) == v.toDateTime()
}

func (s rDateTime) less(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return datetime.DateTime(s).Before(datetime.ZeroDate())
	}
	switch val := v.(type) {
	case rString:
		return s.toString() < string(val)
	case rDateTime:
		return s.toDateTime().Before(val.toDateTime())
	}
	return datetime.DateTime(s).Before(v.toDateTime())
}

func (s rDateTime) greater(v Variant, _ *collate.Collator) bool {
	if v == nil {
		return datetime.ZeroDate().Before(datetime.DateTime(s))
	}
	switch val := v.(type) {
	case rString:
		return string(val) < s.toString()
	case rDateTime:
		return val.toDateTime().Before(s.toDateTime())
	}
	return v.toDateTime().Before(datetime.DateTime(s))
}

func (s rDateTime) toString() string {
	return datetime.DateTime(s).ToString()
}

func (s rDateTime) toStringWithDateFormat(df string) string {
	return datetime.DateTime(s).ToStringWithDateFormat(df)
}

func (s rDateTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.toString() + "\""), nil
}

// ToFloat64 converts the type to float64.
func (s rDateTime) toFloat64() float64 {
	return datetime.DateTime(s).Sub(datetime.ZeroDate()).Hours() / 24
}

func (s rDateTime) toDateTime() datetime.DateTime {
	return datetime.DateTime(s)
}

func (s rDateTime) toDate() date.Date {
	return date.NewAt(datetime.DateTime(s).ToTime())
}

func (s rDateTime) toBool() bool {
	return !s.toDateTime().IsZero()
}

func (s rDateTime) toInt() int {
	return int(s.toFloat64())
}

func (s rDateTime) toDecimal() decimal.Decimal {
	return decimal.FromFloat(s.toFloat64())
}

func (s rDateTime) isZero() bool {
	return datetime.DateTime(s).IsZero()
}

func (s rDateTime) abs() Variant {
	return s
}
func (s rDateTime) val() Variant {
	return RFloat(s.toFloat64())
}

func (s rDateTime) plus(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt:
		return rDateTime(s.toDateTime().AddDate(0, 0, p.toInt()))
	case TypeRDecimal:
		return rDateTime(s.toDateTime().Add(time.Duration(p.toDecimal() * 100 * 1000)))
	case TypeRBool:
		return rDateTime(s.toDateTime().AddDate(0, 0, p.toInt()))
	case TypeRDateTime:
		return nil
	case TypeRDate:
		return nil
	case TypeRFloat:
		return rDateTime(s.toDateTime().Add(time.Duration(p.toFloat64() * 86400 * 1000 * 1000 * 1000)))
	case TypeRString:
		if !p.isNumeric() {
			return nil
		}
		return rDateTime(s.toDateTime().Add(time.Duration(p.toFloat64() * 86400 * 1000 * 1000 * 1000)))
	default:
		return nil
	}
}

func (s rDateTime) minus(p Variant) Variant {
	switch GetType(p) {
	case TypeRInt:
		return rDateTime(s.toDateTime().AddDate(0, 0, -p.toInt()))
	case TypeRDecimal:
		return rDateTime(s.toDateTime().Add(time.Duration(-p.toDecimal() * 100 * 1000)))
	case TypeRBool:
		return rDateTime(s.toDateTime().AddDate(0, 0, -p.toInt()))
	case TypeRDateTime:
		return rFloat(s.toDateTime().Sub(p.toDateTime()).Hours() / 24)
	case TypeRDate:
		return rFloat(s.toDateTime().Sub(p.toDateTime()).Hours() / 24)
	case TypeRFloat:
		return rDateTime(s.toDateTime().Add(time.Duration(-p.toFloat64() * 86400 * 1000 * 1000 * 1000)))
	case TypeRString:
		if !p.isNumeric() {
			return nil
		}
		return rDateTime(s.toDateTime().Add(time.Duration(-p.toFloat64() * 86400 * 1000 * 1000 * 1000)))
	default:
		return nil
	}
}

func (s rDateTime) mult(p Variant) Variant {
	return RFloat(s.toFloat64()).mult(p)
}

func (s rDateTime) div(p Variant) Variant {
	return RFloat(s.toFloat64()).div(p)
}

func (s rDateTime) divInt(p Variant) Variant {
	return RFloat(s.toFloat64()).divInt(p)
}

func (s rDateTime) mod(p Variant) Variant {
	return RFloat(s.toFloat64()).mod(p)
}

func (s rDateTime) exp(p Variant) Variant {
	return RFloat(s.toFloat64()).exp(p)
}

func (s rDateTime) isNumeric() bool {
	return true
}

func (s rDateTime) toJSON(hashDates bool) []byte {
	if s.isZero() {
		return []byte("\"\"")
	}
	if hashDates {
		return []byte("\"#" + s.toDateTime().ToString() + "#\"")
	}
	return []byte("\"" + s.toDateTime().ToString() + "\"")
}

func (s rDateTime) writeAsJSON(w io.Writer, hashDates bool) {
	_, _ = w.Write(s.toJSON(hashDates))
}
