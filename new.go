// +build ignore

// This is a helper "script" to create a directory structure for a new dojo
// session with today's date. Run it as:
//
//   go run new.go PACKAGE_NAME
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const usage = `go run new.go PACKAGE_NAME`

func createDatedPackage(name string) (path string, err error) {
	date := time.Now().Format("2006-01-02")
	path = filepath.Join(date, name)
	return path, os.MkdirAll(path, 0775)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:", usage)
		os.Exit(1)
	}
	name := strings.Replace(os.Args[1], " ", "_", -1)
	path, err := createDatedPackage(name)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Printf("cd %q\n", path)
}
