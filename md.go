package table

import (
	"fmt"
	"strings"
)

type md struct {
	Content Content
	Options Options
}

func (m md) format() string {
	return fmt.Sprintf("%s%s", m.header(), m.body())
}

func (m md) header() string {
	widths := m.Content.Widths
	txt := m.fline(m.Content.Records[0].Fields, widths)
	sep := m.fline(mdSeperator(widths), widths)
	out := fmt.Sprintf("%s%s", txt, sep)
	return out
}

func (m md) body() string {
	var out string
	for i := 1; i < m.Content.NR; i++ {
		record := m.Content.Records[i]
		txt := m.fline(record.Fields, m.Content.Widths)
		out = fmt.Sprintf("%s%s", out, txt)
	}
	return out
}

func (m md) fline(s []string, widths []int) (out string) {
	var cell string
	for i, w := range widths {
		f := fmt.Sprintf("%%-%ds", w)
		cell = fmt.Sprintf(f, s[i])
		out = fmt.Sprintf("%s| %s ", out, cell)
	}
	out = fmt.Sprintf("%s|\n", out)
	return out
}

func mdSeperator(widths []int) (out []string) {
	for _, w := range widths {
		out = append(out, strings.Repeat("-", w))
	}
	return out
}
