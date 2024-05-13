package main

import (
	"YADRO-parser/internal/parser"
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
	err := parser.Parse(filePath)
	if err != nil {
		fmt.Println(err)
	}
}
