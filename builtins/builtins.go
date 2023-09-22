package builtins

import (
	"fmt"
	"os"
	"strings"
)

type builtinsfn func(...string) error

func cd(args ...string) error {
	path := strings.Join(args[1:], "")

	err := os.Chdir(path)

	return err
}

func exit(args ...string) error {
	os.Exit(0)

	return nil
}

func help(args ...string) error {
	fmt.Println("*** G Shell ***")
	fmt.Println("Type program names and arguments and press enter.")
	fmt.Println("The following commands are built in:")

	builtinsMap := GetMap()

	for key := range builtinsMap {
		fmt.Printf(" %s\n", key)
	}

	fmt.Println("Use the man command for information on other programs.")

	return nil
}

func GetMap() map[string]builtinsfn {
	return map[string]builtinsfn{
		"cd":   cd,
		"exit": exit,
		"help": help,
	}
}
