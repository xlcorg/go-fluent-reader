package benchmark

import (
	"regexp"
	"testing"

	"github.com/xlcorg/go-fluent-reader"
)

const input = "gnsbgyг:200260797,jsiczw:434278153,qkuyencyih:88313277"

func BenchmarkValidateFluentReader(b *testing.B) {
	b.ReportAllocs()
	reader := fluent_reader.NewFluentReader(input)
	for i := 0; i < b.N; i++ {
		validateResult(reader.Clone())
	}
}

func BenchmarkValidateRegex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		validateResultRegex(input)
	}
}

var regex = regexp.MustCompile("^((([a-z]+:{1}[1-9][0-9]*),?)+).*[^,]$")

func validateResultRegex(s string) bool {
	matched := regex.FindAllString(s, len(s))
	var counter int
	for _, s := range matched {
		counter += len(s)
	}
	return counter == len(s)
}

func validateResult(reader *fluent_reader.FluentReader) bool {
	if reader.Next() == ',' || reader.Last() == ',' {
		return false
	}

	for reader.HasNext() {
		key := reader.ReadUntil(':')
		if !key.HasLettersOnly() {
			return false
		}

		val := reader.SkipOne().ReadUntil(',')
		if !val.HasDigitsOnly() || val.Next() == '0' {
			return false
		}

		reader.SkipOne()
	}

	return true
}
