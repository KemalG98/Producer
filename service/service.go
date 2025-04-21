package service

import (
	"sync"
)

//go:generate go run github.com/vektra/mockery/v2@v2.53.3 --name=Producer
type Producer interface {
	Produce() ([]string, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.53.3 --name=Presenter
type Presenter interface {
	Present([]string) error
}

type Service struct {
	prod Producer
	pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod: prod, pres: pres}
}

func (s *Service) Mask(data string) string {
	byteData := []byte(data)
	maskedData := make([]byte, 0, len(byteData))
	i := 0
	for i < len(byteData) {
		if i+7 <= len(byteData) && string(byteData[i:i+7]) == "http://" {
			maskedData = append(maskedData, byteData[i:i+7]...)
			i += 7

			for i < len(byteData) && byteData[i] != ' ' && byteData[i] != '\n' && byteData[i] != '\t' {
				i++
			}

			maskLength := i - (len(maskedData) - 7)
			for j := 0; j < maskLength; j++ {
				maskedData = append(maskedData, '*')
			}
		} else {

			maskedData = append(maskedData, maskedData[i])
			i++
		}
	}
	return string(maskedData)
}

// Run Основной метод запуска
func (s *Service) Run() error {
	data, err := s.prod.Produce()
	if err != nil {
		return err
	}

	// Канал для input инфармации
	inputChannel := make(chan string)

	// Канал для результата
	resultChannel := make(chan string)

	var wg sync.WaitGroup
	var mu sync.Mutex

	//Цикл на 10 экземпляров
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for text := range inputChannel {
				masked := s.Mask(text)
				mu.Lock()
				resultChannel <- masked
				mu.Unlock()
			}
		}()
	}
	//Отправляем data в inputChannel
	go func() {
		defer close(inputChannel)
		for _, text := range data {
			inputChannel <- text
		}
	}()
	// Закрываем resultChannel после выполнения
	go func() {
		wg.Wait()
		close(resultChannel)
	}()
	//Возвращаем в Presenter маскираванную data
	return s.pres.Present(data)
}
