package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gaoliveira21/gsh/builtins"
)

func getCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

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

	builtinMap := builtins.GetMap()
	builtinFn, found := builtinMap[name]

	if found {
		err := builtinFn(args...)

		if err != nil {
			log.Println(err)
		}
	} else {
		cmd := exec.Command(name, args[1:]...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Println(err)
		}
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
