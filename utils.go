package main

import (
	"os"
	"path/filepath"
)

func getPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
