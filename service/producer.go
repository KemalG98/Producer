package service

import (
	"bufio"
	"os"
)

// FileProducer реализует Producer для чтения данных из файла
type FileProducer struct {
	FilePath string
}

// Produce читает данные из файла
func (fp *FileProducer) Produce() ([]string, error) {
	file, err := os.Open(fp.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []string
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
