package parser

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name          string
		inputFile     string
		outputFile    string
		expectedError error
	}{
		{
			name:          "simple test",
			inputFile:     "../../test/case1/simple_test.txt",
			outputFile:    "../../test/case1/simple_test_out.txt",
			expectedError: nil,
		},
		{
			name:          "client already in club test",
			inputFile:     "../../test/case2/client_already_in_club.txt",
			outputFile:    "../../test/case2/client_already_in_club_out.txt",
			expectedError: nil,
		},
		{
			name:          "table switch test",
			inputFile:     "../../test/case3/table_switch.txt",
			outputFile:    "../../test/case3/table_switch_out.txt",
			expectedError: nil,
		},
		{
			name:          "queue is full test",
			inputFile:     "../../test/case4/queue_full.txt",
			outputFile:    "../../test/case4/queue_full_out.txt",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// файл для вывода результата работы
			output, err := os.Create("../../test/output.txt")
			if err != nil {
				t.Fatalf("failed to create output file: %v", err)
			}
			defer output.Close()

			oldStdout := os.Stdout
			defer func() { os.Stdout = oldStdout }()
			os.Stdout = output

			err = Parse(tc.inputFile)
			if err != tc.expectedError {
				t.Errorf("Parse() вернул не ожиданную ошибку: %v", err)
			}

			file1, err := os.Open(tc.outputFile)
			if err != nil {
				t.Errorf("Ошибка при открытии тестового файла: %v", err)
				return
			}
			defer file1.Close()

			file2, err := os.Open("../../test/output.txt")
			if err != nil {
				t.Errorf("Ошибка при открытии результирующего файла: %v", err)
				return
			}
			defer file2.Close()

			scanner := bufio.NewScanner(file1)
			scanner2 := bufio.NewScanner(file2)
			for scanner.Scan() {
				scanner2.Scan()
				if !reflect.DeepEqual(scanner.Text(), scanner2.Text()) {
					t.Fatalf("Вывод не совпадает у кейса %s получено %v ожидалось %v", tc.name, scanner2.Text(), scanner.Text())
				}
			}
			if err := scanner.Err(); err != nil {
				t.Errorf("Ошибка при чтении первого файла: %v", err)
			}
			if err := scanner2.Err(); err != nil {
				t.Errorf("Ошибка при чтении второго файла: %v", err)
			}
		})
	}
}

//go test ./... -coverprofile=coverage.out -coverpkg=./...
//go tool cover -html .\coverage -o cover.html
