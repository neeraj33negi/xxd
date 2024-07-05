package xxd

import (
	"bufio"
	"encoding/hex"
	"os"
)

func FileHex(file *os.File) string {
	if file == nil {
		panic("File expected")
	}
	stats, statsErr := file.Stat()
	if statsErr != nil {
		panic(statsErr)
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	reader := bufio.NewReader(file)
	_, err := reader.Read(bytes) // //read file into a buffer
	if err != nil {
		panic(err)
	} else {
		return hex.Dump(bytes)
	}
}
