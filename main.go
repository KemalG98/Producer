package main

import (
	"flag"
	"log"
	"src/service"
)

func main() {
	// Определяем флаг для пути к входному файлу
	inputPath := flag.String("input", "", "Path to input file (required)")
	flag.Parse()

	// Проверка обязательного аргумента
	if *inputPath == "" {
		log.Fatal("Input file path must be specified.")
	}
	// Выводим, если выходной файл не указан
	outputPath := flag.String("output", "", "Path to output file (required)")
	if *outputPath == "" {
		log.Println("No output path specified, using default: output.txt")
	} else {
		log.Printf("Using output file: %s\n", *outputPath)
	}
	// Определяем флаги командной строки для ввода и вывода
	inputPath = flag.String("input", "input.txt", "Path to input file")
	outputPath = flag.String("output", "output.txt", "Path to output file")

	// Парсим флаги командной строки(анализ)
	flag.Parse()

	// Создаем экземпляры FileProducer и FilePresenter с заданными путями
	producer := &service.FileProducer{FilePath: *inputPath}
	presenter := &service.FilePresenter{FilePath: *outputPath}

	// Создаем сервис
	var svc = service.NewService(producer, presenter) // Запускаем сервис
	(*svc).Run()
}

// Для запуска go run main.go -input input.txt -output output.txt
