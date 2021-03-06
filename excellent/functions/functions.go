package functions

import (
	"bytes"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/nyaruka/gocommon/urns"
	"github.com/nyaruka/goflow/excellent/types"
	"github.com/nyaruka/goflow/utils"

	humanize "github.com/dustin/go-humanize"
	"github.com/shopspring/decimal"
)

var nonPrintableRegex = regexp.MustCompile(`[\p{Cc}\p{C}]`)

// XFunction defines the interface that Excellent functions must implement
type XFunction func(env utils.Environment, args ...types.XValue) types.XValue

// RegisterXFunction registers a new function in Excellent
func RegisterXFunction(name string, function XFunction) {
	XFUNCTIONS[name] = function
}

// XFUNCTIONS is our map of functions available in Excellent which aren't tests
var XFUNCTIONS = map[string]XFunction{
	// type conversion
	"text":     OneArgFunction(Text),
	"boolean":  OneArgFunction(Boolean),
	"number":   OneArgFunction(Number),
	"datetime": OneTextFunction(DateTime),
	"array":    Array,

	// text functions
	"char":              OneNumberFunction(Char),
	"code":              OneTextFunction(Code),
	"split":             TwoTextFunction(Split),
	"join":              TwoArgFunction(Join),
	"title":             OneTextFunction(Title),
	"word":              TextAndIntegerFunction(Word),
	"remove_first_word": OneTextFunction(RemoveFirstWord),
	"word_count":        OneTextFunction(WordCount),
	"word_slice":        ArgCountCheck(2, 3, WordSlice),
	"field":             ArgCountCheck(3, 3, Field),
	"clean":             OneTextFunction(Clean),
	"left":              TextAndIntegerFunction(Left),
	"lower":             OneTextFunction(Lower),
	"right":             TextAndIntegerFunction(Right),
	"text_compare":      TwoTextFunction(TextCompare),
	"repeat":            TextAndIntegerFunction(Repeat),
	"replace":           ThreeTextFunction(Replace),
	"upper":             OneTextFunction(Upper),
	"percent":           OneNumberFunction(Percent),
	"url_encode":        OneTextFunction(URLEncode),

	// bool functions
	"and": ArgCountCheck(1, -1, And),
	"if":  ThreeArgFunction(If),
	"or":  ArgCountCheck(1, -1, Or),

	// number functions
	"round":        OneNumberAndOptionalIntegerFunction(Round, 0),
	"round_up":     OneNumberAndOptionalIntegerFunction(RoundUp, 0),
	"round_down":   OneNumberAndOptionalIntegerFunction(RoundDown, 0),
	"max":          ArgCountCheck(1, -1, Max),
	"min":          ArgCountCheck(1, -1, Min),
	"mean":         ArgCountCheck(1, -1, Mean),
	"mod":          TwoNumberFunction(Mod),
	"rand":         NoArgFunction(Rand),
	"rand_between": TwoNumberFunction(RandBetween),
	"abs":          OneNumberFunction(Abs),

	// datetime functions
	"parse_datetime":      ArgCountCheck(2, 3, ParseDateTime),
	"datetime_from_parts": ArgCountCheck(3, 3, DateTimeFromParts),
	"datetime_diff":       DateTimeDiff,
	"datetime_add":        DateTimeAdd,
	"weekday":             OneDateTimeFunction(Weekday),
	"tz":                  OneDateTimeFunction(TZ),
	"tz_offset":           OneDateTimeFunction(TZOffset),
	"today":               NoArgFunction(Today),
	"now":                 NoArgFunction(Now),
	"from_epoch":          OneNumberFunction(FromEpoch),
	"to_epoch":            OneDateTimeFunction(ToEpoch),

	// json functions
	"json":       OneArgFunction(JSON),
	"parse_json": OneTextFunction(ParseJSON),

	// formatting functions
	"format_datetime": FormatDateTime,
	"format_location": OneTextFunction(FormatLocation),
	"format_number":   FormatNumber,
	"format_urn":      FormatURN,

	// utility functions
	"length":     OneArgFunction(Length),
	"default":    TwoArgFunction(Default),
	"legacy_add": TwoArgFunction(LegacyAdd),
	"read_chars": OneTextFunction(ReadChars),
}

//------------------------------------------------------------------------------------------
// Type Conversion Functions
//------------------------------------------------------------------------------------------

// Text tries to convert `value` to text. An error is returned if the value can't be converted.
//
//   @(text(3 = 3)) -> true
//   @(json(text(123.45))) -> "123.45"
//   @(text(1 / 0)) -> ERROR
//
// @function text(value)
func Text(env utils.Environment, value types.XValue) types.XValue {
	str, xerr := types.ToXText(env, value)
	if xerr != nil {
		return xerr
	}
	return str
}

// Boolean tries to convert `value` to a boolean. An error is returned if the value can't be converted.
//
//   @(boolean(array(1, 2))) -> true
//   @(boolean("FALSE")) -> false
//   @(boolean(1 / 0)) -> ERROR
//
// @function boolean(value)
func Boolean(env utils.Environment, value types.XValue) types.XValue {
	str, xerr := types.ToXBoolean(env, value)
	if xerr != nil {
		return xerr
	}
	return str
}

// Number tries to convert `value` to a number. An error is returned if the value can't be converted.
//
//   @(number(10)) -> 10
//   @(number("123.45000")) -> 123.45
//   @(number("what?")) -> ERROR
//
// @function number(value)
func Number(env utils.Environment, value types.XValue) types.XValue {
	num, xerr := types.ToXNumber(env, value)
	if xerr != nil {
		return xerr
	}
	return num
}

// DateTime turns `text` into a date according to the environment's settings. It will return an error
// if it is unable to convert the text to a date.
//
//   @(datetime("1979-07-18")) -> 1979-07-18T00:00:00.000000-05:00
//   @(datetime("1979-07-18T10:30:45.123456Z")) -> 1979-07-18T10:30:45.123456Z
//   @(datetime("2010 05 10")) -> 2010-05-10T00:00:00.000000-05:00
//   @(datetime("NOT DATE")) -> ERROR
//
// @function datetime(text)
func DateTime(env utils.Environment, str types.XText) types.XValue {
	date, err := utils.DateFromString(env, str.Native())
	if err != nil {
		return types.NewXError(err)
	}

	return types.NewXDateTime(date)
}

