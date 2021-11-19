// Generated code; DO NOT EDIT.
package test

import (
	"net/url"
	"strings"
	
)

type String string

func (s String) IsEmpty() bool {
	return &s == nil || s == "" || s == " "
}
func (s String) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s String) ToLower() String {
	return String(strings.ToLower(s.Str()))
}

func (s String) ToUpper() String {
	return String(strings.ToUpper(s.Str()))
}

func (s String) ToTitle() String {
	return String(strings.ToTitle(s.Str()))
}

func (s String) Contains(that string) bool {
	return strings.Contains(s.Str(), that)
}

func (s String) NotContains(that string) bool {
	return !strings.Contains(s.Str(), that)
}

func (s String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.Str(), prefix)
}

func (s String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.Str(), suffix)
}

func (s String) Map(f func(string) string) String {
	return String(f(s.Str()))
}

func (s String) Split(sep string) StringList {
	xs := strings.Split(s.Str(), sep)
	return EmptyStringList().AppendAll(xs...)
}

func (s String) Append(str string) String {
	return String(s.Str() + str)
}

func (s String) Count() int {
	return len(s.Str())
}

func (s String) ToString() string {
	return string(s)
}

func (s String) ToBytes() []byte {
	return []byte(s)
}

func (s String) QueryEscape() String {
	return String(url.QueryEscape(s.Str()))
}

func (s String) QueryUnEscape() (String, error) {
	str, err := url.QueryUnescape(s.Str())
	return String(str), err
}

func (s String) TrimSpace() String {
	return String(strings.TrimSpace(s.Str()))
}

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
func (s String) TrimLeft(cutset string) String {
	return String(strings.TrimLeft(s.Str(), cutset))
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func (s String) TrimPrefix(prefix string) String {
	return String(strings.TrimPrefix(s.Str(), prefix))
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func (s String) TrimSuffix(suffix string) String {
	return String(strings.TrimSuffix(s.Str(), suffix))
}

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
func (s String) ReplaceAll(oldString, newString string) String {
	return String(strings.ReplaceAll(s.Str(), oldString, newString))
}

func (s String) Str() string {
	return string(s)
}
