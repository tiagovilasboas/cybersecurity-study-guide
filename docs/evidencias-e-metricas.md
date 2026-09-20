# Evidências e métricas de estudo

Este arquivo separa o que pode ser demonstrado pelo repositório do que só pode ser confirmado no Coursera.

## Regra de evidência

O GitHub registra o processo autoral: objetivo, raciocínio, artefato produzido, critério de validação e status. A plataforma Coursera continua sendo a fonte oficial para progresso, nota, conclusão e certificado.

Não registramos uma atividade como correta apenas porque o texto foi escrito. O status só muda depois de uma validação observável na plataforma.

## Matriz atual

| Atividade | Artefato esperado | Métrica verificável | Fonte | Status |
| --- | --- | --- | --- | --- |
| Vídeos de proteção do sistema operacional | conteúdo assistido | item marcado como concluído | Coursera | concluído |
| Leitura sobre força bruta e hardening | leitura compreendida | leitura marcada como concluída | Coursera | concluído |
| Aplicar técnicas de hardening | relatório de incidente | protocolo identificado, incidente documentado e correção justificada | relatório autoral + validação da atividade | em elaboração |
| Teste de conhecimentos do módulo | respostas próprias | resultado exibido após envio | Coursera | pendente |
| Hardening de rede | atividade do módulo | artefato e resultado da plataforma | Coursera + registro autoral | pendente |
| Hardening de nuvem | atividade do módulo | artefato e resultado da plataforma | Coursera + registro autoral | pendente |

## Métricas que fazem sentido

- itens concluídos na plataforma;
- atividades entregues;
- resultado ou percentual exibido pelo Coursera;
- artefatos produzidos;
- critérios cobertos pelo artefato;
- correções feitas após feedback;
- data da última validação.

Evite métricas de vaidade, como quantidade de páginas ou linhas escritas. O indicador mais importante é a combinação entre evidência, validação e aprendizado demonstrável.

## Evidências deste repositório

| Registro | Origem | O que prova | Limite |
| --- | --- | --- | --- |
| `docs/atividades.md` | cenário e materiais da Coursera | compreensão do caso e dos critérios | não prova envio ou nota |
| `relatorio.md` | gerado pelo programa em Go | relatório estruturado para revisão | não substitui a entrega oficial |
| `labs/go/bruteforce-detector/sample.log` | exemplo autoral | entrada reproduzível do detector | não é log real do curso |
| `docs/resultados-execucao.md` | execução local | comando e saída observados | não é métrica de produção |

O relatório mistura contexto conhecido do caso com uma evidência calculada pelo detector. A parte dinâmica é recalculada sempre que outro log é informado:

```bash
go run ./labs/go/incident-report caminho/para/log | tee relatorio.md
```

## Como atualizar

Para cada atividade concluída, registre:

1. o artefato produzido;
2. o critério que ele atende;
3. o resultado observado no Coursera;
4. a data da verificação;
5. uma reflexão curta sobre o que foi aprendido.
