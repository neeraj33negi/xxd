package xxd

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileHex(file *os.File) {
	if file == nil {
		panic("File expected")
	}

	bytes := make([]byte, 256)

	reader := bufio.NewReader(file)
	for {
		_, err := reader.Read(bytes) // //read file into a buffer
		if err != nil {
			if !errors.Is(err, io.EOF) {
				panic(err)
			}
			break
		}
		printHex(bytes)
	}
}

func printHex(data []byte) {
	if len(data) == 0 {
		return
	}
	fmt.Println(hex.Dump(data))
}
