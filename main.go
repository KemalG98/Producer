package main

import (
	"fmt"
	"log"
	"os"
	"src/service"
)

func main() {
	// Проверка наличия аргументов
	if len(os.Args) < 2 {
		log.Fatal("Применяем: go run main.go inputPath outputPath")
	}
	inputPath := os.Args[1]
	outputPath := "output.txt" // значение по умолчанию, если не указан выходной файл

	// Проверка обязательного аргумента
	if inputPath == "" {
		log.Fatal("Input должен быть уточнен.")
	}
	// Выводим, если указан выходной файл
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}
	fmt.Printf("Input file: %s\n", inputPath)
	fmt.Printf("Output file: %s\n", outputPath)

	// Создаем экземпляры FileProducer и FilePresenter с заданными путями
	var producer = &service.FileProducer{FilePath: inputPath}
	var presenter = &service.FilePresenter{FilePath: outputPath}

	// Создаем сервис
	var svc = service.NewService(producer, presenter) // Запускаем сервис
	(*svc).Run()
}

// Для запуска go run main.go input.txt output.txt
