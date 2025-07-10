# Ferramenta de Teste de Carga em Go

Este repositório contém uma ferramenta de linha de comando (CLI) desenvolvida em Go para a realização de testes de carga em serviços web. A aplicação é capaz de disparar um número configurável de requisições HTTP de forma concorrente, simulando múltiplos usuários, e ao final apresenta um relatório com as métricas da execução.

## Funcionalidades

- Execução de testes de carga via CLI
- Configuração de URL do alvo, número total de requisições e nível de concorrência
- Geração de relatório com tempo total, contagem de requisições por status e distribuição de códigos HTTP
- Implementação concorrente utilizando goroutines para simulação de chamadas simultâneas
- Containerização via Docker para portabilidade e execução em ambientes isolados

## Arquitetura

O projeto foi desenvolvido seguindo os princípios da Clean Architecture, separando as responsabilidades em camadas distintas (entidades, casos de uso e infraestrutura). Essa abordagem garante baixo acoplamento, alta coesão e maior testabilidade do código-fonte.

## Pré-requisitos

- Go: versão 1.22 ou superior
- Docker: para execução em contêineres

## Instalação e Execução

Existem duas maneiras principais de executar a ferramenta: localmente ou via Docker.

### Opção 1: Execução Local

Clone o repositório:

```bash
git clone <url-do-seu-repositorio>
cd stress-test
```

Compile a aplicação:

```bash
go build -o stress-tester ./cmd/cli
```

Execute a ferramenta:

```bash
./stress-tester --url=http://exemplo.com --requests=1000 --concurrency=100
```

### Opção 2: Execução via Docker

Construa a imagem Docker:

```bash
docker build -t stress-tester .
```

Execute o contêiner:

```bash
docker run --rm stress-tester --url=http://google.com --requests=1000 --concurrency=100
```

> **Nota**: Para testar um serviço local, utilize `host.docker.internal` no lugar de `localhost`.

## Uso

A ferramenta é configurada através de flags de linha de comando:

- `--url` (string, obrigatório): A URL completa do serviço web a ser testado
- `--requests` (int, obrigatório): O número total de requisições a serem realizadas
- `--concurrency` (int, obrigatório): O número de requisições simultâneas (trabalhadores/goroutines)

## Saída do Relatório

Ao final da execução, um relatório é impresso no console com as seguintes informações:

```
--- Relatório Final ---
Tempo total gasto: 1.52345s
Quantidade total de requests: 1000
Requests com status 200 (OK): 998
Distribuição de status HTTP:
  - Status 200: 998
  - Status 429: 2
---------------------
```

- **Tempo total gasto**: Duração total da execução do teste
- **Quantidade total de requests**: Total de requisições efetivamente realizadas
- **Requests com status 200 (OK)**: Número de respostas bem-sucedidas
- **Distribuição de status HTTP**: Contagem de respostas agrupadas por código de status HTTP

## Testes

O projeto contém uma suíte de testes automatizados para validar a lógica de negócio principal.

### Executando a Suíte de Testes

```bash
go test ./...
```

### Verificando a Cobertura de Código

Gerar perfil de cobertura:

```bash
go test -coverprofile=coverage.out ./...
```

Visualizar relatório em HTML:

```bash
go tool cover -html=coverage.out
```