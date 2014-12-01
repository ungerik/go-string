package str

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type String string

func (s String) String() string {
	return string(s)
}

func (s String) Contains(substr String) bool {
	return strings.Contains(string(s), string(substr))
}

func (s String) ContainsAny(chars String) bool {
	return strings.ContainsAny(string(s), string(chars))
}

func (s String) ContainsRune(r rune) bool {
	return strings.ContainsRune(string(s), r)
}

func (s String) Count(sep String) int {
	return strings.Count(string(s), string(sep))
}

func (s String) EqualFold(t String) bool {
	return strings.EqualFold(string(s), string(t))
}

func (s String) Fields() []String {
	return s.FieldsFunc(unicode.IsSpace)
}

func (s String) FieldsFunc(f func(rune) bool) []String {
	// First count the fields.
	n := 0
	inField := false
	for _, rune := range s {
		wasInField := inField
		inField = !f(rune)
		if inField && !wasInField {
			n++
		}
	}

	// Now create them.
	a := make([]String, n)
	na := 0
	fieldStart := -1 // Set to -1 when looking for start of field.
	for i, rune := range s {
		if f(rune) {
			if fieldStart >= 0 {
				a[na] = s[fieldStart:i]
				na++
				fieldStart = -1
			}
		} else if fieldStart == -1 {
			fieldStart = i
		}
	}
	if fieldStart >= 0 { // Last field might end at EOF.
		a[na] = s[fieldStart:]
	}
	return a
}

func (s String) HasPrefix(prefix String) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

func (s String) HasSuffix(suffix String) bool {
	return strings.HasSuffix(string(s), string(suffix))
}

func (s String) Index(sep String) int {
	return strings.Index(string(s), string(sep))
}

func (s String) IndexAny(chars String) int {
	return strings.IndexAny(string(s), string(chars))
}

func (s String) IndexByte(c byte) int {
	return strings.IndexByte(string(s), c)
}

func (s String) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(string(s), f)
}

func (s String) IndexRune(r rune) int {
	return strings.IndexRune(string(s), r)
}

func (sep String) Join(a []String) String {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return String(b)
}

func (s String) LastIndex(sep String) int {
	return strings.LastIndex(string(s), string(sep))
}

func (s String) LastIndexAny(chars String) int {
	return strings.LastIndexAny(string(s), string(chars))
}

func (s String) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(string(s), f)
}

func (s String) Map(mapping func(rune) rune) String {
	return String(strings.Map(mapping, string(s)))
}

func (s String) Repeat(count int) String {
	return String(strings.Repeat(string(s), count))
}

func (s String) Replace(old, new String, n int) String {
	return String(strings.Replace(string(s), string(old), string(new), n))
}

func (s String) Split(sep String) []String {
	return genSplit(s, sep, 0, -1)
}

func (s String) SplitAfter(sep String) []String {
	return genSplit(s, sep, len(sep), -1)
}

func (s String) SplitAfterN(sep String, n int) []String {
	return genSplit(s, sep, len(sep), n)
}

func (s String) SplitN(sep String, n int) []String {
	return genSplit(s, sep, 0, n)
}

func (s String) Title() String {
	return String(strings.Title(string(s)))
}

func (s String) ToLower() String {
	return String(strings.ToLower(string(s)))
}

func (s String) ToLowerSpecial(_case unicode.SpecialCase) String {
	return String(strings.ToLowerSpecial(_case, string(s)))
}

func (s String) ToTitle() String {
	return String(strings.ToTitle(string(s)))
}

func (s String) ToTitleSpecial(_case unicode.SpecialCase) String {
	return String(strings.ToTitleSpecial(_case, string(s)))
}

func (s String) ToUpper() String {
	return String(strings.ToUpper(string(s)))
}

func (s String) ToUpperSpecial(_case unicode.SpecialCase) String {
	return String(strings.ToUpperSpecial(_case, string(s)))
}

func (s String) Trim(cutset String) String {
	return String(strings.Trim(string(s), string(cutset)))
}

func (s String) TrimFunc(f func(rune) bool) String {
	return String(strings.TrimFunc(string(s), f))
}

func (s String) TrimLeft(cutset String) String {
	return String(strings.TrimLeft(string(s), string(cutset)))
}

func (s String) TrimLeftFunc(f func(rune) bool) String {
	return String(strings.TrimLeftFunc(string(s), f))
}

func (s String) TrimPrefix(prefix String) String {
	return String(strings.TrimPrefix(string(s), string(prefix)))
}

func (s String) TrimRight(cutset String) String {
	return String(strings.TrimRight(string(s), string(cutset)))
}

func (s String) TrimRightFunc(f func(rune) bool) String {
	return String(strings.TrimRightFunc(string(s), f))
}

func (s String) TrimSpace() String {
	return String(strings.TrimSpace(string(s)))
}

func (s String) TrimSuffix(suffix String) String {
	return String(strings.TrimSuffix(string(s), string(suffix)))
}

func genSplit(s, sep String, sepSave, n int) []String {
	if n == 0 {
		return nil
	}
	if sep == "" {
		return explode(s, n)
	}
	if n < 0 {
		n = s.Count(sep) + 1
	}
	c := sep[0]
	start := 0
	a := make([]String, n)
	na := 0
	for i := 0; i+len(sep) <= len(s) && na+1 < n; i++ {
		if s[i] == c && (len(sep) == 1 || s[i:i+len(sep)] == sep) {
			a[na] = s[start : i+sepSave]
			na++
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	a[na] = s[start:]
	return a[0 : na+1]
}

func explode(s String, n int) []String {
	if n == 0 {
		return nil
	}
	l := utf8.RuneCountInString(string(s))
	if n <= 0 || n > l {
		n = l
	}
	a := make([]String, n)
	var size int
	var ch rune
	i, cur := 0, 0
	for ; i+1 < n; i++ {
		ch, size = utf8.DecodeRuneInString(string(s[cur:]))
		if ch == utf8.RuneError {
			a[i] = String(utf8.RuneError)
		} else {
			a[i] = s[cur : cur+size]
		}
		cur += size
	}
	// add the rest, if there is any
	if cur < len(s) {
		a[i] = s[cur:]
	}
	return a
}
