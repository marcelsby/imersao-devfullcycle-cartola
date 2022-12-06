package main

import (
	"context"
	"database/sql"

	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/infra/db"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/infra/repository"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

func main() {
	ctx := context.Background()

	dbConnection, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cartola?parseTime=true")
	if err != nil {
		panic(err)
	}

	defer dbConnection.Close()

	uow, err := uow.NewUow(ctx, dbConnection)
	if err != nil {
		panic(err)
	}

	registerRepositories(uow)

}

func registerRepositories(uow *uow.Uow) {
	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repository := repository.NewPlayerRepository(uow.Db)
		repository.Queries = db.New(tx)
		return repository
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repository := repository.NewMatchRepository(uow.Db)
		repository.Queries = db.New(tx)
		return repository
	})

	uow.Register("TeamRepository", func(tx *sql.Tx) interface{} {
		repository := repository.NewTeamRepository(uow.Db)
		repository.Queries = db.New(tx)
		return repository
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repository := repository.NewMyTeamRepository(uow.Db)
		repository.Queries = db.New(tx)
		return repository
	})
}
