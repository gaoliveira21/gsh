package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getCwd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	cwd := filepath.Dir(ex)

	return cwd
}

func main() {
	cwd := getCwd()

	fmt.Printf("%s> ", cwd)

}
