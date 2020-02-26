package table

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestMax(t *testing.T) {
	// Really just a test if the testing works :)
	if max(1, 3) != 3 {
		t.Error()
	}
}

func TestParseText(t *testing.T) {
	r := regexp.MustCompile(" +")

	var parseTests = []struct {
		name     string
		in       string
		expected Content
	}{
		{"normal case", "a b\nc    d", Content{NR: 2, MaxFS: 2}},
		{"leading spaces", "123 234\n    456 789", Content{NR: 2, MaxFS: 2}},
		{"empty lines", "abc def\n\nghi jkl\n", Content{NR: 2, MaxFS: 2}},
	}

	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			f := strings.NewReader(tt.in)
			content := ParseText(f, r)
			if content.NR != tt.expected.NR {
				t.Errorf("ParseText(%s) wrong NR. Got %v, want %v\ninput text:\n%v\n", tt.name, content.NR, tt.expected.NR, tt.in)
			}
			if content.MaxFS != tt.expected.MaxFS {
				t.Errorf("ParseText(%s) wrong MaxFS. Got %v, want %v\ninput text:\n%v\n", tt.name, content.MaxFS, tt.expected.MaxFS, tt.in)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	r := regexp.MustCompile(" +")
	var lineTests = []struct {
		name     string
		in       string
		expected LineRecord
	}{
		{"normal case", "a b c", LineRecord{NF: 3, Fields: []string{"a", "b", "c"}}},
		{"trailing seperator", "1 2 3 ", LineRecord{NF: 3, Fields: []string{"1", "2", "3"}}},
		{"starts with seperator", " 1 2 3", LineRecord{NF: 3, Fields: []string{"1", "2", "3"}}},
	}
	for _, test := range lineTests {
		t.Run(test.name, func(t *testing.T) {
			result := true
			got, _ := parseLine(test.in, r)

			if got.NF != test.expected.NF {
				result = false
			} else if len(got.Fields) != len(test.expected.Fields) {
				fmt.Printf("len(got.Fields) = %v, len(test.expected.Fields) = %v\n", len(got.Fields), len(test.expected.Fields))
				result = false
			} else {
				for k, v := range got.Fields {
					if v != test.expected.Fields[k] {
						result = false
					}
				}
			}

			if !result {
				t.Errorf("parseLine(%s). Input: <%s>. Got: <%v>. Expected: <%v>", test.name, test.in, got, test.expected)
			}
		})
	}
}
