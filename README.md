# Cybersecurity Study Guide

Guia autoral de estudo do curso **Security Hardening**, realizado no Coursera, com foco em fundamentos de cibersegurança, proteção de sistemas operacionais, hardening de rede e segurança em nuvem.

> Este repositório registra aprendizados, modelos mentais e reflexões. Não redistribui vídeos, leituras, notebooks, avaliações ou respostas da Coursera.

## Por que este guia existe?

Hardening é como preparar uma casa antes de uma viagem: fechar entradas desnecessárias, trocar chaves fracas, limitar quem pode entrar e manter registros do que aconteceu. O objetivo é reduzir a superfície de ataque antes que um incidente aconteça.

## Percurso

| Etapa | O que estudamos | Evidência autoral |
| --- | --- | --- |
| Fundamentos | ameaças, controles e modelo CIA | [modelos mentais](docs/modelos-mentais.md) |
| Sistema operacional | configuração segura, força bruta e resposta a incidente | [atividades](docs/atividades.md) |
| Rede | redução de exposição e controles de comunicação | [checklist](docs/checklist-hardening.md) |
| Nuvem | identidade, configuração e responsabilidade compartilhada | [checklist](docs/checklist-hardening.md) |
| Revisão | conectar controles, evidências e decisões | [trajetória](docs/trajetoria.md) |
| Evidências | métricas, critérios e fonte de validação | [matriz de evidências](docs/evidencias-e-metricas.md) |

## Evidências e limites

Este repositório separa três tipos de registro:

- **Evidência do curso:** cenário, protocolos e critérios descritos nos materiais oficiais da Coursera.
- **Evidência autoral:** análise escrita, checklist e relatório produzido para consolidar o aprendizado.
- **Evidência executável:** resultados calculados pelos programas em Go a partir de um log local de exemplo.

O laboratório em Go é reproduzível, mas o `sample.log` é sintético. Ele demonstra o detector e não deve ser apresentado como captura real do ambiente da Coursera. Progresso, nota, conclusão e certificado continuam sendo confirmados na plataforma oficial.

## O que fica registrado

- resumos em linguagem simples;
- atividades descritas como estudos de caso, sem respostas copiadas;
- checklists reutilizáveis em projetos próprios;
- trade-offs: segurança, disponibilidade, custo, operação e experiência do usuário;
- dúvidas e decisões que ajudam a transformar teoria em prática;
- [rascunho gerado do relatório](relatorio.md);
- [resultados de execução](docs/resultados-execucao.md).

## Fontes oficiais

- [Curso Security Hardening no Coursera](https://www.coursera.org/learn/google-security-hardening)
- [Google Cybersecurity Certificate](https://www.coursera.org/professional-certificates/google-cybersecurity)
- [OWASP](https://owasp.org/)
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)

## Status

A documentação e os laboratórios autorais estão registrados. A validação da atividade, o progresso e a conclusão do curso devem ser confirmados na Coursera.

## Licença

Código, modelos e textos autorais: MIT. A licença não se aplica ao conteúdo proprietário da Coursera ou do Google.
