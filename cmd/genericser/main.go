package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jochenczemmel/genericser/convert"
)

func main() {

	inFile := flag.String("in", "", "go source file")
	outFile := flag.String("out", "", "go target file")
	replacements := flag.String("r", "",
		`replacements, format: "from1=to1[,from2=to2[,...]]"`)
	flag.Parse()

	if *inFile == "" || *outFile == "" || *replacements == "" {
		fmt.Fprintf(os.Stderr,
			"usage: %s -in in.go -out out.go -r x1=y1,x2=y2", os.Args[0])
		os.Exit(1)
	}

	rep := convert.Replacer{}
	definitions := strings.Split(*replacements, ",")
	for _, def := range definitions {
		defElem := strings.Split(def, "=")
		if len(defElem) != 2 {
			fmt.Fprintf(os.Stderr, "ERROR: invalid replacement: %s", defElem)
			os.Exit(2)
		}
		rep[defElem[0]] = defElem[1]
	}

	err := convert.ConvertFile(*inFile, *outFile, rep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(2)
	}

}
