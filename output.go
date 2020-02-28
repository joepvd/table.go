package table

import (
	"errors"
)

// Output is entry function for outputting the table
func Output(c Content, o Options) (string, error) {
	var e error
	var formatter format
	switch style := o.Style; style {
	case "md":
		formatter = md{Content: c, Options: o}
	case "psql":
		formatter = psql{Content: c, Options: o}
	case "jira":
		formatter = jira{Content: c, Options: o}
	default:
		e = errors.New("Formatter not implemented error")
	}
	table := formatter.format()
	return table, e
}

type format interface {
	format() string
}
