package builtins

import (
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

func GetMap() map[string]builtinsfn {
	return map[string]builtinsfn{
		"cd":   cd,
		"exit": exit,
	}
}
