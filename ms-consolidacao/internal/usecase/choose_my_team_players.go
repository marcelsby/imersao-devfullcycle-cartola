package usecase

import (
	"context"

	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/repository"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/service"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type ChooseMyTeamPlayersInput struct {
	ID        string   `json:"my_team_id"`
	PlayersID []string `json:"players"`
}

type ChooseMyTeamPlayersUseCase struct {
	Uow uow.UowInterface
}

func NewChooseMyTeamPlayersUseCase(uow uow.UowInterface) *ChooseMyTeamPlayersUseCase {
	return &ChooseMyTeamPlayersUseCase{
		Uow: uow,
	}
}

func (u *ChooseMyTeamPlayersUseCase) Execute(ctx context.Context, input ChooseMyTeamPlayersInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		playerRepository := u.getPlayerRepository(ctx)
		myTeamRepository := u.getMyTeamRepository(ctx)

		myTeam, err := myTeamRepository.FindByID(ctx, input.ID)

		if err != nil {
			return err
		}

		newChoosenPlayers, err := playerRepository.FindAllByIDs(ctx, input.PlayersID)

		if err != nil {
			return err
		}

		service.ChoosePlayers(myTeam, newChoosenPlayers)

		err = myTeamRepository.SavePlayers(ctx, myTeam)

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (u *ChooseMyTeamPlayersUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := u.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (u *ChooseMyTeamPlayersUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := u.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
