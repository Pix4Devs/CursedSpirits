package filesystem

import (
	"io"
	"log"
	"os"
	"strings"
)

func Read(filepath string) []string {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	var data string

	for {
		chunk := make([]byte, 1042)
		rL, err := file.Read(chunk); if err == io.EOF {
			break
		}

		data += string(chunk[:rL])
	}

	return strings.Split(strings.ReplaceAll(strings.TrimSpace(data),"\r",""), "\n")
}