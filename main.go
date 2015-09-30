package main

import (
	"flag"
	"fmt"
	"github.com/rhysd/api-dts/apidts"
	"os"
)

func ParseArgv() string {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	vf := fs.Bool("version", false, "Display version")
	of := fs.String("out", "", "Output file name")
	fs.Parse(os.Args[1:])

	if *vf {
		fmt.Println("0.0.0")
		os.Exit(0)
	}

	return *of
}

func main() {
	out := ParseArgv()

	var err error
	dts, err := apidts.ConvertJsonToDts(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	stringized, err := apidts.StringizeDts(dts, out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if out != "" {
		f, err := os.Create(out)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		_, err = f.WriteString(stringized)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Println(stringized)
	}
}
