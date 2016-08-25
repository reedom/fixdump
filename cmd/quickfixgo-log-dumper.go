package main

import (
	"flag"
	"fmt"
	"os"

	dumper "github.com/reedom/quickfixgo-log-dumper"
)

const version = "0.0.1"

func init() {
	ver := flag.Bool("version", false, "print version")
	flag.Usage = usage
	flag.Parse()
	if *ver {
		fmt.Fprintf(os.Stdout, "%v version %s\n", os.Args[0], version)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %v [flags] <path to log file> ... \n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	if flag.NArg() == 0 {
		dumper.Dump(os.Stdin)
		return
	}

	for _, logFilePath := range flag.Args() {
		file, err := os.Open(logFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file: %s\n", err.Error())
			continue
		}
		defer file.Close()

		dumper.Dump(file)
	}
}
