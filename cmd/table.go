package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/jessevdk/go-flags"
	table "github.com/joepvd/table.go"
)

var files []string
var opts table.Options

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
	case 1:
		fileHandle, err = os.Open(files[0])
		check(err)
	default:
		fmt.Printf("input must be from stdin or single file\n")
		os.Exit(1)
	}
	defer fileHandle.Close()

	fs := regexp.MustCompile(opts.FS)
	contents := table.ParseText(fileHandle, fs)
	fmt.Printf("%s", table.Output(contents, opts))
}