// Array takes a list of `values` and returns them as an array
//
//   @(array("a", "b", 356)[1]) -> b
//   @(join(array("a", "b", "c"), "|")) -> a|b|c
//   @(length(array())) -> 0
//   @(length(array("a", "b"))) -> 2
//
// @function array(values...)
func Array(env utils.Environment, args ...types.XValue) types.XValue {
	// check none of our args are errors
	for _, arg := range args {
		if types.IsXError(arg) {
			return arg
		}
	}

	return types.NewXArray(args...)
}

//------------------------------------------------------------------------------------------
// Bool Functions
//------------------------------------------------------------------------------------------

// And returns whether all the passed in arguments are truthy
//
//   @(and(true)) -> true
//   @(and(true, false, true)) -> false
//
// @function and(tests...)
func And(env utils.Environment, args ...types.XValue) types.XValue {
	for _, arg := range args {
		asBool, xerr := types.ToXBoolean(env, arg)
		if xerr != nil {
			return xerr
		}
		if !asBool.Native() {
			return types.XBooleanFalse
		}
	}
	return types.XBooleanTrue
}

// Or returns whether if any of the passed in arguments are truthy
//
//   @(or(true)) -> true
//   @(or(true, false, true)) -> true
//
// @function or(tests...)
func Or(env utils.Environment, args ...types.XValue) types.XValue {
	for _, arg := range args {
		asBool, xerr := types.ToXBoolean(env, arg)
		if xerr != nil {
			return xerr
		}
		if asBool.Native() {
			return types.XBooleanTrue
		}
	}
	return types.XBooleanFalse
}

// If evaluates the `test` argument, and if truthy returns `true_value`, if not returning `false_value`
//
// If the first argument is an error that error is returned
//
//   @(if(1 = 1, "foo", "bar")) -> foo
//   @(if("foo" > "bar", "foo", "bar")) -> ERROR
//
// @function if(test, true_value, false_value)
func If(env utils.Environment, test types.XValue, arg1 types.XValue, arg2 types.XValue) types.XValue {
	asBool, err := types.ToXBoolean(env, test)
	if err != nil {
		return err
	}

	if asBool.Native() {
		return arg1
	}
	return arg2
}

//------------------------------------------------------------------------------------------
// Text Functions
//------------------------------------------------------------------------------------------

// Code returns the numeric code for the first character in `text`, it is the inverse of char
//
//   @(code("a")) -> 97
//   @(code("abc")) -> 97
//   @(code("😀")) -> 128512
//   @(code("15")) -> 49
//   @(code(15)) -> 49
//   @(code("")) -> ERROR
//
// @function code(text)
func Code(env utils.Environment, text types.XText) types.XValue {
	if text.Length() == 0 {
		return types.NewXErrorf("requires a string of at least one character")
	}

	r, _ := utf8.DecodeRuneInString(text.Native())
	return types.NewXNumberFromInt(int(r))
}

// Split splits `text` based on the passed in `delimeter`
//
// Empty values are removed from the returned list
//
//   @(split("a b c", " ")) -> ["a","b","c"]
//   @(split("a", " ")) -> ["a"]
//   @(split("abc..d", ".")) -> ["abc","d"]
//   @(split("a.b.c.", ".")) -> ["a","b","c"]
//   @(split("a && b && c", " && ")) -> ["a","b","c"]
//
// @function split(text, delimiter)
func Split(env utils.Environment, text types.XText, sep types.XText) types.XValue {
	splits := types.NewXArray()
	allSplits := strings.Split(text.Native(), sep.Native())
	for i := range allSplits {
		if allSplits[i] != "" {
			splits.Append(types.NewXText(allSplits[i]))
		}
	}
	return splits
}

// Join joins the passed in `array` of strings with the passed in `delimeter`
//
//   @(join(array("a", "b", "c"), "|")) -> a|b|c
//   @(join(split("a.b.c", "."), " ")) -> a b c
//
// @function join(array, delimiter)
func Join(env utils.Environment, array types.XValue, delimiter types.XValue) types.XValue {
	indexable, isIndexable := array.(types.XIndexable)
	if !isIndexable {
		return types.NewXErrorf("requires an indexable as its first argument")
	}

	sep, xerr := types.ToXText(env, delimiter)
	if xerr != nil {
		return xerr
	}

	var output bytes.Buffer
	for i := 0; i < indexable.Length(); i++ {
		if i > 0 {
			output.WriteString(sep.Native())
		}
		itemAsStr, xerr := types.ToXText(env, indexable.Index(i))
		if xerr != nil {
			return xerr
		}

		output.WriteString(itemAsStr.Native())
	}

	return types.NewXText(output.String())
}

// Char returns the rune for the passed in codepoint, `num`, which may be unicode, this is the reverse of code
//
//   @(char(33)) -> !
//   @(char(128512)) -> 😀
//   @(char("foo")) -> ERROR
//
// @function char(num)
func Char(env utils.Environment, num types.XNumber) types.XValue {
	code, xerr := types.ToInteger(env, num)
	if xerr != nil {
		return xerr
	}

	return types.NewXText(string(rune(code)))
}

// Title titlecases the passed in `text`, capitalizing each word
//
//   @(title("foo")) -> Foo
//   @(title("ryan lewis")) -> Ryan Lewis
//   @(title(123)) -> 123
//
// @function title(text)
func Title(env utils.Environment, text types.XText) types.XValue {
	return types.NewXText(strings.Title(text.Native()))
}

// Word returns the word at the passed in `index` for the passed in `text`
//
//   @(word("bee cat dog", 0)) -> bee
//   @(word("bee.cat,dog", 0)) -> bee
//   @(word("bee.cat,dog", 1)) -> cat
//   @(word("bee.cat,dog", 2)) -> dog
//   @(word("bee.cat,dog", -1)) -> dog
//   @(word("bee.cat,dog", -2)) -> cat
//
// @function word(text, index)
func Word(env utils.Environment, text types.XText, index int) types.XValue {
	words := utils.TokenizeString(text.Native())

	offset := index
	if offset < 0 {
		offset += len(words)
	}

	if !(offset >= 0 && offset < len(words)) {
		return types.NewXErrorf("index %d is out of range for the number of words %d", index, len(words))
	}

	return types.NewXText(words[offset])
}

// RemoveFirstWord removes the 1st word of `text`
//
//   @(remove_first_word("foo bar")) -> bar
//
// @function remove_first_word(text)
func RemoveFirstWord(env utils.Environment, text types.XText) types.XValue {
	words := utils.TokenizeString(text.Native())
	if len(words) > 1 {
		return types.NewXText(strings.Join(words[1:], " "))
	}

	return types.XTextEmpty
}

