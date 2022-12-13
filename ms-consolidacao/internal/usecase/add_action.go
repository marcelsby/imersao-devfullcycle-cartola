package usecase

import (
	"context"

	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/entity"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/repository"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type AddActionInput struct {
	MatchID  string `json:"match_id"`
	TeamID   string `json:"team_id"`
	PlayerID string `json:"player_id"`
	Minute   int    `json:"minute"`
	Action   string `json:"action"`
}

type AddActionUseCase struct {
	Uow         uow.UowInterface
	ActionTable entity.ActionTableInterface
}

func NewAddActionUseCase(uow uow.UowInterface, actionTable entity.ActionTableInterface) *AddActionUseCase {
	return &AddActionUseCase{
		Uow:         uow,
		ActionTable: actionTable,
	}
}

func (a *AddActionUseCase) Execute(ctx context.Context, input AddActionInput) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		matchRepository := a.getMatchRepository(ctx)
		myTeamRepository := a.getMyTeamRepository(ctx)
		playerRepository := a.getPlayerRepository(ctx)

		match, err := matchRepository.FindByID(ctx, input.MatchID)

		if err != nil {
			return err
		}

		score, err := a.ActionTable.GetScore(input.Action)

		if err != nil {
			return err
		}

		newAction := entity.NewGameAction(input.PlayerID, input.TeamID, input.Minute, input.Action, score)
		match.Actions = append(match.Actions, *newAction)

		err = matchRepository.SaveActions(ctx, match, float64(score))

		if err != nil {
			return err
		}

		player, err := playerRepository.FindByID(ctx, input.PlayerID)

		if err != nil {
			return err
		}

		player.Price += float64(score)
		err = playerRepository.Update(ctx, player)

		if err != nil {
			return err
		}

		myTeam, err := myTeamRepository.FindByID(ctx, "22087246-01bc-46ad-a9d9-a99a6d734167")

		if err != nil {
			return err
		}

		err = myTeamRepository.AddScore(ctx, myTeam, float64(score))

		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddActionUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := a.Uow.GetRepository(ctx, "MatchRepository")

	if err != nil {
		panic(err)
	}

	return matchRepository.(repository.MatchRepositoryInterface)
}

func (a *AddActionUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	MyTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")

	if err != nil {
		panic(err)
	}

	return MyTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (a *AddActionUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	PlayerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")

	if err != nil {
		panic(err)
	}

	return PlayerRepository.(repository.PlayerRepositoryInterface)
}
