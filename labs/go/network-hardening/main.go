package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tiagovilasboas/cybersecurity-study-guide/pkg/studyreport"
)

func main() {
	inputPath := flag.String("input", "labs/go/network-hardening/scenario.json", "path to report input JSON")
	format := flag.String("format", "markdown", "output format: markdown or json")
	flag.Parse()
	data, err := os.ReadFile(*inputPath)
	if err != nil {
		fail(err)
	}
	report, err := studyreport.Parse(data)
	if err != nil {
		fail(err)
	}
	switch strings.ToLower(*format) {
	case "markdown", "md":
		fmt.Print(report.Markdown())
	case "json":
		fmt.Println(string(mustJSON(report)))
	default:
		fail(fmt.Errorf("unsupported format %q", *format))
	}
}

func mustJSON(v any) []byte {
	b, err := jsonMarshal(v)
	if err != nil {
		fail(err)
	}
	return b
}
func jsonMarshal(v any) ([]byte, error) { return json.MarshalIndent(v, "", "  ") }
func fail(err error)                    { fmt.Fprintf(os.Stderr, "error: %v\n", err); os.Exit(1) }
