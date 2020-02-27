package table

import (
	"fmt"
)

// What do I need to do?
// Want to have a basic type here.
// The output type.
// which ensures that some methods are implemented:
// output.writeHeader()
// output.writeRecord()
//
// and a general method output.writeTable()
// hmmmmm.... maybe this should not operate on output, but on Content?
// Meaning... Need to... include `opts` in Content. !?!
//
// depending on style, a different concrete function should be selected.
//
// Need to have something like a dispatch table?

// Output is entry function for outputting the table
func Output(c Content, o Options) string {
	var outputter format
	if o.Style == "md" {
		outputter = markdownOutputter{Content: c, Options: o}
	}

	h := outputter.header()
	b := outputter.record()
	return fmt.Sprintf("%s%s", h, b)
}

type format interface {
	header() string
	record() string
}
