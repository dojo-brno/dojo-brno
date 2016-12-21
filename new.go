// +build ignore

// This is a helper "script" to create a directory structure for a new dojo
// session with today's date. Run it as:
//
//   go run new.go PACKAGE_NAME [EXERCISE_URL]
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const usage = `go run new.go PACKAGE_NAME [EXERCISE_URL]`

func createDatedPackage(name string) (path string, err error) {
	date := time.Now().Format("2006/2006-01-02")
	path = filepath.Join(date, name)
	if _, err := os.Stat(path); err == nil {
		return path, fmt.Errorf("path already exists, aborting: %s", path)
	}
	return path, os.MkdirAll(path, 0775)
}

func createCommonFiles(path string, url string) error {
	f, err := os.Create(filepath.Join(path, "AUTHORS"))
	if err != nil {
		return err
	}
	f.Close()
	f, err = os.Create(filepath.Join(path, "README.md"))
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "# Exercise\n\n%s\n", url)
	return nil
}

func main() {
	if l := len(os.Args); l < 2 || l > 3 {
		fmt.Println("usage:", usage)
		os.Exit(1)
	}
	name := strings.Replace(os.Args[1], " ", "_", -1)
	var url string
	if len(os.Args) > 2 {
		url = os.Args[2]
	}
	path, err := createDatedPackage(name)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	if err := createCommonFiles(path, url); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Printf("cd %q\n", path)
}
