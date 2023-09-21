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

func GetMap() map[string]builtinsfn {
	return map[string]builtinsfn{
		"cd": cd,
	}
}
