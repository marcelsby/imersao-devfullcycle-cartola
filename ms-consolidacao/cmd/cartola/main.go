package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/infra/db"
	httphandler "github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/infra/http"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/infra/kafka/consumer"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/infra/repository"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dbConnection, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/cartola?parseTime=true")
	if err != nil {
		panic(err)
	}

	defer dbConnection.Close()

	uow, err := uow.NewUow(ctx, dbConnection)
	if err != nil {
		panic(err)
	}

	registerRepositories(uow)

	router := chi.NewRouter()

	router.Get("/players", httphandler.ListPlayersHandler(ctx, *db.New(dbConnection)))
	router.Get("/my-teams/{teamID}/players", httphandler.ListMyTeamPlayers(ctx, *db.New(dbConnection)))
	router.Get("/my-teams/{teamID}/balance", httphandler.GetMyTeamBalanceHandler(ctx, *db.New(dbConnection)))
	router.Get("/matches", httphandler.ListMatchesHandler(ctx, repository.NewMatchRepository(dbConnection)))
	router.Get("/matches/{matchID}", httphandler.ListMatchByIDHandler(ctx, repository.NewMatchRepository(dbConnection)))

	go http.ListenAndServe(":8080", router)

	var topics = []string{"newMatch", "chooseTeam", "newPlayer", "updateMatchResult", "newAction"}
	msgChan := make(chan *kafka.Message)
	go consumer.Consume(topics, "broker:9094", msgChan)
	consumer.ProcessEvents(ctx, msgChan, uow)
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
