package main

import (
	"fmt"
	"os"

	"github.com/hkupty/joule-reader/pkg"
)

func main() {
	for _, arg := range os.Args[1:] {
		f := pkg.ReadFile(arg)
		fmt.Println(f)
	}
}
