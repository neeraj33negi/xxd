package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		panic("Provide the file")
	} else {
		defer f.Close()
		printFileBinary(f)
	}
}

func printFileBinary(file *os.File) {
	if file == nil {
		return
	}
	stats, statsErr := file.Stat()
	if statsErr != nil {
		panic(statsErr)
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err := buffer.Read(bytes)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(hex.Dump(bytes))
	}
}
