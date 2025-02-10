package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/hkupty/joule-reader/pkg"
)

var exitCodeOnFailureOrError = flag.Bool("x", false, "Fail executable if any of the passed files contain errors or failures")
var verbose = flag.Bool("v", false, "Print all test suites and test cases independently of them containing errors or not")

var shortCircuitOut = flag.Bool("s", false, "Stop executing on the first failure")
var openEditor = flag.Bool("o", false, "Open the editor on the location of the error")

var ignoreErrors = flag.Bool("e", false, "Will ignore errors, only count failures")
var ignoreFailures = flag.Bool("f", false, "Will ignore failures, only count errors")

func main() {
	ret := 0
	flag.Parse()
	for _, arg := range flag.Args() {
		f := pkg.ReadFile(arg)
		hasFailures := f.Failures > 0 && !(*ignoreFailures)
		hasErrors := f.Errors > 0 && !(*ignoreErrors)

		failureOrError := hasFailures || hasErrors

		if *exitCodeOnFailureOrError && failureOrError {
			ret += 1
		}

		if *verbose || failureOrError {
			fmt.Println(f)
		}

		if *shortCircuitOut && failureOrError {
			break
		}

		if *openEditor && failureOrError {
			var editor string
			if editor = os.ExpandEnv("$EDITOR"); editor == "" {
				slog.Warn("Editor is not set, aborting")
				os.Exit(ret + 1)
			}

			slog.Info("Opening editor")

			editorCmd := strings.Split(editor, " ")
			editorCmd = append(editorCmd, arg)
			editorPath, err := exec.LookPath(editorCmd[0])
			if err != nil {
				slog.Warn("Failed to locate $EDITOR", "error", err)
			}

			editorCmd[0] = editorPath

			err = syscall.Exec(editorCmd[0], editorCmd, syscall.Environ())
			if err != nil {
				slog.Warn("Failed to open editor", "editor-command", editorCmd, "error", err)
			}
		}

	}

	os.Exit(ret)
}
