package service

import (
	"os"
)

// FilePresenter использует Presenter для записи данных в файл
type FilePresenter struct {
	FilePath string
}

// Present записывает данные в файл
func (fp *FilePresenter) Present(data []string) error {
	file, err := os.Create(fp.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, d := range data {
		_, err := file.WriteString(d + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
