package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCwd() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	cwd := filepath.Dir(ex)

	return cwd
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return line
}

func splitLine(line string) []string {
	args := strings.Fields(line)

	return args
}

func main() {
	log.SetPrefix("gsh: ")
	log.SetFlags(0)

	cwd := getCwd()

	fmt.Printf("%s> ", cwd)

	line := readLine()
	args := splitLine(line)

	fmt.Println(args)
}
