package table

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

// LineRecord stores stuff from a line
type LineRecord struct {
	NF     int
	Fields []string
}

// Content stores content of the table. Contains LineRecords
type Content struct {
	NR      int
	MaxFS   int
	Widths  []int
	Records []LineRecord
}

func parseLine(s string, re *regexp.Regexp) (LineRecord, bool) {
	var nf int
	ok := false
	reStart := regexp.MustCompile(fmt.Sprintf("^%s", re.String()))
	s = reStart.ReplaceAllString(s, "")

	reEnd := regexp.MustCompile(fmt.Sprintf("%s$", re.String()))
	s = reEnd.ReplaceAllString(s, "")

	fields := re.Split(s, -1)
	for _, v := range fields {
		if len(v) > 0 {
			ok = true
		}
	}
	if ok {
		nf = len(fields)
	} else {
		nf = 0
	}

	return LineRecord{
		nf,
		fields,
	}, ok
}

// ParseFile passes file handle on to ParseText
func ParseFile(file *os.File, re *regexp.Regexp) Content {
	reader := bufio.NewReader(file)
	return ParseText(reader, re)
}

// ParseText takes a reader and converts it to Content
func ParseText(reader io.Reader, re *regexp.Regexp) Content {
	var content Content
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		splitted, ok := parseLine(line, re)
		if ok {
			content.Records = append(content.Records, splitted)
			content.NR = content.NR + 1
			content.MaxFS = max(content.MaxFS, splitted.NF)
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error")
	}
	return content
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
