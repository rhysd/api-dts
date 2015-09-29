package main

import (
	"flag"
	"fmt"
	"github.com/rhysd/api-dts/apidts"
	"io"
	"os"
)

func ParseArgv() io.Reader {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	vf := fs.Bool("version", false, "Display version")
	ff := fs.String("file", "", "Read from file")
	fs.Parse(os.Args[1:])

	if *vf {
		fmt.Println("0.0.0")
		os.Exit(0)
	}

	if len(*ff) != 0 {
		f, err := os.Open(*ff)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return f
	}

	return os.Stdin
}

func main() {
	target := ParseArgv()
	dts := apidts.ConvertJsonToDts(target)
	stringized := apidts.StringizeDts(dts)
	fmt.Println(stringized)
}
