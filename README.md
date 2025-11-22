# Capital Gains Calculator

Uma aplicação de linha de comando em Go para cálculo de impostos sobre ganhos de capital em operações do mercado financeiro.

## Arquitetura

O projeto segue os princípios de **Clean Architecture** e **SOLID**:

### Camadas:

1. **Domain**: Entidades e interfaces de repositório
2. **Application**: Casos de uso da aplicação
3. **Infrastructure**: Implementações concretas (CLI, JSON, Repositórios)
4. **CMD**: Ponto de entrada da aplicação

### Princípios SOLID aplicados:

- **S**: Cada classe tem uma única responsabilidade
- **O**: Aberto para extensão, fechado para modificação
- **L**: Substituível por suas abstrações
- **I**: Interfaces segregadas e específicas
- **D**: Dependência de abstrações, não implementações

## Como executar

### Pré-requisitos
- Go 1.25 ou superior
- Ou Docker

### Build e execução nativa

```bash
# Build
make build

# Executar
./capital-gains

# Ou com input redirection
./capital-gains < input.txt