package pkg

import (
	"encoding/xml"
	"os"
)

func ReadFile(path string) *Testsuite {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	var test Testsuite
	err = xml.Unmarshal(data, &test)

	if err != nil {
		return nil
	}

	return &test
}
