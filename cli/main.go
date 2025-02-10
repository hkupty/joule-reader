package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hkupty/joule-reader/pkg"
)

var failOnError = flag.Bool("x", false, "Fail executable if any of the passed files contain errors or failures")
var verbose = flag.Bool("v", false, "Print all test suites and test cases independently of them containing errors or not")

func main() {
	ret := 0
	flag.Parse()
	for _, arg := range flag.Args() {
		f := pkg.ReadFile(arg)
		hasFailures := f.Failures > 0
		hasErrors := f.Errors > 0

		switch {
		case *failOnError && (hasFailures || hasErrors):
			ret += 1
		}

		if *verbose || (hasFailures || hasErrors) {
			fmt.Println(f)
		}
	}

	os.Exit(ret)
}
