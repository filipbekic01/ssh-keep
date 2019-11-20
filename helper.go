package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	return text
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
