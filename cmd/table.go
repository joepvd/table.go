package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/jessevdk/go-flags"
	"github.com/joepvd/table.go"
)

var opts struct {
	Style string `short:"s" long:"style"`
	FS    string `short:"F" long:"field-seperator" default:"[ \t]+"`
}
var files []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	var err error
	files, err = flags.Parse(&opts)
	check(err)
}

func main() {
	var fileHandle *os.File
	var err error

	switch len(files) {
	case 0:
		fileHandle = os.Stdin
		break
	case 1:
		fileHandle, err = os.Open(files[0])
		check(err)
		break
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}
	defer fileHandle.Close()

	fs := regexp.MustCompile(opts.FS)
	contents := table.ParseText(fileHandle, fs)
	fmt.Printf("%#v\n", contents)

}