// WordSlice extracts a substring from `text` spanning from `start` up to but not-including `end`. (first word is 0). A negative
// end value means that all words after the start should be returned.
//
//   @(word_slice("bee cat dog", 0, 1)) -> bee
//   @(word_slice("bee cat dog", 0, 2)) -> bee cat
//   @(word_slice("bee cat dog", 1, -1)) -> cat dog
//   @(word_slice("bee cat dog", 1)) -> cat dog
//   @(word_slice("bee cat dog", 2, 3)) -> dog
//   @(word_slice("bee cat dog", 3, 10)) ->
//
// @function word_slice(text, start, end)
func WordSlice(env utils.Environment, args ...types.XValue) types.XValue {
	str, xerr := types.ToXText(env, args[0])
	if xerr != nil {
		return xerr
	}

	start, xerr := types.ToInteger(env, args[1])
	if xerr != nil {
		return xerr
	}
	if start < 0 {
		return types.NewXErrorf("must start with a positive index")
	}

	end := -1
	if len(args) == 3 {
		if end, xerr = types.ToInteger(env, args[2]); xerr != nil {
			return xerr
		}
	}
	if end > 0 && end <= start {
		return types.NewXErrorf("must have a end which is greater than the start")
	}

	words := utils.TokenizeString(str.Native())

	if start >= len(words) {
		return types.XTextEmpty
	}
	if end >= len(words) {
		end = len(words)
	}

	if end > 0 {
		return types.NewXText(strings.Join(words[start:end], " "))
	}
	return types.NewXText(strings.Join(words[start:], " "))
}

// WordCount returns the number of words in `text`
//
//   @(word_count("foo bar")) -> 2
//   @(word_count(10)) -> 1
//   @(word_count("")) -> 0
//   @(word_count("😀😃😄😁")) -> 4
//
// @function word_count(text)
func WordCount(env utils.Environment, text types.XText) types.XValue {
	words := utils.TokenizeString(text.Native())
	return types.NewXNumberFromInt(len(words))
}

// Field splits `text` based on the passed in `delimiter` and returns the field at `offset`.  When splitting
// with a space, the delimiter is considered to be all whitespace.  (first field is 0)
//
//   @(field("a,b,c", 1, ",")) -> b
//   @(field("a,,b,c", 1, ",")) ->
//   @(field("a   b c", 1, " ")) -> b
//   @(field("a		b	c	d", 1, "	")) ->
//   @(field("a\t\tb\tc\td", 1, " ")) ->
//   @(field("a,b,c", "foo", ",")) -> ERROR
//
// @function field(text, offset, delimeter)
func Field(env utils.Environment, args ...types.XValue) types.XValue {
	source, xerr := types.ToXText(env, args[0])
	if xerr != nil {
		return xerr
	}

	field, xerr := types.ToInteger(env, args[1])
	if xerr != nil {
		return xerr
	}

	if field < 0 {
		return types.NewXErrorf("cannot use a negative index to FIELD")
	}

	sep, xerr := types.ToXText(env, args[2])
	if xerr != nil {
		return xerr
	}

	fields := strings.Split(source.Native(), sep.Native())
	if field >= len(fields) {
		return types.XTextEmpty
	}

	// when using a space as a delimiter, we consider it splitting on whitespace, so remove empty values
	if sep.Native() == " " {
		var newFields []string
		for _, field := range fields {
			if field != "" {
				newFields = append(newFields, field)
			}
		}
		fields = newFields
	}

	return types.NewXText(strings.TrimSpace(fields[field]))
}

// Clean strips any non-printable characters from `text`
//
//   @(clean("😃 Hello \nwo\tr\rld")) -> 😃 Hello world
//   @(clean(123)) -> 123
//
// @function clean(text)
func Clean(env utils.Environment, text types.XText) types.XValue {
	return types.NewXText(nonPrintableRegex.ReplaceAllString(text.Native(), ""))
}

// Left returns the `count` most left characters of the passed in `text`
//
//   @(left("hello", 2)) -> he
//   @(left("hello", 7)) -> hello
//   @(left("😀😃😄😁", 2)) -> 😀😃
//   @(left("hello", -1)) -> ERROR
//
// @function left(text, count)
func Left(env utils.Environment, text types.XText, count int) types.XValue {
	if count < 0 {
		return types.NewXErrorf("can't take a negative count")
	}

	// this weird construct does the right thing for multi-byte unicode
	var output bytes.Buffer
	i := 0
	for _, r := range text.Native() {
		if i >= count {
			break
		}
		output.WriteRune(r)
		i++
	}

	return types.NewXText(output.String())
}

// Lower lowercases the passed in `text`
//
//   @(lower("HellO")) -> hello
//   @(lower("hello")) -> hello
//   @(lower("123")) -> 123
//   @(lower("😀")) -> 😀
//
// @function lower(text)
func Lower(env utils.Environment, text types.XText) types.XValue {
	return types.NewXText(strings.ToLower(text.Native()))
}

// Right returns the `count` most right characters of the passed in `text`
//
//   @(right("hello", 2)) -> lo
//   @(right("hello", 7)) -> hello
//   @(right("😀😃😄😁", 2)) -> 😄😁
//   @(right("hello", -1)) -> ERROR
//
// @function right(text, count)
func Right(env utils.Environment, text types.XText, count int) types.XValue {
	if count < 0 {
		return types.NewXErrorf("can't take a negative count")
	}

	start := utf8.RuneCountInString(text.Native()) - count

	// this weird construct does the right thing for multi-byte unicode
	var output bytes.Buffer
	i := 0
	for _, r := range text.Native() {
		if i >= start {
			output.WriteRune(r)
		}
		i++
	}

	return types.NewXText(output.String())
}

// TextCompare returns the comparison between the strings `text1` and `text2`.
// The return value will be -1 if str1 is smaller than str2, 0 if they
// are equal and 1 if str1 is greater than str2
//
//   @(text_compare("abc", "abc")) -> 0
//   @(text_compare("abc", "def")) -> -1
//   @(text_compare("zzz", "aaa")) -> 1
//
// @function text_compare(text1, text2)
func TextCompare(env utils.Environment, text1 types.XText, text2 types.XText) types.XValue {
	return types.NewXNumberFromInt(text1.Compare(text2))
}

