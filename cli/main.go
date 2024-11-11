package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/hkupty/joule-reader/pkg"
)

var args struct {
	FailuresOrErrors bool     `arg:"-x"`
	Files            []string `arg:"positional"`
}

func main() {
	ret := 0
	arg.MustParse(&args)
	for _, arg := range args.Files {
		f := pkg.ReadFile(arg)
		hasFailures := f.Failures > 0
		hasErrors := f.Errors > 0

		switch {
		case args.FailuresOrErrors && (hasFailures || hasErrors):
			ret += 1
			fmt.Println(f)
		case !args.FailuresOrErrors:
			fmt.Println(f)
		}

	}

	os.Exit(ret)
}
