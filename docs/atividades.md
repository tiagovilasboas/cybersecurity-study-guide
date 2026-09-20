# Registro das atividades

Este arquivo registra o que foi praticado e aprendido. Os textos são autorais e não substituem a atividade original da Coursera.

## Atividade: aplicar técnicas de proteção do sistema operacional

**Cenário estudado:** um site de receitas foi comprometido depois que um atacante adivinhou a senha administrativa padrão. O código da página foi alterado para induzir visitantes a baixar um executável, que redirecionava para um domínio malicioso.

**O que a atividade exercita:**

- interpretar uma sequência de eventos de DNS e HTTP;
- separar evidência técnica de suposição;
- documentar um incidente com linguagem objetiva;
- recomendar um controle proporcional ao risco de força bruta.

**Como raciocinar:**

- DNS resolve nomes; a comunicação web busca a página e o conteúdo solicitado.
- A documentação deve registrar quem foi afetado, como o incidente foi descoberto, quais evidências existem e qual foi a sequência observada.
- A correção deve reduzir tentativas automatizadas e impedir que uma senha padrão seja suficiente para comprometer a conta.

**Status:** relatório autoral gerado; validação e envio continuam pendentes na Coursera.

## Estrutura usada para o relatório

O modelo abaixo ajuda a organizar o pensamento sem fornecer uma resposta pronta:

1. protocolo observado e seu papel;
2. linha do tempo do incidente e evidências;
3. um controle recomendado, com justificativa, custo operacional e forma de validação.

Consulte também o [modelo de relatório](../templates/relatorio-incidente.md).

## Passo a passo da atividade no Coursera

A atividade foi analisada seguindo esta sequência:

1. **Ler o cenário:** o site `yummyrecipesforme.com` foi comprometido após um ataque de força bruta contra a conta administrativa.
2. **Observar o comportamento:** em um ambiente sandbox, o navegador acessa o site, recebe uma solicitação para baixar um executável e é redirecionado para `greatrecipesforme.com`.
3. **Analisar o `tcpdump`:** o fluxo mostra resolução DNS para os dois domínios e requisições HTTP para carregar a página, baixar o arquivo e acessar o domínio falso.
4. **Identificar os protocolos:** registrar no relatório o papel de DNS e HTTP na sequência observada.
5. **Documentar o incidente:** descrever ativo afetado, acesso obtido, alteração no JavaScript, impacto nos visitantes, descoberta e fontes de evidência.
6. **Recomendar um controle:** escolher uma medida contra força bruta, como MFA, senha forte, limite de tentativas ou monitoramento, explicando como validar o resultado.
7. **Comparar com o exemplo do curso:** usar o material de apoio depois de elaborar o próprio relatório.
8. **Validar na plataforma:** confirmar a atividade na Coursera somente depois de revisar o texto e concluir o envio pelo próprio usuário.

## Laboratórios e materiais analisados

| Material | Origem | Uso no estudo |
| --- | --- | --- |
| Atividade “Aplicar técnicas de proteção do sistema operacional” | Coursera | cenário, critérios e sequência do incidente |
| Registro de tráfego `tcpdump` | material de apoio da Coursera | identificação de DNS e HTTP |
| Guia “Como ler o registro do tcpdump” | material de apoio da Coursera | interpretação das camadas e dos pacotes |
| Modelo de relatório de incidente | template do curso | estrutura das três seções do relatório |
| Detector de força bruta em Go | laboratório autoral | prática reproduzível com log local |
| Gerador de relatório em Go | laboratório autoral | consolidação do contexto e da evidência calculada |

Os materiais do curso sustentam a análise do caso. Os programas em Go ajudam a praticar e reproduzir uma parte do raciocínio, mas o `sample.log` é sintético e não representa o `tcpdump` oficial.
