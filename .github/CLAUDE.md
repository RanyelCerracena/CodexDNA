# BioDigital Codec — Master Plan

Nada deve ser feito automático, tudo deve ser instruido, para que eu o desenvolva.

## 1. Visão Geral

O BioDigital Codec é um conversor de arquivos digitais para sequências de DNA usando Go. O objetivo é transformar dados binários em bases nucleotídicas de longo prazo de forma robusta, confiável e extensível.

O mapeamento utilizado é baseado em base 4, onde cada par de bits representa uma base de DNA:

| Bits | Base |
| ---- | ---- |
| `00` | `A`  |
| `01` | `C`  |
| `10` | `G`  |
| `11` | `T`  |

A conversão de dados segue uma lógica atômica:

- cada byte é decomposto em pares de bits,
- cada par de bits é convertido na base correspondente,
- a sequência resultante é preparada para preservação biológica.

## 2. Arquitetura

A estrutura do projeto segue princípios de Clean Architecture / Hexagonal Architecture em Go:

```text
.
├── cmd/
│   ├── biodigital-codec/
│   │   └── main.go
│   └── cli/
│       └── root.go
├── internal/
│   ├── core/
│   │   ├── service.go
│   │   ├── model.go
│   │   └── usecase.go
│   ├── encoder/
│   │   ├── encode.go
│   │   ├── decode.go
│   │   └── rules.go
│   ├── repository/
│   │   ├── postgres/
│   │   │   └── storage.go
│   │   └── metadata.go
│   └── app/
│       ├── server.go
│       └── config.go
├── deployments/
│   ├── docker-compose.yml
│   └── Dockerfile
├── docs/
│   └── CLAUDE.md
└── go.mod
```

- `cmd/` contém o entrypoint da aplicação e a CLI.
- `internal/core` concentra a lógica de domínio e os casos de uso principais.
- `internal/encoder` cuida da conversão bit ↔ base e das regras biológicas.
- `internal/repository` abstrai persistência e metadados.
- `deployments/` contém a infraestrutura de containerização e orquestração.

## 3. Stack Tecnológica

- Go (Golang) — linguagem principal para alto desempenho e portabilidade.
- PostgreSQL — banco de dados relacional para persistência de metadados e artefatos.
- Docker — ambiente reproducível e containerização do serviço.
- Metodologias aplicadas:
  - Clean Architecture / Hexagonal Architecture
  - TDD e testes de unidade/integrados
  - CI/CD orientado a qualidade de código e integração contínua

## 4. Plano de Sprints

### Sprint 1 — Core Engine (Bit-to-Base)

Objetivo: construir o motor de conversão inicial e uma CLI básica.

- Desenvolver conversão binária para DNA usando o mapeamento `00=A`, `01=C`, `10=G`, `11=T`.
- Implementar decodificação reversa de DNA para binário.
- Criar CLI com comandos de `encode`, `decode` e opções de entrada/saída.
- Validar consistência com casos de teste básicos.

### Sprint 2 — Estabilidade Biológica

Objetivo: aplicar regras de qualidade biológica para a sequência de DNA.

- Adicionar detecção de homopolímeros excessivos.
- Implementar filtros para evitar longas repetições da mesma base.
- Validar e ajustar conteúdo GC para manter equilíbrio entre 40% e 60%.
- Garantir que a sequência convertida seja compatível com princípios de síntese e leitura.

### Sprint 3 — Resiliência

Objetivo: tornar o codec tolerante a erros com correção de erros.

- Implementar Error Correction usando Reed-Solomon.
- Encapsular dados com códigos de paridade e redundância.
- Validar a capacidade de recuperar dados após perda ou corrupção parcial.
- Testar com simulação de erros em diferentes comprimentos de bloco.

### Sprint 4 — Persistência e Futuro

Objetivo: integrar armazenamento de metadados e preparar o projeto para expansão.

- Conectar a camada de persistência a PostgreSQL.
- Modelar metadados de preservação: hash, timestamp, tipo de arquivo, versão do codec.
- Criar funções de ingestão e recuperação de artefatos codificados.
- Definir diretrizes para futuras integrações (e.g. armazenamento em nuvem, conversão reversa automatizada, suporte a múltiplos formatos).

## 5. Critérios de Sucesso

Um arquivo codificado é considerado “viável” quando:

- A conversão de ida e volta preserva 100% dos dados originais.
- As sequências geradas atendem a regras biológicas básicas:
  - sem homopolímeros longos;
  - conteúdo GC equilibrado;
  - ausência de padrões que causem falhas na síntese.
- O mecanismo de correção de erros recupera dados mesmo após simulação de perda/corrupção.
- Metadados essenciais são persistidos corretamente em PostgreSQL.
- A CLI oferece comandos claros e reproduzíveis.
