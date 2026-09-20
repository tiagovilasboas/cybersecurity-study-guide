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

**Status:** leitura concluída; relatório autoral em elaboração.

## Estrutura usada para o relatório

O modelo abaixo ajuda a organizar o pensamento sem fornecer uma resposta pronta:

1. protocolo observado e seu papel;
2. linha do tempo do incidente e evidências;
3. um controle recomendado, com justificativa, custo operacional e forma de validação.

Consulte também o [modelo de relatório](../templates/relatorio-incidente.md).
