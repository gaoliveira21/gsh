package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func runCmd(args []string) {
	name := args[0]

	cmd := exec.Command(name, args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetPrefix("gsh: ")
	log.SetFlags(0)

	for {
		cwd := getCwd()

		fmt.Printf("%s> ", cwd)

		line := readLine()
		args := splitLine(line)

		runCmd(args)
	}
}
