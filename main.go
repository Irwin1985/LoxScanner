package main

import (
	"LoxScanner/lox"
	"fmt"
	"os"
)

func main() {
	var argCount = len(os.Args)
	if argCount > 3 {
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if argCount == 3 {
		lox.RunFile(os.Args[2])
	} else {
		lox.RunPrompt()
	}
}
