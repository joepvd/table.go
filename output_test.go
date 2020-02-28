package table

import (
	"testing"
)

var input = Content{
	NR:     2,
	MaxFS:  3,
	Widths: []int{3, 4, 5},
	Records: []LineRecord{
		{NF: 3, Fields: []string{"aaa", "b", "ccccc"}},
		{NF: 3, Fields: []string{"d", "eeee", "f"}},
	},
}

func TestOutputTypes(t *testing.T) {
	var tests = []struct {
		name           string
		input          Content
		expectedHeader string
		expectedBody   string
	}{
		{
			name:           "jira",
			input:          input,
			expectedHeader: "|| aaa || b    || ccccc ||",
			expectedBody:   "\n|  d   |  eeee |  f      |",
		},
		{
			name:           "md",
			input:          input,
			expectedHeader: "| aaa | b    | ccccc |\n| --- | ---- | ----- |\n",
			expectedBody:   "| d   | eeee | f     |\n",
		},
		{
			name:           "psql",
			input:          input,
			expectedHeader: "┌─────┬──────┬───────┐\n│ aaa │ b    │ ccccc │\n├─────┼──────┼───────┤",
			expectedBody:   "\n│ d   │ eeee │ f     │\n└─────┴──────┴───────┘",
		},
	}

	for _, test := range tests {
		var got string
		var o format
		switch test.name {
		case "jira":
			o = jira{Content: test.input, Options: Options{}}
		case "md":
			o = md{Content: test.input, Options: Options{}}
		case "psql":
			o = psql{Content: test.input, Options: Options{}}
		default:
			panic("Not implemented!")
		}
		t.Run(test.name, func(t *testing.T) {
			t.Run("Header", func(t *testing.T) {
				got = o.header()
				if got != test.expectedHeader {
					t.Errorf("%sHeader: got: %s, expected: %s", test.name, got, test.expectedHeader)
				}
			})
			t.Run("Body", func(t *testing.T) {
				got := o.body()
				if got != test.expectedBody {
					t.Errorf("%sBody: got: %s, expected: %s", test.name, got, test.expectedBody)
				}
			})
		})
	}
}

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
