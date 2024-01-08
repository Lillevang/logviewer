package main

import (
	"io"
	"log"

	"github.com/lillevang/logviewer/filehandler"
	"github.com/lillevang/logviewer/logviewer"
)

func main() {
	reader, file, err := filehandler.OpenFile("log.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	for {
		_, err := filehandler.ReadChunk(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading chunk: %v", err)
		}

	}

	logviewer.Run()

}
