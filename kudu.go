package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"fmt"

	"os"

	"github.com/rokob/kudu/parser"
	"github.com/rokob/kudu/repl"
)

func main() {
	outputFilename := flag.String("o", "a.out", "Output filename")
	flag.Parse()

	if len(flag.Args()) > 1 {
		flag.Usage()
		return
	}
	if len(flag.Args()) == 1 {
		compile(flag.Args()[0], *outputFilename)
		return
	}
	repl.Run()
}

func compile(infile string, outfile string) {
	input, err := ioutil.ReadFile(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s\n", infile)
		return
	}
	p := parser.New(parser.CompilerMode)
	ok, _, es := p.Parse(string(input))
	if ok {
		var b []byte
		var err error
		if len(es) == 1 {
			b, err = json.MarshalIndent(es[0], "", "  ")
		} else {
			b, err = json.MarshalIndent(es, "", "  ")
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return
		}
		err = ioutil.WriteFile(outfile, b, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		return
	}
	fmt.Fprintf(os.Stderr, "Problem parsing input\n")
}
