package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: task.exe <filename>")
		return
	}

	filePath := args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаем буфер для чтения файла
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	tablesNum := scanner.Text()
	fmt.Println(tablesNum)

	scanner.Scan()
	workingTime := scanner.Text()
	fmt.Println(workingTime)

	scanner.Scan()
	hourPrice := scanner.Text()
	fmt.Println(hourPrice)

	// Читаем файл построчно
	for scanner.Scan() {
		// Выводим содержимое каждой строки файла
		fmt.Println(scanner.Text())
	}

	// Проверяем наличие ошибок при сканировании файла
	if err = scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
	}
}
