package lox

import (
	"LoxScanner/scanner"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var HadError bool = false

func RunFile(path string) {
	file, err := filepath.Abs(path)
	check(err)

	fileInfo, err := os.Stat(file)
	check(err)

	// Leemos todos los bytes del fichero
	bytes, err := os.ReadFile(fileInfo.Name())
	check(err)
	// Ejecutamos el código fuente
	run(string(bytes))
	if HadError {
		os.Exit(65)
	}
}

func RunPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to <<Lox>> programming language.")
	fmt.Println("Please type some commands and press INTRO.")
	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			continue
		}
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		run(input)
		HadError = false
	}
}

func run(source string) {
	lexer := scanner.New(source)
	tokens := lexer.ScanTokens()

	// imprime los tokens
	for _, tok := range tokens {
		fmt.Println(tok.ToString())
	}
}

func Error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	msg := fmt.Errorf("[line %d] Error %s: %s", line, where, message)
	fmt.Println(msg)
	HadError = true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
