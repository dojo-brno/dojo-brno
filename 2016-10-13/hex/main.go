package main

import (
	"io"
	"os"

	"github.com/dojo-brno/dojo-brno/2016-10-13/hex/hex"
)

func main() {
	io.Copy(hex.NewWriter(os.Stdout), os.Stdin)
}
