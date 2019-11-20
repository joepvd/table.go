package main

import (
	"flag"
	"fmt"
	"github.com/joepvd/table.go"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()
	var err error
	var fileHandle *os.File

	switch flag.NArg() {
	case 0:
		fileHandle = os.Stdin
		break
	case 1:
		fileHandle, err = os.Open(flag.Arg(0))
		check(err)
		break
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}
	defer fileHandle.Close()

	fs := regexp.MustCompile(`[ \t]+`)
	contents := table.ParseText(fileHandle, fs)
	fmt.Printf("%#v\n", contents)
}
