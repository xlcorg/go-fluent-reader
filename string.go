package fluent_reader

import (
	"strconv"
	"unicode"
	"unicode/utf8"
)

type String string

func (s String) FluentReader() *FluentReader {
	return NewFluentReader(string(s))
}

func (s String) Next() rune {
	ch, _ := utf8.DecodeRuneInString(string(s))
	return ch
}

func (s String) Last() rune {
	ch, _ := utf8.DecodeLastRuneInString(string(s))
	return ch
}

func (s String) String() string {
	return string(s)
}

func (s String) Length() int {
	return len(s)
}

func (s String) Empty() bool {
	return len(s) == 0
}

func (s String) HasLettersOnly() bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range string(s) {
		if !isLetter(c) {
			return false
		}
	}
	return true
}

func (s String) HasDigitsOnly() bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range string(s) {
		if !isDigit(c) {
			return false
		}
	}
	return true
}

func (s String) Int() int {
	result, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return result
}

func (s String) Int64() int64 {
	result, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func isLetter(c rune) bool {
	return unicode.IsLetter(c)
}

func isDigit(c rune) bool {
	return unicode.IsDigit(c)
}
