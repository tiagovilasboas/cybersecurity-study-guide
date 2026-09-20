package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Vulnerability struct {
	Name      string `json:"name"`
	Risk      string `json:"risk"`
	Control   string `json:"control"`
	Frequency string `json:"frequency"`
	Effect    string `json:"effect"`
}

type Scenario struct {
	Name            string          `json:"name"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

func main() {
	path := "labs/go/network-hardening/scenario.json"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fail(err)
	}
	var scenario Scenario
	if err := json.Unmarshal(data, &scenario); err != nil {
		fail(err)
	}
	if len(scenario.Vulnerabilities) == 0 {
		fail(fmt.Errorf("scenario has no vulnerabilities"))
	}

	fmt.Println("# Network Hardening Security Risk Assessment")
	fmt.Println()
	fmt.Println("> Scenario-derived draft for review. Complete or adapt it before using it in the course activity.")
	fmt.Printf("\nScenario: %s\n\n## Vulnerabilities and controls\n\n", scenario.Name)
	for i, v := range scenario.Vulnerabilities {
		fmt.Printf("%d. **%s**\n   - Risk: %s\n   - Control: %s\n   - Frequency: %s\n   - Why it helps: %s\n\n", i+1, v.Name, v.Risk, v.Control, v.Frequency, v.Effect)
	}
	fmt.Println("## Selected methods (up to three)\n")
	for i, v := range scenario.Vulnerabilities {
		if i == 3 {
			break
		}
		fmt.Printf("%d. %s\n", i+1, v.Control)
	}
	fmt.Println("\n## Validation evidence\n\n- Review configuration evidence for each selected control.\n- Confirm access logs and change records show the control is active.\n- Repeat the review at the frequency defined for each control.")
}

func fail(err error) { fmt.Fprintf(os.Stderr, "error: %v\n", err); os.Exit(1) }
