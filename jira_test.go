package table

import (
	"fmt"
	"testing"
)

func TestJira(t *testing.T) {
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
			expectedHeader: "|| aaa || b    || ccccc ||",
			expectedBody:   "\n|  d   |  eeee |  f      |",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%sHeader", test.name), func(t *testing.T) {
			o := jira{Content: test.input, Options: Options{}}
			got := o.header()
			if got != test.expectedHeader {
				t.Errorf("TestJiraHeader: got: %s, expected: %s", got, test.expectedHeader)
			}
		})
		t.Run(fmt.Sprintf("%sBody", test.name), func(t *testing.T) {
			o := jira{Content: test.input, Options: Options{}}
			got := o.record()
			if got != test.expectedBody {
				t.Errorf("TestJiraBody: got: %s, expected: %s", got, test.expectedBody)
			}
		})
	}
}
