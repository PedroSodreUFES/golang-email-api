![Go](https://img.shields.io/badge/Language-Go-00ADD8?logo=go)

## API de Emails

### Contexto
+ API de emails
+ Contempla concorrência

### Tecnologias usadas
+ Go Language
+ PostgreSQL
+ Gin framework
+ Pacote errors
+ Pacote json
+ Pacote jwt
+ Pacote http
+ Pacote log
+ Pacote pgx/v5 
+ Pacote validator
+ SQLc
+ Tern
+ Docker

### Como rodar o programa
Após criar um arquivo .env, faça:
```bash
docker-compose up -d
go mod tidy
go run ./cmd/tern
sqlc generate -f internal/store/pgstore/sqlc.yml
go run ./cmd/api
```
OBS: Requer Docker, Go e SQLc instalados em seu computador

### Como parar o container Docker e excluir os dados
```bash
docker-compose down -v
```

### Diagrama relacional

![Diagrama relacional](diagrama_relacional.png)