// Repeat return `text` repeated `count` number of times
//
//   @(repeat("*", 8)) -> ********
//   @(repeat("*", "foo")) -> ERROR
//
// @function repeat(text, count)
func Repeat(env utils.Environment, text types.XText, count int) types.XValue {
	if count < 0 {
		return types.NewXErrorf("must be called with a positive integer, got %d", count)
	}

	var output bytes.Buffer
	for j := 0; j < count; j++ {
		output.WriteString(text.Native())
	}

	return types.NewXText(output.String())
}

// Replace replaces all occurrences of `needle` with `replacement` in `text`
//
//   @(replace("foo bar", "foo", "zap")) -> zap bar
//   @(replace("foo bar", "baz", "zap")) -> foo bar
//
// @function replace(text, needle, replacement)
func Replace(env utils.Environment, text types.XText, needle types.XText, replacement types.XText) types.XValue {
	return types.NewXText(strings.Replace(text.Native(), needle.Native(), replacement.Native(), -1))
}

// Upper uppercases all characters in the passed `text`
//
//   @(upper("Asdf")) -> ASDF
//   @(upper(123)) -> 123
//
// @function upper(text)
func Upper(env utils.Environment, text types.XText) types.XValue {
	return types.NewXText(strings.ToUpper(text.Native()))
}

// Percent converts `num` to text represented as a percentage
//
//   @(percent(0.54234)) -> 54%
//   @(percent(1.2)) -> 120%
//   @(percent("foo")) -> ERROR
//
// @function percent(num)
func Percent(env utils.Environment, num types.XNumber) types.XValue {
	// multiply by 100 and floor
	percent := num.Native().Mul(decimal.NewFromFloat(100)).Round(0)

	// add on a %
	return types.NewXText(fmt.Sprintf("%d%%", percent.IntPart()))
}

// URLEncode URL encodes `text` for use in a URL parameter
//
//   @(url_encode("two words")) -> two+words
//   @(url_encode(10)) -> 10
//
// @function url_encode(text)
func URLEncode(env utils.Environment, text types.XText) types.XValue {
	return types.NewXText(url.QueryEscape(text.Native()))
}

//------------------------------------------------------------------------------------------
// Number Functions
//------------------------------------------------------------------------------------------

// Abs returns the absolute value of `num`
//
//   @(abs(-10)) -> 10
//   @(abs(10.5)) -> 10.5
//   @(abs("foo")) -> ERROR
//
// @function abs(num)
func Abs(env utils.Environment, num types.XNumber) types.XValue {
	return types.NewXNumber(num.Native().Abs())
}

// Round rounds `num` to the nearest value. You can optionally pass in the number of decimal places to round to as `places`.
//
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
//   @(round(12)) -> 12
//   @(round(12.141)) -> 12
//   @(round(12.6)) -> 13
//   @(round(12.141, 2)) -> 12.14
//   @(round(12.146, 2)) -> 12.15
//   @(round(12.146, -1)) -> 10
//   @(round("notnum", 2)) -> ERROR
//
// @function round(num [,places])
func Round(env utils.Environment, num types.XNumber, places int) types.XValue {
	return types.NewXNumber(num.Native().Round(int32(places)))
}

// RoundUp rounds `num` up to the nearest integer value. You can optionally pass in the number of decimal places to round to as `places`.
//
//   @(round_up(12)) -> 12
//   @(round_up(12.141)) -> 13
//   @(round_up(12.6)) -> 13
//   @(round_up(12.141, 2)) -> 12.15
//   @(round_up(12.146, 2)) -> 12.15
//   @(round_up("foo")) -> ERROR
//
// @function round_up(num [,places])
func RoundUp(env utils.Environment, num types.XNumber, places int) types.XValue {
	dec := num.Native()
	if dec.Round(int32(places)).Equal(dec) {
		return num
	}

	halfPrecision := decimal.New(5, -int32(places)-1)
	roundedDec := dec.Add(halfPrecision).Round(int32(places))

	return types.NewXNumber(roundedDec)
}

// RoundDown rounds `num` down to the nearest integer value. You can optionally pass in the number of decimal places to round to as `places`.
//
//   @(round_down(12)) -> 12
//   @(round_down(12.141)) -> 12
//   @(round_down(12.6)) -> 12
//   @(round_down(12.141, 2)) -> 12.14
//   @(round_down(12.146, 2)) -> 12.14
//   @(round_down("foo")) -> ERROR
//
// @function round_down(num [,places])
func RoundDown(env utils.Environment, num types.XNumber, places int) types.XValue {
	dec := num.Native()
	if dec.Round(int32(places)).Equal(dec) {
		return num
	}

	halfPrecision := decimal.New(5, -int32(places)-1)
	roundedDec := dec.Sub(halfPrecision).Round(int32(places))

	return types.NewXNumber(roundedDec)
}

// Max takes a list of `values` and returns the greatest of them
//
//   @(max(1, 2)) -> 2
//   @(max(1, -1, 10)) -> 10
//   @(max(1, 10, "foo")) -> ERROR
//
// @function max(values...)
func Max(env utils.Environment, args ...types.XValue) types.XValue {
	max, xerr := types.ToXNumber(env, args[0])
	if xerr != nil {
		return xerr
	}

	for _, v := range args[1:] {
		val, xerr := types.ToXNumber(env, v)
		if xerr != nil {
			return xerr
		}

		if val.Compare(max) > 0 {
			max = val
		}
	}
	return max
}

// Min takes a list of `values` and returns the smallest of them
//
//   @(min(1, 2)) -> 1
//   @(min(2, 2, -10)) -> -10
//   @(min(1, 2, "foo")) -> ERROR
//
// @function min(values)
func Min(env utils.Environment, args ...types.XValue) types.XValue {
	max, xerr := types.ToXNumber(env, args[0])
	if xerr != nil {
		return xerr
	}

	for _, v := range args[1:] {
		val, xerr := types.ToXNumber(env, v)
		if xerr != nil {
			return xerr
		}

		if val.Compare(max) < 0 {
			max = val
		}
	}
	return max
}

// Mean takes a list of `values` and returns the arithmetic mean of them
//
//   @(mean(1, 2)) -> 1.5
//   @(mean(1, 2, 6)) -> 3
//   @(mean(1, "foo")) -> ERROR
//
// @function mean(values)
func Mean(env utils.Environment, args ...types.XValue) types.XValue {
	sum := decimal.Zero

	for _, val := range args {
		num, xerr := types.ToXNumber(env, val)
		if xerr != nil {
			return xerr
		}
		sum = sum.Add(num.Native())
	}

	return types.NewXNumber(sum.Div(decimal.New(int64(len(args)), 0)))
}

