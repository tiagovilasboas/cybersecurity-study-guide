// Exercício autoral: detectar possíveis ataques de força bruta em logs.
// Não reproduz o laboratório da Coursera; transforma o conceito em uma prática própria.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const threshold = 3

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "uso: go run . caminho-do-log\\n")
		os.Exit(2)
	}

	counts, err := failedAttempts(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for ip, count := range counts {
		if count >= threshold {
			fmt.Printf("alerta: %s teve %d tentativas de login com falha\\n", ip, count)
		}
	}
}

func failedAttempts(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("abrir log: %w", err)
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 3 || fields[2] != "FAIL" {
			continue
		}
		counts[fields[1]]++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ler log: %w", err)
	}
	return counts, nil
}
