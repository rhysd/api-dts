package main

import (
	"flag"
	"fmt"
	"github.com/rhysd/api-dts/apidts"
	"io"
	"os"
)

func ParseArgv() (string, io.Reader) {
	fs := flag.NewFlagSet(fmt.Sprintf("%s [file]", os.Args[0]), flag.ExitOnError)
	vf := fs.Bool("version", false, "Display version")
	fs.Parse(os.Args[1:])

	if *vf {
		fmt.Println("0.0.0")
		os.Exit(0)
	}

	args := fs.Args()
	if len(args) == 0 {
		return "", os.Stdin
	}

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return args[0], f
}

func main() {
	file_name, input := ParseArgv()

	var err error
	dts, err := apidts.ConvertJsonToDts(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	stringized, err := apidts.StringizeDts(dts, file_name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(stringized)
}
