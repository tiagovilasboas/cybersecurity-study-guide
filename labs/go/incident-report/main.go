package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	logPath := "labs/go/bruteforce-detector/sample.log"
	if len(os.Args) > 1 {
		logPath = os.Args[1]
	}

	alerts, err := failedAttempts(logPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro ao ler log: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("# Rascunho — Security Incident Report")
	fmt.Println()
	fmt.Println("> Revise e adapte o texto antes de usar na atividade da Coursera.")
	fmt.Println()
	fmt.Println("## 1. Protocolo observado")
	fmt.Println()
	fmt.Println("O tráfego mostrou consultas DNS para resolver os domínios **yummyrecipesforme.com** e **greatrecipesforme.com**. Também foram observadas requisições HTTP para carregar a página, baixar o arquivo malicioso e redirecionar os usuários.")
	fmt.Println()
	fmt.Println("## 2. Descrição do incidente")
	fmt.Println()
	fmt.Println("O site **yummyrecipesforme.com** foi comprometido depois que o atacante descobriu a senha padrão da conta administrativa usando várias tentativas de login.")
	fmt.Println()
	fmt.Println("Após obter acesso, o atacante alterou o código JavaScript do site para induzir os visitantes a baixar um arquivo executável. O arquivo redirecionava os usuários para **greatrecipesforme.com**.")
	fmt.Println()
	fmt.Println("O impacto foi a exposição dos visitantes a malware e o redirecionamento para um site controlado pelo atacante.")
	fmt.Println()
	fmt.Println("## 3. Recomendação")
	fmt.Println()
	fmt.Println("Remover senhas padrão, exigir senhas fortes e habilitar autenticação multifator para a conta administrativa. Também é recomendável limitar tentativas de login e bloquear temporariamente endereços IP com comportamento de força bruta.")
	fmt.Println()
	fmt.Println("A eficácia pode ser validada monitorando novas tentativas falhas de login e confirmando que acessos repetidos são bloqueados ou geram alertas.")
	fmt.Println()
	fmt.Println("## Evidência complementar do laboratório em Go")
	fmt.Println()
	fmt.Printf("Log analisado: `%s`\n\n", logPath)
	if len(alerts) == 0 {
		fmt.Println("Nenhum IP atingiu o limiar de tentativas falhas configurado no detector.")
	} else {
		for ip, count := range alerts {
			fmt.Printf("Resultado observado: `alerta: %s teve %d tentativas de login com falha`\n", ip, count)
		}
	}
	fmt.Println()
	fmt.Println("A evidência acima é calculada a partir do log informado e complementa o relatório; ela não substitui a validação da atividade na Coursera.")
}

func failedAttempts(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 3 && fields[2] == "FAIL" {
			counts[fields[1]]++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	alerts := make(map[string]int)
	for ip, count := range counts {
		if count >= 3 {
			alerts[ip] = count
		}
	}
	return alerts, nil
}
