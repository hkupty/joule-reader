package pkg

import (
	"strconv"
	"strings"
)

type Failure struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
}

type Testcase struct {
	Id        string  `xml:"id,attr"`
	Name      string  `xml:"name,attr"`
	Classname string  `xml:"classname,attr"`
	Time      float64 `xml:"time,attr"`
	Failure   Failure `xml:"failure"`
}

type Testsuite struct {
	Id       string     `xml:"id,attr"`
	Name     string     `xml:"name,attr"`
	Tests    int        `xml:"tests,attr"`
	Skipped  int        `xml:"skipped,attr"`
	Failures int        `xml:"failures,attr"`
	Errors   int        `xml:"errors,attr"`
	Time     float64    `xml:"time,attr"`
	Testcase []Testcase `xml:"testcase"`
}

type Testsuites struct {
	Id        string      `xml:"id,attr"`
	Name      string      `xml:"name,attr"`
	Tests     int         `xml:"tests,attr"`
	Failures  int         `xml:"failures,attr"`
	Time      float64     `xml:"time,attr"`
	Testsuite []Testsuite `xml:"testsuite"`
}

func (t Testcase) String() string {
	var sb strings.Builder
	sb.WriteString("Test -> ")
	sb.WriteString(t.Name)
	sb.WriteString(" (")
	sb.WriteString(strconv.FormatFloat(t.Time, 'f', -1, 64))
	sb.WriteString("s)\n")

	return sb.String()
}

func (t Testsuite) String() string {
	var sb strings.Builder
	sb.WriteString("Testsuite { ")
	sb.WriteString(t.Name)
	sb.WriteString(" } ")
	sb.WriteString("\n\tTests: ")
	sb.WriteString(strconv.Itoa(t.Tests))
	sb.WriteString("\n\tFailures: ")
	sb.WriteString(strconv.Itoa(t.Failures))
	sb.WriteString("\n\tErrors: ")
	sb.WriteString(strconv.Itoa(t.Errors))
	sb.WriteString("\n\tSkipped: ")
	sb.WriteString(strconv.Itoa(t.Skipped))
	sb.WriteString("\n\n --- \n\n")

	for _, test := range t.Testcase {
		sb.WriteString(test.String())
	}

	return sb.String()
}
