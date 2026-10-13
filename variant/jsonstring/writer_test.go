package jsonstring

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"testing"
	"unicode/utf8"
)

// plainWriter hides any WriteString method of the wrapped writer,
// so the io.Writer fallback path of io.WriteString is exercised as well.
type plainWriter struct{ w io.Writer }

func (p plainWriter) Write(b []byte) (int, error) { return p.w.Write(b) }

// writeToString runs WriteJSONString with both a StringWriter and a plain Writer
// and fails the test if the results differ.
func writeToString(t testing.TB, s string) string {
	t.Helper()
	var sb strings.Builder
	WriteJSONString(s, &sb)

	var pb bytes.Buffer
	WriteJSONString(s, plainWriter{&pb})

	if sb.String() != pb.String() {
		t.Fatalf("StringWriter and plain Writer differ for %q: %q vs %q", s, sb.String(), pb.String())
	}
	return sb.String()
}

// reference encodes s with encoding/json without HTML escaping.
func reference(t testing.TB, s string) string {
	t.Helper()
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(s); err != nil {
		t.Fatalf("reference encoding failed: %v", err)
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

func TestWriteJSONString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", `""`},
		{"plain", "hello", `"hello"`},
		{"html chars", "<a href=\"x\">&</a>", `"<a href=\"x\">&</a>"`},
		{"quote", `say "hi"`, `"say \"hi\""`},
		{"backslash", `C:\temp`, `"C:\\temp"`},
		{"newline tab cr", "a\nb\tc\rd", `"a\nb\tc\rd"`},
		{"backspace formfeed", "a\bb\fc", `"a\bb\fc"`},
		{"other control chars", "\x00\x01\x1f", `"\u0000\u0001\u001f"`},
		{"DEL is not escaped", "\x7f", "\"\x7f\""},
		{"multibyte utf8", "Grüße 日本 🏃", `"Grüße 日本 🏃"`},
		{"line separators", "a\u2028b\u2029c", `"a\u2028b\u2029c"`},
		{"invalid utf8", "a\xffb", `"a\ufffdb"`},
		{"truncated utf8", "a\xe2\x82", `"a\ufffd\ufffd"`},
		{"escape at start and end", "\n<x>\n", `"\n<x>\n"`},
		{"only escapes", "\"\\\"", `"\"\\\""`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := writeToString(t, tt.in)
			if got != tt.want {
				t.Errorf("WriteJSONString(%q) = %s, want %s", tt.in, got, tt.want)
			}
			if ref := reference(t, tt.in); got != ref {
				t.Errorf("WriteJSONString(%q) = %s, encoding/json = %s", tt.in, got, ref)
			}
		})
	}
}

func FuzzWriteJSONString(f *testing.F) {
	seeds := []string{
		"", "hello", "<>&", "\"\\", "\x00\x1f\x7f",
		"\u2028\u2029", "a\xffb", "Grüße 🏃", strings.Repeat("<tag>", 100),
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		got := writeToString(t, s)

		// must match encoding/json byte for byte
		if ref := reference(t, s); got != ref {
			t.Fatalf("mismatch for %q:\n got: %s\nwant: %s", s, got, ref)
		}

		// must be valid JSON that decodes back to the input (for valid UTF-8)
		var decoded string
		if err := json.Unmarshal([]byte(got), &decoded); err != nil {
			t.Fatalf("output is not valid JSON for %q: %v", s, err)
		}
		if utf8.ValidString(s) && decoded != s {
			t.Fatalf("round trip failed: %q -> %s -> %q", s, got, decoded)
		}

		// no HTML escaping
		if strings.Contains(got, `\u003c`) || strings.Contains(got, `\u003e`) || strings.Contains(got, `\u0026`) {
			t.Fatalf("unexpected HTML escaping in %s", got)
		}
	})
}
