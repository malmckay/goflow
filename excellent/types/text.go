package types

import (
	"encoding/json"
	"strings"
	"unicode/utf8"

	"github.com/nyaruka/goflow/utils"
)

// XText is a simple tex value
type XText struct {
	native string
}

// NewXText creates a new text value
func NewXText(value string) XText {
	return XText{native: value}
}

// Reduce returns the primitive version of this type (i.e. itself)
func (x XText) Reduce() XPrimitive { return x }

// ToXText converts this type to text
func (x XText) ToXText() XText { return x }

// ToXBoolean converts this type to a bool
func (x XText) ToXBoolean() XBoolean {
	return NewXBoolean(!x.Empty() && strings.ToLower(x.Native()) != "false")
}

// ToXJSON is called when this type is passed to @(json(...))
func (x XText) ToXJSON() XText { return MustMarshalToXText(x.Native()) }

// Native returns the native value of this type
func (x XText) Native() string { return x.native }

// String returns the native string representation of this type
func (x XText) String() string { return x.Native() }

// Equals determines equality for this type
func (x XText) Equals(other XText) bool {
	return x.Native() == other.Native()
}

// Compare compares this string to another
func (x XText) Compare(other XText) int {
	return strings.Compare(x.Native(), other.Native())
}

// Length returns the length of this string
func (x XText) Length() int { return utf8.RuneCountInString(x.Native()) }

// Empty returns whether this is an empty string
func (x XText) Empty() bool { return x.Native() == "" }

// MarshalJSON is called when a struct containing this type is marshaled
func (x XText) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Native())
}

// UnmarshalJSON is called when a struct containing this type is unmarshaled
func (x *XText) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &x.native)
}

// XTextEmpty is the empty text value
var XTextEmpty = NewXText("")
var _ XPrimitive = XTextEmpty
var _ XLengthable = XTextEmpty

// ToXText converts the given value to a string
func ToXText(x XValue) (XText, XError) {
	if utils.IsNil(x) {
		return XTextEmpty, nil
	}
	if IsXError(x) {
		return XTextEmpty, x.(XError)
	}

	return x.Reduce().ToXText(), nil
}
