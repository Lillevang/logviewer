package filehandler

import (
	"bufio"
	"os"
)

const ChunkSize = 1024

func OpenFile(filePath string) (*bufio.Reader, *os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	reader := bufio.NewReader(file)
	return reader, file, nil
}

func ReadChunk(reader *bufio.Reader) ([]byte, error) {
	chunk := make([]byte, ChunkSize)
	bytesRead, err := reader.Read(chunk)
	if err != nil {
		return nil, err
	}

	return chunk[:bytesRead], nil
}
