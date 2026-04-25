# Sprint 1 — Core Engine (Bit-to-Base)

Este documento descreve os passos a serem implementados para concluir o objetivo 1 da Sprint 1 do projeto `BioDigital Codec`.

## Objetivo

Construir o motor básico de conversão entre dados binários e DNA, e entregar uma CLI mínima com comandos `encode` e `decode`.

## Passos

1. Estruturar o repositório
   - Criar ou verificar as pastas:
     - `cmd/biodigital-codec/`
     - `internal/encoder/`
     - `internal/core/`
     - `internal/repository/`
   - Inicializar ou atualizar `go.mod` com o módulo do projeto.
   - Criar `cmd/biodigital-codec/main.go` e `cmd/biodigital-codec/root.go`.

2. Implementar o encoder básico
   - No pacote `internal/encoder`, criar os arquivos:
     - `encode.go`
     - `decode.go`
     - `rules.go`
   - Implementar funções principais:
     - `Encode(data []byte) (string, error)`
     - `Decode(dna string) ([]byte, error)`
   - O mapeamento deve ser:
     - `00 -> A`
     - `01 -> C`
     - `10 -> G`
     - `11 -> T`
   - A lógica deve:
     - converter cada byte em 8 bits;
     - agrupar os bits em pares;
     - mapear cada par para a base correspondente;
     - montar a sequência de DNA resultante.
   - A decodificação deve reverter o processo e recuperar os bytes originais.

3. Escrever testes unitários básicos
   - Criar `internal/encoder/encode_test.go`.
   - Validar casos simples:
     - `[]byte{0x00}` -> `AAAA`
     - `[]byte{0xFF}` -> `TTTT`
     - `[]byte(