// Mod returns the remainder of the division of `divident` by `divisor`
//
//   @(mod(5, 2)) -> 1
//   @(mod(4, 2)) -> 0
//   @(mod(5, "foo")) -> ERROR
//
// @function mod(dividend, divisor)
func Mod(env utils.Environment, num1 types.XNumber, num2 types.XNumber) types.XValue {
	return types.NewXNumber(num1.Native().Mod(num2.Native()))
}

// Rand returns a single random number between [0.0-1.0).
//
//   @(rand()) -> 0.3849275689214193274523267973563633859157562255859375
//   @(rand()) -> 0.607552015674623913099594574305228888988494873046875
//
// @function rand()
func Rand(env utils.Environment) types.XValue {
	return types.NewXNumber(utils.RandDecimal())
}

// RandBetween a single random integer in the given inclusive range.
//
//   @(rand_between(1, 10)) -> 5
//   @(rand_between(1, 10)) -> 10
//
// @function rand_between()
func RandBetween(env utils.Environment, min types.XNumber, max types.XNumber) types.XValue {
	span := (max.Native().Sub(min.Native())).Add(decimal.New(1, 0))

	val := utils.RandDecimal().Mul(span).Add(min.Native()).Floor()

	return types.NewXNumber(val)
}

//------------------------------------------------------------------------------------------
// Date & Time Functions
//------------------------------------------------------------------------------------------

// ParseDateTime turns `text` into a date according to the `format` and optional `timezone` specified
//
// The format string can consist of the following characters. The characters
// ' ', ':', ',', 'T', '-' and '_' are ignored. Any other character is an error.
//
// * `YY`        - last two digits of year 0-99
// * `YYYY`      - four digits of your 0000-9999
// * `M`         - month 1-12
// * `MM`        - month 01-12
// * `D`         - day of month, 1-31
// * `DD`        - day of month, zero padded 0-31
// * `h`         - hour of the day 1-12
// * `hh`        - hour of the day 01-12
// * `tt`        - twenty four hour of the day 01-23
// * `m`         - minute 0-59
// * `mm`        - minute 00-59
// * `s`         - second 0-59
// * `ss`        - second 00-59
// * `fff`       - milliseconds
// * `ffffff`    - microseconds
// * `fffffffff` - nanoseconds
// * `aa`        - am or pm
// * `AA`        - AM or PM
// * `Z`         - hour and minute offset from UTC, or Z for UTC
// * `ZZZ`       - hour and minute offset from UTC
//
// Timezone should be a location name as specified in the IANA Time Zone database, such
// as "America/Guayaquil" or "America/Los_Angeles". If not specified the timezone of your
// environment will be used. An error will be returned if the timezone is not recognized.
//
// Note that fractional seconds will be parsed even without an explicit format identifier.
// You should only specify fractional seconds when you want to assert the number of places
// in the input format.
//
// parse_datetime will return an error if it is unable to convert the text to a datetime.
//
//   @(parse_datetime("1979-07-18", "YYYY-MM-DD")) -> 1979-07-18T00:00:00.000000-05:00
//   @(parse_datetime("2010 5 10", "YYYY M DD")) -> 2010-05-10T00:00:00.000000-05:00
//   @(parse_datetime("2010 5 10 12:50", "YYYY M DD tt:mm", "America/Los_Angeles")) -> 2010-05-10T12:50:00.000000-07:00
//   @(parse_datetime("NOT DATE", "YYYY-MM-DD")) -> ERROR
//
// @function parse_datetime(text, format [,timezone])
func ParseDateTime(env utils.Environment, args ...types.XValue) types.XValue {
	str, xerr := types.ToXText(env, args[0])
	if xerr != nil {
		return xerr
	}

	format, xerr := types.ToXText(env, args[1])
	if xerr != nil {
		return xerr
	}

	// try to turn it to a go format
	goFormat, err := utils.ToGoDateFormat(format.Native())
	if err != nil {
		return types.NewXError(err)
	}

	// grab our location
	location := env.Timezone()
	if len(args) == 3 {
		tzStr, xerr := types.ToXText(env, args[2])
		if xerr != nil {
			return xerr
		}

		location, err = time.LoadLocation(tzStr.Native())
		if err != nil {
			return types.NewXError(err)
		}
	}

	// finally try to parse the date
	parsed, err := time.ParseInLocation(goFormat, str.Native(), location)
	if err != nil {
		return types.NewXError(err)
	}

	return types.NewXDateTime(parsed.In(location))
}

// DateTimeFromParts converts the passed in `year`, `month` and `day`
//
//   @(datetime_from_parts(2017, 1, 15)) -> 2017-01-15T00:00:00.000000-05:00
//   @(datetime_from_parts(2017, 2, 31)) -> 2017-03-03T00:00:00.000000-05:00
//   @(datetime_from_parts(2017, 13, 15)) -> ERROR
//
// @function datetime_from_parts(year, month, day)
func DateTimeFromParts(env utils.Environment, args ...types.XValue) types.XValue {
	year, xerr := types.ToInteger(env, args[0])
	if xerr != nil {
		return xerr
	}
	month, xerr := types.ToInteger(env, args[1])
	if xerr != nil {
		return xerr
	}
	if month < 1 || month > 12 {
		return types.NewXErrorf("invalid value for month, must be 1-12")
	}

	day, xerr := types.ToInteger(env, args[2])
	if xerr != nil {
		return xerr
	}

	return types.NewXDateTime(time.Date(year, time.Month(month), day, 0, 0, 0, 0, env.Timezone()))
}

