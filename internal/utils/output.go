package utils

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func CreateFile(output string) (*os.File, error) {
	fileOuput, err := os.Create(output)
	if err != nil {
		log.Debug("create file: %v", err)
		return nil, err
	}
	return fileOuput, nil
}

func WriteToOutput(fileOuput *os.File, resp io.Reader) (int64, error) {
	n, err := io.Copy(fileOuput, resp)
	defer fileOuput.Close()
	return n, err
}
