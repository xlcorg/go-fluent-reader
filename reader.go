package fluent_reader

import (
	"strings"
	"unicode/utf8"
)

type FluentReader struct {
	index  int
	source string
}

func NewFluentReader(s string) *FluentReader {
	return &FluentReader{
		source: s,
		index:  0,
	}
}

func (s *FluentReader) Next() rune {
	ch, _ := utf8.DecodeRuneInString(s.source[s.index:])
	return ch
}

func (s *FluentReader) Last() rune {
	ch, _ := utf8.DecodeLastRuneInString(s.source[s.index:])
	return ch
}

func (s *FluentReader) HasNext() bool {
	return s.index < len(s.source)
}

func (s *FluentReader) HasNextChar(r rune) bool {
	return s.Next() == r
}

func (s *FluentReader) Empty() bool {
	return s.index == len(s.source)
}

func (s *FluentReader) Clone() *FluentReader {
	return &FluentReader{
		source: s.source,
		index:  s.index,
	}
}

func (s *FluentReader) ReadAll() String {
	return String(s.source[s.index:])
}

func (s *FluentReader) String() string {
	return s.source[s.index:]
}

func (s *FluentReader) ReadUntil(ch rune) (result String) {
	found := strings.IndexRune(s.source[s.index:], ch)
	if found == -1 {
		result = String(s.source[s.index:])
		s.index = len(s.source)
	} else {
		newIndex := found + s.index
		result = String(s.source[s.index:newIndex])
		s.index = newIndex
	}
	return
}

func (s *FluentReader) SkipOne() *FluentReader {
	_, size := utf8.DecodeRuneInString(s.source[s.index:])
	s.index = min(s.index+size, len(s.source))
	return s
}

func (s *FluentReader) Skip(count int) *FluentReader {
	for i := 0; i < count && s.HasNext(); i++ {
		s.SkipOne()
	}
	return s
}

func (s *FluentReader) SkipUntil(ch rune) *FluentReader {
	found := strings.IndexRune(s.source[s.index:], ch)
	if found == -1 {
		s.index = len(s.source)
	} else {
		s.index += found
	}
	return s
}

func (s *FluentReader) SkipAfter(ch rune) *FluentReader {
	found := strings.IndexRune(s.source[s.index:], ch)
	if found == -1 {
		s.index = len(s.source)
	} else {

		s.index = min(s.index+found, len(s.source))
		s.SkipOne()
	}
	return s
}