// DateTimeDiff returns the integer duration between `date1` and `date2` in the `unit` specified.
//
// Valid durations are "Y" for years, "M" for months, "W" for weeks, "D" for days, "h" for hour,
// "m" for minutes, "s" for seconds
//
//   @(datetime_diff("2017-01-17", "2017-01-15", "D")) -> 2
//   @(datetime_diff("2017-01-17 10:50", "2017-01-17 12:30", "h")) -> -1
//   @(datetime_diff("2017-01-17", "2015-12-17", "Y")) -> 2
//
// @function datetime_diff(date1, date2, unit)
func DateTimeDiff(env utils.Environment, args ...types.XValue) types.XValue {
	if len(args) != 3 {
		return types.NewXErrorf("takes exactly three arguments, received %d", len(args))
	}

	date1, xerr := types.ToXDateTime(env, args[0])
	if xerr != nil {
		return xerr
	}

	date2, xerr := types.ToXDateTime(env, args[1])
	if xerr != nil {
		return xerr
	}

	unit, xerr := types.ToXText(env, args[2])
	if xerr != nil {
		return xerr
	}

	// find the duration between our dates
	duration := date1.Native().Sub(date2.Native())

	// then convert based on our unit
	switch unit.Native() {
	case "s":
		return types.NewXNumberFromInt(int(duration / time.Second))
	case "m":
		return types.NewXNumberFromInt(int(duration / time.Minute))
	case "h":
		return types.NewXNumberFromInt(int(duration / time.Hour))
	case "D":
		return types.NewXNumberFromInt(utils.DaysBetween(date1.Native(), date2.Native()))
	case "W":
		return types.NewXNumberFromInt(int(utils.DaysBetween(date1.Native(), date2.Native()) / 7))
	case "M":
		return types.NewXNumberFromInt(utils.MonthsBetween(date1.Native(), date2.Native()))
	case "Y":
		return types.NewXNumberFromInt(date1.Native().Year() - date2.Native().Year())
	}

	return types.NewXErrorf("unknown unit: %s, must be one of s, m, h, D, W, M, Y", unit)
}

// DateTimeAdd calculates the date value arrived at by adding `offset` number of `unit` to the `date`
//
// Valid durations are "Y" for years, "M" for months, "W" for weeks, "D" for days, "h" for hour,
// "m" for minutes, "s" for seconds
//
//   @(datetime_add("2017-01-15", 5, "D")) -> 2017-01-20T00:00:00.000000-05:00
//   @(datetime_add("2017-01-15 10:45", 30, "m")) -> 2017-01-15T11:15:00.000000-05:00
//
// @function datetime_add(date, offset, unit)
func DateTimeAdd(env utils.Environment, args ...types.XValue) types.XValue {
	if len(args) != 3 {
		return types.NewXErrorf("takes exactly three arguments, received %d", len(args))
	}

	date, xerr := types.ToXDateTime(env, args[0])
	if xerr != nil {
		return xerr
	}

	duration, xerr := types.ToInteger(env, args[1])
	if xerr != nil {
		return xerr
	}

	unit, xerr := types.ToXText(env, args[2])
	if xerr != nil {
		return xerr
	}

	switch unit.Native() {
	case "s":
		return types.NewXDateTime(date.Native().Add(time.Duration(duration) * time.Second))
	case "m":
		return types.NewXDateTime(date.Native().Add(time.Duration(duration) * time.Minute))
	case "h":
		return types.NewXDateTime(date.Native().Add(time.Duration(duration) * time.Hour))
	case "D":
		return types.NewXDateTime(date.Native().AddDate(0, 0, duration))
	case "W":
		return types.NewXDateTime(date.Native().AddDate(0, 0, duration*7))
	case "M":
		return types.NewXDateTime(date.Native().AddDate(0, duration, 0))
	case "Y":
		return types.NewXDateTime(date.Native().AddDate(duration, 0, 0))
	}

	return types.NewXErrorf("unknown unit: %s, must be one of s, m, h, D, W, M, Y", unit)
}

// Weekday returns the day of the week for `date`, 0 is sunday, 1 is monday..
//
//   @(weekday("2017-01-15")) -> 0
//   @(weekday("foo")) -> ERROR
//
// @function weekday(date)
func Weekday(env utils.Environment, date types.XDateTime) types.XValue {
	return types.NewXNumberFromInt(int(date.Native().Weekday()))
}

// TZ returns the timezone for `date``
//
// If not timezone information is present in the date, then the environment's
// timezone will be returned
//
//   @(tz("2017-01-15T02:15:18.123456Z")) -> UTC
//   @(tz("2017-01-15 02:15:18PM")) -> America/Guayaquil
//   @(tz("2017-01-15")) -> America/Guayaquil
//   @(tz("foo")) -> ERROR
//
// @function tz(date)
func TZ(env utils.Environment, date types.XDateTime) types.XValue {
	return types.NewXText(date.Native().Location().String())
}

// TZOffset returns the offset for the timezone as text +/- HHMM for `date`
//
// If no timezone information is present in the date, then the environment's
// timezone offset will be returned
//
//   @(tz_offset("2017-01-15T02:15:18.123456Z")) -> +0000
//   @(tz_offset("2017-01-15 02:15:18PM")) -> -0500
//   @(tz_offset("2017-01-15")) -> -0500
//   @(tz_offset("foo")) -> ERROR
//
// @function tz_offset(date)
func TZOffset(env utils.Environment, date types.XDateTime) types.XValue {
	// this looks like we are returning a set offset, but this is how go describes formats
	return types.NewXText(date.Native().Format("-0700"))

}

// Today returns the current date in the current timezone, time is set to midnight in the environment timezone
//
//   @(today()) -> 2018-04-11T00:00:00.000000-05:00
//
// @function today()
func Today(env utils.Environment) types.XValue {
	nowTZ := env.Now()
	return types.NewXDateTime(time.Date(nowTZ.Year(), nowTZ.Month(), nowTZ.Day(), 0, 0, 0, 0, env.Timezone()))
}

// FromEpoch returns a new date created from `num` which represents number of nanoseconds since January 1st, 1970 GMT
//
//   @(from_epoch(1497286619000000000)) -> 2017-06-12T11:56:59.000000-05:00
//
// @function from_epoch(num)
func FromEpoch(env utils.Environment, num types.XNumber) types.XValue {
	return types.NewXDateTime(time.Unix(0, num.Native().IntPart()).In(env.Timezone()))
}

// ToEpoch converts `date` to the number of nanoseconds since January 1st, 1970 GMT
//
//   @(to_epoch("2017-06-12T16:56:59.000000Z")) -> 1497286619000000000
//
// @function to_epoch(date)
func ToEpoch(env utils.Environment, date types.XDateTime) types.XValue {
	return types.NewXNumberFromInt64(date.Native().UnixNano())
}

// Now returns the current date and time in the environment timezone
//
//   @(now()) -> 2018-04-11T13:24:30.123456-05:00
//
// @function now()
func Now(env utils.Environment) types.XValue {
	return types.NewXDateTime(env.Now())
}

