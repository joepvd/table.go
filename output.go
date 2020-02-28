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
		outputter = md{Content: c, Options: o}
	case "psql":
		outputter = psql{Content: c, Options: o}
	case "jira":
		outputter = jira{Content: c, Options: o}
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
