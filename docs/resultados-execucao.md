# Resultados de execução

Registro da execução local do laboratório autoral de força bruta em Go.

## Comandos

```bash
go test ./...
go vet ./...
go run ./labs/go/bruteforce-detector ./labs/go/bruteforce-detector/sample.log
```

## Resultado observado

```text
?    github.com/tiagovilasboas/cybersecurity-study-guide/labs/go/bruteforce-detector [no test files]
alerta: 203.0.113.10 teve 3 tentativas de login com falha
```

`go vet ./...` não encontrou problemas e, por isso, não produziu saída.

## Interpretação

O programa identificou o IP `203.0.113.10` ao atingir o limite de três falhas. O evento `SUCCESS` do IP `198.51.100.7` não gerou alerta. Isso demonstra o comportamento do protótipo, não uma validação completa de um ambiente de produção.

A nota, o progresso e a correção da atividade da Coursera continuam sendo validados na própria plataforma.