//------------------------------------------------------------------------------------------
// JSON Functions
//------------------------------------------------------------------------------------------

// ParseJSON tries to parse `text` as JSON, returning a fragment you can index into
//
// If the passed in value is not JSON, then an error is returned
//
//   @(parse_json("[1,2,3,4]").2) -> 3
//   @(parse_json("invalid json")) -> ERROR
//
// @function parse_json(text)
func ParseJSON(env utils.Environment, text types.XText) types.XValue {
	return types.JSONToXValue([]byte(text.Native()))
}

// JSON tries to return a JSON representation of `value`. An error is returned if there is
// no JSON representation of that object.
//
//   @(json("string")) -> "string"
//   @(json(10)) -> 10
//   @(json(contact.uuid)) -> "5d76d86b-3bb9-4d5a-b822-c9d86f5d8e4f"
//
// @function json(value)
func JSON(env utils.Environment, value types.XValue) types.XValue {
	asJSON, xerr := types.ToXJSON(env, value)
	if xerr != nil {
		return xerr
	}
	return asJSON
}

//----------------------------------------------------------------------------------------
// Formatting Functions
//----------------------------------------------------------------------------------------

// FormatDateTime turns `date` into text according to the `format` specified and in
// the optional `timezone`.
//
// The format string can consist of the following characters. The characters
// ' ', ':', ',', 'T', '-' and '_' are ignored. Any other character is an error.
//
// * `YY`        - last two digits of year 0-99
// * `YYYY`      - four digits of your 0000-9999
// * `M`         - month 1-12
// * `MM`        - month 01-12
// * `D`         - day of month, 1-31
// * `DD`        - day of month, zero padded 0-31
// * `h`         - hour of the day 1-12
// * `hh`        - hour of the day 01-12
// * `tt`        - twenty four hour of the day 01-23
// * `m`         - minute 0-59
// * `mm`        - minute 00-59
// * `s`         - second 0-59
// * `ss`        - second 00-59
// * `fff`       - milliseconds
// * `ffffff`    - microseconds
// * `fffffffff` - nanoseconds
// * `aa`        - am or pm
// * `AA`        - AM or PM
// * `Z`         - hour and minute offset from UTC, or Z for UTC
// * `ZZZ`       - hour and minute offset from UTC
//
// Timezone should be a location name as specified in the IANA Time Zone database, such
// as "America/Guayaquil" or "America/Los_Angeles". If not specified the timezone of your
// environment will be used. An error will be returned if the timezone is not recognized.
//
//   @(format_datetime("1979-07-18T15:00:00.000000Z")) -> 1979-07-18 10:00
//   @(format_datetime("1979-07-18T15:00:00.000000Z", "YYYY-MM-DD")) -> 1979-07-18
//   @(format_datetime("2010-05-10T19:50:00.000000Z", "YYYY M DD tt:mm")) -> 2010 5 10 14:50
//   @(format_datetime("2010-05-10T19:50:00.000000Z", "YYYY-MM-DD tt:mm AA", "America/Los_Angeles")) -> 2010-05-10 12:50 PM
//   @(format_datetime("1979-07-18T15:00:00.000000Z", "YYYY")) -> 1979
//   @(format_datetime("1979-07-18T15:00:00.000000Z", "M")) -> 7
//   @(format_datetime("NOT DATE", "YYYY-MM-DD")) -> ERROR
//
// @function format_datetime(date, format [,timezone])
func FormatDateTime(env utils.Environment, args ...types.XValue) types.XValue {
	if len(args) < 1 || len(args) > 3 {
		return types.NewXErrorf("takes one or two arguments, got %d", len(args))
	}
	date, xerr := types.ToXDateTime(env, args[0])
	if xerr != nil {
		return xerr
	}

	format := types.NewXText(fmt.Sprintf("%s %s", env.DateFormat().String(), env.TimeFormat().String()))
	if len(args) >= 2 {
		format, xerr = types.ToXText(env, args[1])
		if xerr != nil {
			return xerr
		}
	}

	// try to turn it to a go format
	goFormat, err := utils.ToGoDateFormat(format.Native())
	if err != nil {
		return types.NewXError(err)
	}

	// grab our location
	location := env.Timezone()
	if len(args) == 3 {
		arg3, xerr := types.ToXText(env, args[2])
		if xerr != nil {
			return xerr
		}

		location, err = time.LoadLocation(arg3.Native())
		if err != nil {
			return types.NewXError(err)
		}
	}

	// convert to our timezone if we have one (otherwise we remain in the date's default)
	if location != nil {
		date = types.NewXDateTime(date.Native().In(location))
	}

	// return the formatted date
	return types.NewXText(date.Native().Format(goFormat))
}

// FormatNumber returns `num` formatted with the passed in number of decimal `places` and optional `commas` dividing thousands separators
//
//   @(format_number(31337)) -> 31,337.00
//   @(format_number(31337, 2)) -> 31,337.00
//   @(format_number(31337, 2, true)) -> 31,337.00
//   @(format_number(31337, 0, false)) -> 31337
//   @(format_number("foo", 2, false)) -> ERROR
//
// @function format_number(num, places, commas)
func FormatNumber(env utils.Environment, args ...types.XValue) types.XValue {
	if len(args) < 1 || len(args) > 3 {
		return types.NewXErrorf("takes 1 to 3 arguments, got %d", len(args))
	}

	num, err := types.ToXNumber(env, args[0])
	if err != nil {
		return err
	}

	places := 2
	if len(args) > 1 {
		if places, err = types.ToInteger(env, args[1]); err != nil {
			return err
		}
		if places < 0 || places > 9 {
			return types.NewXErrorf("must take 0-9 number of places, got %d", args[1])
		}
	}

	commas := types.XBooleanTrue
	if len(args) > 2 {
		if commas, err = types.ToXBoolean(env, args[2]); err != nil {
			return err
		}
	}

	// build our format string
	formatStr := bytes.Buffer{}
	if commas.Native() {
		formatStr.WriteString("#,###.")
	} else {
		formatStr.WriteString("####.")
	}
	if places > 0 {
		for i := 0; i < places; i++ {
			formatStr.WriteString("#")
		}
	}
	f64, _ := num.Native().Float64()
	return types.NewXText(humanize.FormatFloat(formatStr.String(), f64))
}

