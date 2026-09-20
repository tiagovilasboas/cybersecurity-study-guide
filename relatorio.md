# Rascunho — Security Incident Report

> Revise e adapte o texto antes de usar na atividade da Coursera.

## 1. Protocolo observado

O tráfego mostrou consultas DNS para resolver os domínios **yummyrecipesforme.com** e **greatrecipesforme.com**. Também foram observadas requisições HTTP para carregar a página, baixar o arquivo malicioso e redirecionar os usuários.

## 2. Descrição do incidente

O site **yummyrecipesforme.com** foi comprometido depois que o atacante descobriu a senha padrão da conta administrativa usando várias tentativas de login.

Após obter acesso, o atacante alterou o código JavaScript do site para induzir os visitantes a baixar um arquivo executável. O arquivo redirecionava os usuários para **greatrecipesforme.com**.

O impacto foi a exposição dos visitantes a malware e o redirecionamento para um site controlado pelo atacante.

## 3. Recomendação

Remover senhas padrão, exigir senhas fortes e habilitar autenticação multifator para a conta administrativa. Também é recomendável limitar tentativas de login e bloquear temporariamente endereços IP com comportamento de força bruta.

A eficácia pode ser validada monitorando novas tentativas falhas de login e confirmando que acessos repetidos são bloqueados ou geram alertas.

## Evidência complementar do laboratório em Go

Log analisado: `labs/go/bruteforce-detector/sample.log`

Resultado observado: `alerta: 203.0.113.10 teve 3 tentativas de login com falha`

A evidência acima é calculada a partir do log informado e complementa o relatório; ela não substitui a validação da atividade na Coursera.
