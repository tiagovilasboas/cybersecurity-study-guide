package studyreport

import (
	"encoding/json"
	"fmt"
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

type Report struct {
	Title    string    `json:"title"`
	Context  string    `json:"context"`
	Findings []Finding `json:"findings"`
}

func Parse(data []byte) (Report, error) {
	var r Report
	if err := json.Unmarshal(data, &r); err != nil {
		return r, err
	}
	return r, Validate(r)
}

func Validate(r Report) error {
	if strings.TrimSpace(r.Title) == "" {
		return fmt.Errorf("title is required")
	}
	if len(r.Findings) == 0 {
		return fmt.Errorf("at least one finding is required")
	}
	for i, f := range r.Findings {
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

func (r Report) Markdown() string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n", r.Title)
	if r.Context != "" {
		fmt.Fprintf(&b, "> %s\n\n", r.Context)
	}
	b.WriteString("## Findings and conditional recommendations\n\n")
	for i, f := range r.Findings {
		fmt.Fprintf(&b, "%d. **%s** (%s)\n", i+1, f.Title, f.Severity)
		b.WriteString("   - Evidence:\n")
		for _, e := range f.Evidence {
			fmt.Fprintf(&b, "     - %s\n", e)
		}
		fmt.Fprintf(&b, "   - Risk: %s\n   - Recommendation: %s\n   - Cadence: %s\n", f.Risk, f.Recommendation, f.Cadence)
		b.WriteString("   - Validation:\n")
		for _, v := range f.Validation {
			fmt.Fprintf(&b, "     - %s\n", v)
		}
		b.WriteString("\n")
	}
	return b.String()
}