// FormatLocation formats the given location as its name
//
//   @(format_location("Rwanda")) -> Rwanda
//   @(format_location("Rwanda > Kigali")) -> Kigali
//
// @function format_location(location)
func FormatLocation(env utils.Environment, path types.XText) types.XValue {
	parts := strings.Split(path.Native(), ">")
	return types.NewXText(strings.TrimSpace(parts[len(parts)-1]))
}

// FormatURN turns `urn` into human friendly text
//
//   @(format_urn("tel:+250781234567")) -> 0781 234 567
//   @(format_urn("twitter:134252511151#billy_bob")) -> billy_bob
//   @(format_urn(contact.urns)) -> (206) 555-1212
//   @(format_urn(contact.urns.2)) -> foo@bar.com
//   @(format_urn(contact.urns.mailto)) -> foo@bar.com
//   @(format_urn(contact.urns.mailto.0)) -> foo@bar.com
//   @(format_urn(contact.urns.telegram)) ->
//   @(format_urn("NOT URN")) -> ERROR
//
// @function format_urn(urn)
func FormatURN(env utils.Environment, args ...types.XValue) types.XValue {
	if len(args) != 1 {
		return types.NewXErrorf("takes one argument, got %d", len(args))
	}

	// if we've been passed an indexable like a URNList, use first item
	urnArg := args[0]

	indexable, isIndexable := urnArg.(types.XIndexable)
	if isIndexable {
		if indexable.Length() >= 1 {
			urnArg = indexable.Index(0)
		} else {
			return types.XTextEmpty
		}
	}

	urnString, xerr := types.ToXText(env, urnArg)
	if xerr != nil {
		return xerr
	}

	urn := urns.URN(urnString.Native())
	err := urn.Validate()
	if err != nil {
		return types.NewXErrorf("%s is not a valid URN: %s", urnString, err)
	}

	return types.NewXText(urn.Format())
}

//------------------------------------------------------------------------------------------
// Utility Functions
//------------------------------------------------------------------------------------------

// Length returns the length of the passed in text or array.
//
// length will return an error if it is passed an item which doesn't have length.
//
//   @(length("Hello")) -> 5
//   @(length("😀😃😄😁")) -> 4
//   @(length(array())) -> 0
//   @(length(array("a", "b", "c"))) -> 3
//   @(length(1234)) -> ERROR
//
// @function length(value)
func Length(env utils.Environment, value types.XValue) types.XValue {
	// argument must be a value with length
	lengthable, isLengthable := value.(types.XLengthable)
	if isLengthable {
		return types.NewXNumberFromInt(lengthable.Length())
	}

	return types.NewXErrorf("value doesn't have length")
}

// Default takes two arguments, returning `test` if not an error or nil or empty text, otherwise returning `default`
//
//   @(default(undeclared.var, "default_value")) -> default_value
//   @(default("10", "20")) -> 10
//   @(default("", "value")) -> value
//   @(default(array(1, 2), "value")) -> ["1","2"]
//   @(default(array(), "value")) -> value
//   @(default(datetime("invalid-date"), "today")) -> today
//
// @function default(test, default)
func Default(env utils.Environment, test types.XValue, def types.XValue) types.XValue {
	if types.IsEmpty(test) || types.IsXError(test) {
		return def
	}

	return test
}

// LegacyAdd simulates our old + operator, which operated differently based on whether
// one of the parameters was a date or not. If one is a date, then the other side is
// expected to be an integer with a number of days to add to the date, otherwise a normal
// decimal addition is attempted.
func LegacyAdd(env utils.Environment, arg1 types.XValue, arg2 types.XValue) types.XValue {

	// try to parse dates and decimals
	date1, date1Err := types.ToXDateTime(env, arg1)
	date2, date2Err := types.ToXDateTime(env, arg2)

	dec1, dec1Err := types.ToXNumber(env, arg1)
	dec2, dec2Err := types.ToXNumber(env, arg2)

	// if they are both dates, that's an error
	if date1Err == nil && date2Err == nil {
		return types.NewXErrorf("cannot operate on two dates")
	}

	// date and int, do a day addition
	if date1Err == nil && dec2Err == nil {
		if dec2.Native().IntPart() < math.MinInt32 || dec2.Native().IntPart() > math.MaxInt32 {
			return types.NewXErrorf("cannot operate on integers greater than 32 bit")
		}
		return types.NewXDateTime(date1.Native().AddDate(0, 0, int(dec2.Native().IntPart())))
	}

	// int and date, do a day addition
	if date2Err == nil && dec1Err == nil {
		if dec1.Native().IntPart() < math.MinInt32 || dec1.Native().IntPart() > math.MaxInt32 {
			return types.NewXErrorf("cannot operate on integers greater than 32 bit")
		}
		return types.NewXDateTime(date2.Native().AddDate(0, 0, int(dec1.Native().IntPart())))
	}

	// one of these doesn't look like a valid decimal either, bail
	if dec1Err != nil {
		return types.NewXError(dec1Err)
	}

	if dec2Err != nil {
		return types.NewXError(dec2Err)
	}

	// normal decimal addition
	return types.NewXNumber(dec1.Native().Add(dec2.Native()))
}

// ReadChars converts `text` into something that can be read by IVR systems
//
// ReadChars will split the numbers such as they are easier to understand. This includes
// splitting in 3s or 4s if appropriate.
//
//   @(read_chars("1234")) -> 1 2 3 4
//   @(read_chars("abc")) -> a b c
//   @(read_chars("abcdef")) -> a b c , d e f
//
// @function read_chars(text)
func ReadChars(env utils.Environment, val types.XText) types.XValue {
	var output bytes.Buffer

	// remove any leading +
	val = types.NewXText(strings.TrimLeft(val.Native(), "+"))

	length := val.Length()

	// groups of three
	if length%3 == 0 {
		// groups of 3
		for i := 0; i < length; i += 3 {
			if i > 0 {
				output.WriteString(" , ")
			}
			output.WriteString(strings.Join(strings.Split(val.Native()[i:i+3], ""), " "))
		}
		return types.NewXText(output.String())
	}

	// groups of four
	if length%4 == 0 {
		for i := 0; i < length; i += 4 {
			if i > 0 {
				output.WriteString(" , ")
			}
			output.WriteString(strings.Join(strings.Split(val.Native()[i:i+4], ""), " "))
		}
		return types.NewXText(output.String())
	}

	// default, just do one at a time
	for i, c := range val.Native() {
		if i > 0 {
			output.WriteString(" , ")
		}
		output.WriteRune(c)
	}

	return types.NewXText(output.String())
}
