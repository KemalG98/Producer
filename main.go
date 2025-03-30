package main

import (
	"fmt"
	"log"
	"os"
	"src/service"
)

func main() {
	// Проверка наличия аргументов
	if len(os.Args) < 3 {
		log.Fatal("Применяем: go run main.go inputPath outputPath")
	}
	var inputPath string
	var outputPath string
	// Уточнен ли выходной файл
	var outputSpecified bool

	// Обработка аргументов Input, Output
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "input":
			if i+1 < len(os.Args) {
				inputPath = os.Args[i+1]
				i++
			} else {
				log.Fatal("Input должен быть уточнен")
			}
		case "output":
			if i+1 < len(os.Args) {
				outputPath = os.Args[i+1]
				outputSpecified = true
				i++
			}
		}
	}

	// Проверка обязательного аргумента
	if inputPath == "" {
		log.Fatal("Input должен быть уточнен.")
	}
	// Выводим, если выходной файл не указан
	if !outputSpecified {
		outputPath = "output.txt"
		log.Println("Не уточнен выхоной файл,используем стандартный: output.txt")
	} else {
		log.Println("Используем указаный выходной файл: %s\n", outputPath)
	}
	fmt.Printf("Input file: %s\n", inputPath)
	fmt.Printf("Output file: %s\n", outputPath)

	// Создаем экземпляры FileProducer и FilePresenter с заданными путями
	producer := &service.FileProducer{FilePath: inputPath}
	presenter := &service.FilePresenter{FilePath: outputPath}

	// Создаем сервис
	var svc = service.NewService(producer, presenter) // Запускаем сервис
	(*svc).Run()
}

// Для запуска go run main.go -input input.txt -output output.txt
