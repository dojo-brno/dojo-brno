package main

import (
	"fmt"
	"io"
	"os"

	"github.com/dojo-brno/dojo-brno/2016/2016-10-13/hex/hex"
)

func main() {
	n, err := io.Copy(hex.NewWriter(os.Stdout), os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if n > 0 {
		fmt.Fprintln(os.Stdout)
	}
}
