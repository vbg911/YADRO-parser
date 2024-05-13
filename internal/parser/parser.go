package parser

import (
	"YADRO-parser/internal/cyberclub"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const timeLayout = "15:04"

func Parse(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Создаем буфер для чтения файла
	scanner := bufio.NewScanner(file)

	// Получаем количество столов
	scanner.Scan()
	tablesNum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}

	// Получаем рабочие часы клуба
	scanner.Scan()
	workingTime := strings.Split(scanner.Text(), " ")
	workingStart, err := time.Parse(timeLayout, workingTime[0])
	if err != nil {
		return err
	}

	workingEnd, err := time.Parse(timeLayout, workingTime[1])
	if err != nil {
		return err
	}

	// Получаем цену 1 часа
	scanner.Scan()
	hourPrice, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}

	club := cyberclub.Cyberсlub{
		TableAmount:  tablesNum,
		WorkingStart: workingStart,
		WorkingEnd:   workingEnd,
		HourPrice:    hourPrice,
		Scanner:      scanner,
	}

	err = club.RunParsing()

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
