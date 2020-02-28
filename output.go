package table

import (
	"errors"
	"fmt"
)

// Output is entry function for outputting the table
func Output(c Content, o Options) (string, error) {
	var e error
	var outputter format
	switch style := o.Style; style {
	case "md":
		outputter = markdownOutputter{Content: c, Options: o}
	case "psql":
		outputter = psqlOutputter{Content: c, Options: o}
	case "jira":
		outputter = jiraOutputter{Content: c, Options: o}
	default:
		e = errors.New("Outputter not implemented error")
	}
	h := outputter.header()
	b := outputter.record()
	return fmt.Sprintf("%s%s", h, b), e
}

type format interface {
	header() string
	record() string
}
