package variant

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/raceresult/go-model/date"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/variant/jsonstring"
	"golang.org/x/text/collate"
)

// RString creates a string variant type.
func RString(v string) Variant {
	return rString(v)
}

// RString implements a string type.
type rString string

// Type returns the type of the variant.
func (s rString) getType() Type {
	return TypeRString
}

func (s rString) equals(v Variant, caseSensitive bool) bool {
	if v == nil {
		return s == ""
	}
	if val, ok := v.(rBool); ok {
		return s.toBool() == bool(val)
	}
	if caseSensitive {
		return string(s) == v.toString()
	}
	return strings.EqualFold(string(s), v.toString())
}

func (s rString) less(v Variant, collator *collate.Collator) bool {
	if v == nil {
		return false
	}
	if collator == nil || !IsString(v) {
		return string(s) < v.toString()
	}

	r := collator.CompareString(string(s), v.toString())
	if r < 0 {
		return true
	} else if r > 0 {
		return false
	}

	return string(s) < v.toString()
}

func (s rString) greater(v Variant, collator *collate.Collator) bool {
	if v == nil {
		return s != ""
	}
	if collator == nil || !IsString(v) {
		return string(s) > v.toString()
	}

	r := collator.CompareString(string(s), v.toString())
	if r > 0 {
		return true
	} else if r < 0 {
		return false
	}

	return string(s) > v.toString()
}

// ToFloat64 converts the type to float64.
func (s rString) toFloat64() float64 {
	v, err := strconv.ParseFloat(string(s), 64)
	if err == nil {
		return v
	}

	v, err = strconv.ParseFloat(strings.ReplaceAll(string(s), ",", "."), 64)
	if err == nil {
		return v
	}

	return 0
}

func (s rString) toString() string {
	return string(s)
}
func (s rString) toStringWithDateFormat(string) string {
	return s.toString()
}

func (s rString) toDateTime() datetime.DateTime {
	if d, ok := datetime.Parse(string(s)); ok {
		return d
	}
	return datetime.ZeroDate()
}

func (s rString) toDate() date.Date {
	s2 := string(s)
	if len(s2) > 10 {
		s2 = s2[:10]
	}
	if d, err := date.AutoParse(s2); err == nil {
		return d
	}
	return date.ZeroDateVB
}

func (s rString) toBool() bool {
	// strconv.ParseBool exists, but would change the behaviour as it uses different cases to turn to `true`.
	if s == "1" || s == "-1" {
		return true
	}
	if len(s) > 4 {
		return false
	}
	uc := strings.ToUpper(string(s))
	if uc == "TRUE" || uc == "YES" || uc == "JA" || uc == "WAHR" {
		return true
	}
	return false
}

func (s rString) toInt() int {
	v, err := strconv.Atoi(string(s))
	if err != nil {
		return 0
	}
	return v
}

func (s rString) toDecimal() decimal.Decimal {
	v, err := decimal.FromString(string(s))
	if err == nil {
		return v
	}

	v, err = decimal.FromString(strings.ReplaceAll(string(s), ",", "."))
	if err == nil {
		return v
	}

	return 0
}

func (s rString) abs() Variant {
	n, err := ParseNumber(string(s))
	if err != nil {
		return nil
	}
	return n.abs()
}
func (s rString) val() Variant {
	x := decimal.Val(string(s))
	if x.IsInt() {
		return RInt(x.ToInt())
	}
	return RDecimal(x)
}

func (s rString) plus(p Variant) Variant {
	n, err := ParseNumber(string(s))
	if err != nil {
		return nil
	}
	return n.plus(p)
}

func (s rString) minus(p Variant) Variant {
	n, err := ParseNumber(string(s))
	if err != nil {
		return nil
	}
	return n.minus(p)
}

func (s rString) mult(p Variant) Variant {
	n, err := ParseNumber(string(s))
	if err != nil {
		return nil
	}
	return n.mult(p)
}

func (s rString) div(p Variant) Variant {
	n, err := ParseNumber(string(s))
	if err != nil {
		return nil
	}
	return n.div(p)
}

func (s rString) divInt(p Variant) Variant {
	n, err := ParseNumber(string(s))
	if err != nil {
		return nil
	}
	return n.divInt(p)
}

func (s rString) mod(p Variant) Variant {
	n, err := ParseNumber(string(s))
	if err != nil {
		return nil
	}
	return n.mod(p)
}

func (s rString) exp(p Variant) Variant {
	n, err := ParseNumber(string(s))
	if err != nil {
		return nil
	}
	return n.exp(p)
}

func (s rString) isNumeric() bool {
	_, err := ParseNumber(string(s))
	return err == nil
}

func (s rString) toJSON(_ bool) []byte {
	bb, _ := json.Marshal(string(s))
	return bb
}

func (s rString) writeAsJSON(w io.Writer, _ bool) {
	jsonstring.WriteJSONString(string(s), w)
}
