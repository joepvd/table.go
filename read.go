package table

import (
	"bufio"
	"fmt"
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
	Records []LineRecord
}

func parseLine(s string, re *regexp.Regexp) LineRecord {
	fields := re.Split(s, -1)
	return LineRecord{
		len(fields),
		fields,
	}
}

// ParseText reads file and converts it to Content
func ParseText(file *os.File, re *regexp.Regexp) Content {
	reader := bufio.NewScanner(file)
	var content Content

	for reader.Scan() {
		line := reader.Text()
		splitted := parseLine(line, re)
		content.Records = append(content.Records, splitted)
		content.NR = content.NR + 1
		content.MaxFS = max(content.MaxFS, splitted.NF)
	}

	if reader.Err() != nil {
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
