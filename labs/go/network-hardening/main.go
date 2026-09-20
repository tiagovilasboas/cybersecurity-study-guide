package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Finding struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Severity       string   `json:"severity"`
	Evidence       []string `json:"evidence"`
	Risk           string   `json:"risk"`
	Recommendation string   `json:"recommendation"`
	Cadence        string   `json:"cadence"`
	Validation     []string `json:"validation"`
}

type ReportInput struct {
	Title    string    `json:"title"`
	Context  string    `json:"context"`
	Findings []Finding `json:"findings"`
}

func main() {
	inputPath := flag.String("input", "labs/go/network-hardening/scenario.json", "path to report input JSON")
	format := flag.String("format", "markdown", "output format: markdown or json")
	flag.Parse()

	data, err := os.ReadFile(*inputPath)
	if err != nil {
		fail(err)
	}
	var input ReportInput
	if err := json.Unmarshal(data, &input); err != nil {
		fail(err)
	}
	if err := validate(input); err != nil {
		fail(err)
	}

	switch strings.ToLower(*format) {
	case "markdown", "md":
		writeMarkdown(input)
	case "json":
		out, err := json.MarshalIndent(input, "", "  ")
		if err != nil {
			fail(err)
		}
		fmt.Println(string(out))
	default:
		fail(fmt.Errorf("unsupported format %q", *format))
	}
}

func validate(input ReportInput) error {
	if strings.TrimSpace(input.Title) == "" {
		return fmt.Errorf("title is required")
	}
	if len(input.Findings) == 0 {
		return fmt.Errorf("at least one finding is required")
	}
	for i, f := range input.Findings {
		if f.ID == "" || f.Title == "" {
			return fmt.Errorf("finding %d requires id and title", i+1)
		}
		if len(f.Evidence) == 0 {
			return fmt.Errorf("finding %q requires evidence", f.ID)
		}
		if f.Recommendation == "" || f.Cadence == "" {
			return fmt.Errorf("finding %q requires recommendation and cadence", f.ID)
		}
	}
	return nil
}

func writeMarkdown(input ReportInput) {
	fmt.Printf("# %s\n\n", input.Title)
	if input.Context != "" {
		fmt.Printf("> %s\n\n", input.Context)
	}
	fmt.Println("## Findings and conditional recommendations\n")
	for i, f := range input.Findings {
		fmt.Printf("%d. **%s** (%s)\n", i+1, f.Title, f.Severity)
		fmt.Println("   - Evidence:")
		for _, e := range f.Evidence {
			fmt.Printf("     - %s\n", e)
		}
		fmt.Printf("   - Risk: %s\n", f.Risk)
		fmt.Printf("   - Recommendation: %s\n", f.Recommendation)
		fmt.Printf("   - Cadence: %s\n", f.Cadence)
		fmt.Println("   - Validation:")
		for _, v := range f.Validation {
			fmt.Printf("     - %s\n", v)
		}
		fmt.Println()
	}
}

func fail(err error) { fmt.Fprintf(os.Stderr, "error: %v\n", err); os.Exit(1) }
