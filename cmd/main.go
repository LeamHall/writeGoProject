// name:    writeGoProject/cmd/main.go
// version: 0.0.1
// date:    20230114
// author:  Leam Hall

// Notes:
//  - need to take the project module name, and run go mod init.

// Package writeGoProject automates the directory layout of a new Go project.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var dirs = []string{"cmd", "internal", "pkg", "test", "configs", "docs",
	"examples", "api", "web", "build", "scripts", "vendor"}

// createDir creates the directory, mode 0750, or skips if a dir or file
// of that name exists.
func createDir(dir string) {
	err := os.Mkdir(dir, 0750)
	if err != nil {
		if errors.Is(err, fs.ErrExist) {
			fmt.Printf("%s already exists, skipping\n", dir)
		}
	}
}

// isWriteable returns true if the directory is writeable, false otherwise.
func isWriteable(dir string) bool {
	testdir := filepath.Join(dir, "fred12345fred___")
	err := os.Mkdir(testdir, 0750)
	if err != nil {
		fmt.Println("write failed")
		return false
	}
	err = os.Remove(testdir)
	if err != nil {
		fmt.Println("remove failed")
		return false
	}
	return true
}

func main() {
	dir := "."
	if isWriteable(dir) {
		for _, d := range dirs {
			newDir := filepath.Join(dir, d)
			createDir(newDir)
		}
	} else {
		fmt.Printf("cannot create a directory in %s\n", dir)
		os.Exit(1)
	}
}
