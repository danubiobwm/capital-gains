# Capital Gains Calculator

Uma aplicação de linha de comando em Go para cálculo de impostos sobre ganhos de capital em operações do mercado financeiro.

## 1. Estrutura do Projeto
```
capital-gains/
 ├── cmd/capital-gains/main.go       # Entrada da aplicação (CLI)
 ├── internal/
 │    ├── app/app.go                 # Orquestra leitura e escrita
 │    ├── domain/types.go            # Modelos (Operation, TaxOut)
 │    └── service/calculator.go      # Regras de negócio
 ├── test/examples_input.txt         # Exemplo de entrada
 ├── go.mod
 └── README.md
```
## 2. Pré-requisitos

Go 1.18+ (recomendado 1.25+)

Unix, Linux ou macOS (como solicitado no PDF)

Verifique sua versão:
```
go version

```

## 3. Compilar o projeto

No diretório raiz:
```
go build -o capital-gains ./cmd/capital-gains
```
Gerará o binário:
```
./capital-gains
```

## 4. Executar a aplicação
O programa lê cada linha como uma lista JSON de operações.

✔ Forma 1 — Usando redirecionamento de arquivo (recomendado)
```
./capital-gains < test/examples_input.txt
```
✔ Forma 2 — Digitando manualmente no terminal
```
./capital-gains
[{"operation":"buy","unit-cost":10,"quantity":100},{"operation":"sell","unit-cost":15,"quantity":50}]
```
Pressione Enter para enviar a linha.

Para encerrar a aplicação, envie uma linha vazia (como especificado no PDF):
```
<enter>
```

5. Exemplo de entrada e saída

Entrada:

```
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
 {"operation":"sell", "unit-cost":20.00, "quantity": 5000},
 {"operation":"sell", "unit-cost":5.00, "quantity": 5000}]

```

Saída:
```
[{"tax":0},{"tax":10000},{"tax":0}]

```
6. Executar os testes
A aplicação inclui testes unitários cobrindo todos os casos do PDF.

Execute:
```
go test ./...
```

Exemplo de saída esperada:
```
ok  	capital-gains/internal/service	0.120s
```

7. Como usar com múltiplas linhas (múltiplas simulações)

Cada linha é independente.

Entrada:
```
[{"operation":"buy","unit-cost":10,"quantity":100}]
[{"operation":"buy","unit-cost":10,"quantity":10000},{"operation":"sell","unit-cost":20,"quantity":5000}]
```
Saída:
```
[{"tax":0}]
[{"tax":0},{"tax":10000}]
```

8. Erros e Considerações

O PDF afirma que não haverá entradas inválidas, portanto:

O programa não valida campos ausentes

Não imprime mensagens adicionais

A única saída é o JSON dos impostos

9. Como limpar o build
```
rm capital-gains
```