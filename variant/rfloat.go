package variant

import (
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"golang.org/x/text/collate"
)

// RFloat creates a floating point variant type.
func RFloat(v float64) Variant {
	return rFloat(v)
}

// RFloat implements the a float type.
type rFloat float64

// Type returns the type of the variant.
func (s rFloat) getType() Type {
	return TypeRFloat
}

func (s rFloat) equals(v Variant, _ bool) bool {
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
	}
	return float64(s) == v.toFloat64()
}

func (s rFloat) less(v Variant, _ *collate.Collator) bool {
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
		return float64(s) < float64(val)
	}
	return float64(s) < v.toFloat64()
}

func (s rFloat) greater(v Variant, _ *collate.Collator) bool {
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
		return float64(val) < float64(s)
	}
	return v.toFloat64() < float64(s)
}

func (s rFloat) toFloat64() float64 {
	return float64(s)
}

func (s rFloat) toString() string {
	v := strconv.FormatFloat(float64(s), 'f', -1, 64)
	t := strings.IndexByte(v, '.')
	if t > 0 && t < len(v)-10 {
		return strings.TrimRight(strings.TrimRight(strconv.FormatFloat(float64(s), 'f', 8, 64), "0"), ".")
	}
	return v
}
func (s rFloat) toStringWithDateFormat(string) string {
	return s.toString()
}

func (s rFloat) toDateTime() datetime.DateTime {
	return toTime(float64(s))
}

func (s rFloat) toDate() date.Date {
	return date.ZeroDateVB.AddDate(0, 0, s.toInt())
}

func toTime(f float64) datetime.DateTime {
	days := int(f)
	hours := time.Duration((f - float64(days)) * float64(time.Hour) * 24)
	return datetime.ZeroDate().AddDate(0, 0, days).Add(hours)
}

func (s rFloat) toBool() bool {
	return s != 0
}

func (s rFloat) toInt() int {
	return int(s)
}

func (s rFloat) toDecimal() decimal.Decimal {
	return decimal.FromFloat(float64(s))
}

func (s rFloat) abs() Variant {
	if s < 0 {
		return -s
	}
	return s
}
func (s rFloat) val() Variant {
	return s
}
func (s rFloat) plus(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	return rFloat(s.toFloat64() + p.toFloat64())
}

func (s rFloat) minus(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	return rFloat(s.toFloat64() - p.toFloat64())
}

func (s rFloat) mult(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	return rFloat(s.toFloat64() * p.toFloat64())
}

func (s rFloat) div(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	x := p.toFloat64()
	if x == 0 {
		return nil
	}
	return rFloat(s.toFloat64() / x)
}

func (s rFloat) divInt(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	x := p.toFloat64()
	if x == 0 {
		return nil
	}
	return rInt(s.toFloat64() / x)
}

func (s rFloat) mod(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	x := p.toInt()
	if x == 0 {
		return nil
	}
	k := int(s)
	z := float64(s) - float64(k)
	k %= x
	return rFloat(float64(k) + z)
}

func (s rFloat) exp(p Variant) Variant {
	if p == nil || !p.isNumeric() {
		return nil
	}
	e := p.toFloat64()
	if s < 0 && float64(int(e)) != e {
		return nil
	}
	return RFloat(math.Pow(s.toFloat64(), e))
}

func (s rFloat) isNumeric() bool {
	return true
}

func (s rFloat) toJSON(_ bool) []byte {
	return []byte(s.toString())
}

func (s rFloat) writeAsJSON(w io.Writer, hashDates bool) {
	_, _ = w.Write(s.toJSON(hashDates))
}
