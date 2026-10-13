package jsonstring

import (
	"io"
	"unicode/utf8"
)

const hexDigits = "0123456789abcdef"

// ctrlEscapes holds the JSON escape sequence for each control character below 0x20.
var ctrlEscapes [0x20]string

func init() {
	for c := 0; c < 0x20; c++ {
		ctrlEscapes[c] = `\u00` + string(hexDigits[c>>4]) + string(hexDigits[c&0xF])
	}
	ctrlEscapes['\b'] = `\b`
	ctrlEscapes['\f'] = `\f`
	ctrlEscapes['\n'] = `\n`
	ctrlEscapes['\r'] = `\r`
	ctrlEscapes['\t'] = `\t`
}

// WriteJSONString writes s as a JSON string, including the surrounding quotes, to w.
// Escaping rules are the same as in String2JSON (no HTML escaping).
// Write errors are ignored.
func WriteJSONString(s string, w io.Writer) {
	_, _ = io.WriteString(w, `"`)
	start := 0 // start of the section not yet written to w
	for i := 0; i < len(s); {
		c := s[i]
		size := 1
		var esc string
		if c < utf8.RuneSelf {
			switch {
			case c == '"':
				esc = `\"`
			case c == '\\':
				esc = `\\`
			case c < 0x20:
				esc = ctrlEscapes[c]
			default:
				// safe ASCII character, written later as part of the section
				i++
				continue
			}
		} else {
			var r rune
			r, size = utf8.DecodeRuneInString(s[i:])
			switch {
			case r == utf8.RuneError && size == 1:
				// invalid UTF-8
				esc = `\ufffd`
			case r == '\u2028':
				// valid in JSON, but a line terminator in older JavaScript
				esc = `\u2028`
			case r == '\u2029':
				esc = `\u2029`
			default:
				i += size
				continue
			}
		}
		if start < i {
			_, _ = io.WriteString(w, s[start:i])
		}
		_, _ = io.WriteString(w, esc)
		i += size
		start = i
	}
	if start < len(s) {
		_, _ = io.WriteString(w, s[start:])
	}
	_, _ = io.WriteString(w, `"`)
}
