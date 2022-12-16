# Microserviço de Consolidação

Comando de migração do banco de dados do Microserviço de Consolidação

```bash
$ migrate -source file:///go/app/sql/migrations -database 'mysql://root:root@tcp(mysql:3306)/cartola' up
```