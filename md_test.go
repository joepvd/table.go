package table

import (
	"fmt"
	"testing"
)

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
		expectedHeader: "| aaa | b    | ccccc |\n| --- | ---- | ----- |\n",
		expectedBody:   "| d   | eeee | f     |\n",
	},
}

func TestMarkdown(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("%sHeader", test.name), func(t *testing.T) {
			o := md{Content: test.input, Options: Options{}}
			got := o.header()
			if got != test.expectedHeader {
				t.Errorf("TestMarkdownHeader: got: %s, expected: %s", got, test.expectedHeader)
			}
		})
		t.Run(fmt.Sprintf("%sBody", test.name), func(t *testing.T) {
			o := md{Content: test.input, Options: Options{}}
			got := o.body()
			if got != test.expectedBody {
				t.Errorf("TestMarkdownBody: got: %s, expected: %s", got, test.expectedBody)
			}
		})
	}
}
