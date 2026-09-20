# Detector de força bruta em Go

Exercício complementar para praticar o conceito de proteção contra tentativas repetidas de login.

O programa lê linhas no formato:

```text
DATA IP RESULTADO USUARIO
```

Ele conta os eventos `FAIL` por IP e emite um alerta quando o limite chega a três tentativas. Em um sistema real, o limite, a janela de tempo, o bloqueio e o tratamento de falsos positivos precisariam ser definidos com segurança.

## Executar

Na raiz do repositório:

```bash
go run ./labs/go/bruteforce-detector ./labs/go/bruteforce-detector/sample.log
```

Saída esperada:

```text
alerta: 203.0.113.10 teve 3 tentativas de login com falha
```

## Trade-offs

- um limite baixo pode bloquear usuários legítimos;
- contar apenas por IP não identifica bem usuários atrás de NAT;
- produção exige janela de tempo, logs confiáveis, alertas e resposta controlada.
