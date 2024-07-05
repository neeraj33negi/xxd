package main

import (
	"flag"
	"fmt"
	"os"
	"xxd"
)

var (
	littleEndianPtr *bool
	filenamePtr     *string
	littleEndian    bool
	filename        string
)

func initParseCommandLineFlags() {
	littleEndianPtr = flag.Bool("e", false, "Combined hex")
	filenamePtr = flag.String("file", "", "Target file")
	flag.Parse()
	littleEndian = *littleEndianPtr
	filename = *filenamePtr
}

func main() {

	initParseCommandLineFlags()

	f, err := os.Open(filename)
	if err != nil {
		panic("Provide the file")
	} else {
		defer f.Close()
		hexCode := xxd.FileHex(f)
		fmt.Print(hexCode)
	}
}
