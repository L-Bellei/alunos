# 🚀 API RESTful em Go com Gin e PostgreSQL

API RESTful desenvolvida em **Go (Golang)** utilizando o framework **Gin** para roteamento de alta performance e **GORM** como ORM para persistência de dados em um banco de dados **PostgreSQL**. O projeto inclui suporte a operações CRUD completas, criação de registros em lote e gerenciamento do banco de dados via **Docker Compose**.

---

## 🛠️ Tecnologias Utilizadas

- **[Go](https://go.dev/)** (v1.20+) — Linguagem de programação principal
- **[Gin Framework](https://gin-gonic.com/)** — Framework web de alta performance
- **[GORM](https://gorm.io/)** — ORM para manipulação do PostgreSQL
- **[PostgreSQL](https://www.postgresql.org/)** — Banco de dados relacional
- **[pgAdmin 4](https://www.pgadmin.org/)** — Interface visual para gerenciamento do banco
- **[Docker & Docker Compose](https://www.docker.com/)** — Conteinerização do banco de dados e pgAdmin

---

## 📁 Estrutura do Projeto

```text
api-gin-rest/
├── controllers/    # Controladores das rotas (regras de negócio da API)
├── database/       # Conexão com o banco de dados e automigração
├── models/         # Estrutura de dados (Entidade Aluno)
├── routes/         # Definição dos endpoints REST
├── docker-compose.yml  # Configuração dos contêineres PostgreSQL e PgAdmin
├── go.mod          # Módulo e dependências do Go
├── go.sum          # Checksums das dependências
└── main.go         # Ponto de entrada da aplicação
```

---

## 📌 Endpoints da API

A API roda por padrão na porta `:8080`.

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| `GET` | `/alunos` | Retorna a lista com todos os alunos |
| `GET` | `/alunos/:id` | Retorna os dados de um aluno por ID |
| `GET` | `/alunos/cpf/:cpf` | Retorna os dados de um aluno por CPF |
| `POST` | `/alunos` | Cadastra um novo aluno |
| `POST` | `/alunos-em-lote` | Cadastra uma lista de alunos em lote (batches de 10) |
| `PATCH` | `/alunos/:id` | Atualiza os dados de um aluno existente |
| `DELETE` | `/alunos/:id` | Remove um aluno pelo ID |

---

## 📝 Exemplo de Estrutura de Dados (JSON)

### Cadastrar Aluno (`POST /alunos`)
```json
{
  "nome": "João Silva",
  "cpf": "12345678901",
  "rg": "123456789"
}
```

### Cadastrar em Lote (`POST /alunos-em-lote`)
```json
[
  {
    "nome": "Maria Souza",
    "cpf": "98765432100",
    "rg": "987654321"
  },
  {
    "nome": "Carlos Oliveira",
    "cpf": "11122233344",
    "rg": "111222333"
  }
]
```

---

## ⚡ Como Executar o Projeto

### Pré-requisitos
- **[Go](https://go.dev/dl/)** instalado na máquina.
- **[Docker Desktop](https://www.docker.com/products/docker-desktop/)** instalado e em execução.

### 1. Clonar o Repositório
```bash
git clone https://github.com/L-Bellei/api-gin-rest.git
cd api-gin-rest
```

### 2. Iniciar o Banco de Dados (PostgreSQL + pgAdmin)
Utilize o Docker Compose para subir os serviços em segundo plano:
```bash
docker-compose up -d
```
> **Acesso ao pgAdmin:**  
> Acesse no navegador: `http://localhost:54321`  
> - **Email:** `leo.bellei.dev@outlook.com`  
> - **Senha:** `leobellei`

### 3. Executar a Aplicação Go
Com o banco de dados rodando, inicie o servidor Go:
```bash
go run main.go
```

A API estará pronta para receber requisições em `http://localhost:8080`.
