package table

import (
	"fmt"
	"strings"
)

type psql struct {
	Content Content
	Options Options
}

func (m psql) format() string {
	return fmt.Sprintf("%s%s", m.header(), m.body())
}

func (m psql) header() string {
	widths := m.Content.Widths
	top := psqlFiller(widths, []string{"┌", "─", "┬", "┐"})
	txt := psqlTxt(m.Content.Records[0].Fields, widths)
	sep := psqlFiller(widths, []string{"├", "─", "┼", "┤"})
	return fmt.Sprintf("%s\n%s\n%s", top, txt, sep)
}

func (m psql) body() (out string) {
	widths := m.Content.Widths
	for i := 1; i < m.Content.NR; i++ {
		record := m.Content.Records[i]
		txt := psqlTxt(record.Fields, widths)
		out = fmt.Sprintf("%s\n%s", out, txt)
	}
	bottom := psqlFiller(widths, []string{"└", "─", "┴", "┘"})
	return fmt.Sprintf("%s\n%s", out, bottom)
}

func psqlTxt(s []string, widths []int) (out string) {
	var cell string
	for i, w := range widths {
		c := "│"
		if i == 0 {
			c = "│"
		}
		f := fmt.Sprintf("%%-%ds", w)
		word := ""
		if i < len(s) {
			word = s[i]
		}
		cell = fmt.Sprintf(f, word)
		out = fmt.Sprintf("%s%s %s ", out, c, cell)
	}
	out = fmt.Sprintf("%s│", out)
	return out
}

func psqlFiller(widths []int, seps []string) (out string) {
	for i, l := range widths {
		c := seps[2]
		if i == 0 {
			c = seps[0]
		}
		out = fmt.Sprintf("%s%s%s", out, c, strings.Repeat(seps[1], l+2))
	}
	out = fmt.Sprintf("%s%s", out, seps[3])
	return out
}
