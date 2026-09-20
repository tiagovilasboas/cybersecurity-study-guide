package main

import "fmt"

type Vulnerability struct {
	Name      string
	Risk      string
	Control   string
	Frequency string
	Effect    string
}

func main() {
	vulnerabilities := []Vulnerability{
		{
			Name:      "Shared employee passwords",
			Risk:      "Accountability is lost and one leaked password can expose multiple accounts.",
			Control:   "Unique accounts with strong password policy and MFA.",
			Frequency: "Enforce continuously; review access monthly.",
			Effect:    "Reduces credential reuse and improves traceability of network access.",
		},
		{
			Name:      "Default database administrator password",
			Risk:      "A known credential can provide direct privileged access to sensitive data.",
			Control:   "Replace the default secret with a unique managed credential and MFA.",
			Frequency: "Change immediately; rotate and review regularly.",
			Effect:    "Removes an easily guessed entry point and limits privileged access.",
		},
		{
			Name:      "Firewalls without ingress and egress rules",
			Risk:      "Unnecessary or malicious traffic can enter or leave the network.",
			Control:   "Implement deny-by-default firewall rules with documented exceptions.",
			Frequency: "Review monthly and after every network change.",
			Effect:    "Reduces the exposed attack surface and limits lateral movement and exfiltration.",
		},
		{
			Name:      "No multifactor authentication",
			Risk:      "A stolen password may be enough to access the network.",
			Control:   "Require MFA for administrative and remote network access.",
			Frequency: "Enforce continuously; audit adoption monthly.",
			Effect:    "Adds an independent verification factor beyond the password.",
		},
	}

	fmt.Println("# Network Hardening Security Risk Assessment")
	fmt.Println()
	fmt.Println("> Scenario-derived draft for review. Complete or adapt it before using it in the course activity.")
	fmt.Println()
	fmt.Println("## Vulnerabilities and controls")
	fmt.Println()
	for i, v := range vulnerabilities {
		fmt.Printf("%d. **%s**\n", i+1, v.Name)
		fmt.Printf("   - Risk: %s\n", v.Risk)
		fmt.Printf("   - Control: %s\n", v.Control)
		fmt.Printf("   - Frequency: %s\n", v.Frequency)
		fmt.Printf("   - Why it helps: %s\n\n", v.Effect)
	}

	fmt.Println("## Recommended priority")
	fmt.Println()
	fmt.Println("1. Implement deny-by-default firewall rules for inbound and outbound traffic.")
	fmt.Println("2. Remove the default database administrator password and require MFA for privileged access.")
	fmt.Println("3. Eliminate shared passwords by assigning unique accounts and enforcing password policy.")
	fmt.Println()
	fmt.Println("## Validation evidence")
	fmt.Println()
	fmt.Println("- Exported firewall rules show only documented traffic is allowed.")
	fmt.Println("- A credential review confirms that default and shared passwords are no longer active.")
	fmt.Println("- Access logs show unique identities and MFA events for privileged sessions.")
}
