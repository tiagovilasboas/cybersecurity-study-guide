// Gera um rascunho de relatório em Markdown para estudo.
// Os campos entre colchetes devem ser completados pelo aluno.
package main

import "fmt"

func main() {
	fmt.Println(`# Rascunho — Security Incident Report

> Complete os campos com suas próprias palavras antes de usar este material na atividade.

## 1. Protocolo observado

- Protocolo identificado: [preencha após revisar o log tcpdump]
- Papel no fluxo: [explique o que o protocolo fez]
- Evidência: [cite a linha ou sequência observada]

## 2. Descrição do incidente

- Ativo afetado: [preencha]
- Como o acesso foi obtido: [preencha]
- Alterações observadas: [preencha]
- Impacto para os usuários: [preencha]
- Evidências utilizadas: [tcpdump, sandbox, relatos e análise do código]

## 3. Recomendação

- Controle escolhido: [preencha]
- Como reduz o risco: [preencha]
- Como validar a eficácia: [preencha]

## Evidência complementar do laboratório em Go

Comando executado:

` + "`go run ./labs/go/bruteforce-detector ./labs/go/bruteforce-detector/sample.log`" + `

Resultado observado:

` + "`alerta: 203.0.113.10 teve 3 tentativas de login com falha`" + `

Esse resultado demonstra o comportamento do protótipo de detecção e não substitui a validação da atividade na Coursera.
`)
}
