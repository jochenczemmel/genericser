package convert

import "testing"

func TestConvertFile(t *testing.T) {
	// outFile := "generated.go.txt"
	rep := Replacer{
		"T1": "string",
		"T2": "StringList",
	}
	inFile := "../perllist/perllist.go"
	outFile := "perllist_string.go"
	ConvertFile(inFile, outFile, rep)
	/*
		inFile = "../perllist/perllist_test.go"
		outFile = "perllist_string_test.go"
		ConvertFile(inFile, outFile, rep)
	*/
}
