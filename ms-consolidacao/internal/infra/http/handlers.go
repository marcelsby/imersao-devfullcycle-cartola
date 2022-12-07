package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	irepository "github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/repository"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/infra/db"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/infra/presenter"
)

func ListPlayersHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		players, err := queries.FindAllPlayers(ctx)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(players)
	}
}

func ListMyTeamPlayers(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		teamID := chi.URLParam(req, "teamID")

		myTeamPlayers, err := queries.GetPlayersByMyTeamID(ctx, teamID)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(myTeamPlayers)
	}
}

func ListMatchesHandler(ctx context.Context, matchRepository irepository.MatchRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		matches, err := matchRepository.FindAll(ctx)
		var matchesPresenter presenter.Matches
		for _, match := range matches {
			matchesPresenter = append(matchesPresenter, presenter.NewMatchPresenter(match))
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(matchesPresenter)
	}
}

func ListMatchByIDHandler(ctx context.Context, matchRepository irepository.MatchRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchID := chi.URLParam(r, "matchID")
		match, err := matchRepository.FindByID(ctx, matchID)
		matchPresenter := presenter.NewMatchPresenter(match)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(matchPresenter)
	}
}

func GetMyTeamBalanceHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamID := chi.URLParam(r, "teamID")
		balance, err := queries.GetMyTeamBalance(ctx, teamID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resultJson := map[string]float64{"balance": balance}
		json.NewEncoder(w).Encode(resultJson)
	}
}
