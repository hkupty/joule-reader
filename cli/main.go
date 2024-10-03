package main

import (
	"fmt"
	"os"

	"github.com/hkupty/joule-reader/pkg"
)

func main() {
	args := os.Args
	f := pkg.ReadFile(args[1])
	fmt.Println(f)
}
