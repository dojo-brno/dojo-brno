// +build ignore

// This package is a helper to create a directory structure for a new dojo
// session with today's date. Run it with `go run new.go "problem name"`.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const usage = `go run new.go "problem_name"`

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:", usage)
		os.Exit(1)
	}
	date := time.Now().Format("2006-01-02")
	problemName := strings.Replace(os.Args[1], " ", "_", -1)
	path := fmt.Sprintf("%s/%s", date, problemName)
	if err := os.MkdirAll(path, 0775); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Printf("cd %q\n", path)
}
