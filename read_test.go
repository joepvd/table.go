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
		{
			name:     "normal case",
			in:       "a b\nc    d",
			expected: Content{NR: 2, MaxFS: 2}},
		{
			name:     "leading spaces",
			in:       "123 234\n    456 789",
			expected: Content{NR: 2, MaxFS: 2}},
		{
			name:     "empty lines",
			in:       "abc def\n\nghi jkl\n",
			expected: Content{NR: 2, MaxFS: 2}},
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
		{
			name:     "normal case",
			in:       "a b c",
			expected: LineRecord{NF: 3, Fields: []string{"a", "b", "c"}}},
		{
			name:     "trailing seperator",
			in:       "1 2 3 ",
			expected: LineRecord{NF: 3, Fields: []string{"1", "2", "3"}}},
		{
			name:     "starts with seperator",
			in:       " 1 2 3",
			expected: LineRecord{NF: 3, Fields: []string{"1", "2", "3"}}},
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

func TestSetWidths(t *testing.T) {

	tests := []struct {
		name     string
		line     LineRecord
		widths   []int
		expected []int
	}{
		{
			name:     "initial case",
			line:     LineRecord{NF: 3, Fields: []string{"abc", "with some spaces", "owyeah"}},
			widths:   []int{0, 0, 0},
			expected: []int{3, 16, 6},
		},
		{
			name:     "max from initial widths",
			line:     LineRecord{NF: 3, Fields: []string{"abc", "with some spaces", "owyeah"}},
			widths:   []int{20, 20, 20},
			expected: []int{20, 20, 20},
		},
		{
			name:     "growing NF",
			line:     LineRecord{NF: 3, Fields: []string{"abc", "with some spaces", "owyeah"}},
			widths:   []int{0, 0},
			expected: []int{3, 16, 6},
		},
	}

	for _, record := range tests {
		t.Run(record.name, func(t *testing.T) {
			c := Content{
				NR:      1,
				MaxFS:   len(record.widths),
				Widths:  record.widths,
				Records: []LineRecord{},
			}
			ok := true
			c.setWidths(record.line)
			for i, j := range record.expected {
				if j != c.Widths[i] {
					ok = false
				}
			}
			if len(record.expected) != len(c.Widths) {
				ok = false
			}
			if !ok {
				t.Errorf("TestSetWidths(%s): got <%v>, expected <%v>, with input <%#v>", record.name, c.Widths, record.expected, record.line.Fields)
			}
		})
	}
}
