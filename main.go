package main

import (
	"flag"
	"fmt"
	"github.com/rhysd/api-dts/apidts"
	"os"
)

func ParseArgv() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	vf := fs.Bool("version", false, "Display version")
	fs.Parse(os.Args[1:])

	if *vf {
		fmt.Println("0.0.0")
		os.Exit(0)
	}
}

func main() {
	ParseArgv()

	var err error
	dts, err := apidts.ConvertJsonToDts(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	stringized, err := apidts.StringizeDts(dts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(stringized)
}
