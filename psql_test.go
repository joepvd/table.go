package table

import (
	"fmt"
	"testing"
)

func TestPsqlFiller(t *testing.T) {
	var tests = []struct {
		name     string
		seps     []string
		dims     []int
		expected string
	}{
		{
			name:     "top",
			seps:     []string{"┌", "─", "┬", "┐"},
			dims:     []int{3, 4, 5},
			expected: "┌─────┬──────┬───────┐",
		},
		{
			name:     "sep",
			seps:     []string{"├", "─", "┼", "┤"},
			dims:     []int{3, 4, 5},
			expected: "├─────┼──────┼───────┤",
		},
		{
			name:     "bottom",
			seps:     []string{"└", "─", "┴", "┘"},
			dims:     []int{3, 4, 5},
			expected: "└─────┴──────┴───────┘",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := psqlFiller(test.dims, test.seps)
			if got != test.expected {
				t.Errorf("psqlTop:%s. Got: %s. Expected: %s", test.name, got, test.expected)
			}
		})
	}
}

func TestPsqlTxt(t *testing.T) {
	expected := "│ a   │ aaa  │ aaaaa │"
	got := psqlTxt([]string{"a", "aaa", "aaaaa"}, []int{3, 4, 5})
	if got != expected {
		t.Errorf("psqlTxt. Got: <%s>, expected: <%s>", got, expected)
	}
}

func TestPsql(t *testing.T) {
	var tests = []struct {
		name           string
		input          Content
		expectedHeader string
		expectedBody   string
	}{
		{
			name: "simple",
			input: Content{
				NR:     2,
				MaxFS:  3,
				Widths: []int{3, 4, 5},
				Records: []LineRecord{
					{NF: 3, Fields: []string{"aaa", "b", "ccccc"}},
					{NF: 3, Fields: []string{"d", "eeee", "f"}},
				},
			},
			expectedHeader: "┌─────┬──────┬───────┐\n│ aaa │ b    │ ccccc │\n├─────┼──────┼───────┤",
			expectedBody:   "\n│ d   │ eeee │ f     │\n└─────┴──────┴───────┘",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%sHeader", test.name), func(t *testing.T) {
			o := psqlOutputter{Content: test.input, Options: Options{}}
			got := o.header()
			if got != test.expectedHeader {
				t.Errorf("TestPsqlHeader: got: %s, expected: %s", got, test.expectedHeader)
			}
		})
		t.Run(fmt.Sprintf("%sBody", test.name), func(t *testing.T) {
			o := psqlOutputter{Content: test.input, Options: Options{}}
			got := o.record()
			if got != test.expectedBody {
				t.Errorf("TestPsqlBody: got: %s, expected: %s", got, test.expectedBody)
			}
		})
	}
}
