package table

import (
	"fmt"
)

type jira struct {
	Content Content
	Options Options
}

func (m jira) format() string {
	return fmt.Sprintf("%s%s", m.header(), m.body())
}

func jiraTxt(s []string, widths []int, mode string) (out string) {
	var cell string
	c := "| "
	if mode == "header" {
		c = "||"
	}
	for i, w := range widths {
		f := fmt.Sprintf("%%-%ds", w)
		cell = fmt.Sprintf(f, s[i])
		out = fmt.Sprintf("%s%s %s ", out, c, cell)
	}
	if c == "| " {
		c = " |"
	}
	out = fmt.Sprintf("%s%s", out, c)
	return out
}

func (m jira) header() string {
	widths := m.Content.Widths
	txt := jiraTxt(m.Content.Records[0].Fields, widths, "header")
	return txt
}

func (m jira) body() (out string) {
	widths := m.Content.Widths
	for i := 1; i < m.Content.NR; i++ {
		fields := m.Content.Records[i].Fields
		txt := jiraTxt(fields, widths, "col")
		out = fmt.Sprintf("%s\n%s", out, txt)
	}
	return out
}
