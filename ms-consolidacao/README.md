# Microserviço de Consolidação

## 🏃 Executar localmente

1. Na pasta do projeto suba os contêiners:

```bash
$ docker compose up -d
```

2. Acesse o contêiner do Golang:

```bash
$ docker compose exec goapp bash
```

3. Execute as migrações para criar o schema no banco de dados:

```bash
$ migrate -source file:///go/app/sql/migrations -database 'mysql://root:root@tcp(mysql:3306)/cartola' up
```

4. Execute a aplicação:

```bash
$ go run cmd/cartola/main.go
```