package usecase

import (
	"context"

	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/entity"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/repository"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type AddMyTeamInput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type AddMyTeamUseCase struct {
	Uow uow.UowInterface
}

func NewAddMyTeamUseCase(uow uow.UowInterface) *AddMyTeamUseCase {
	return &AddMyTeamUseCase{
		Uow: uow,
	}
}

func (a *AddMyTeamUseCase) Execute(ctx context.Context, input AddMyTeamInput) error {
	myTeamRepository := a.getMyTeamRepository(ctx)
	myTeam := entity.NewMyTeam(input.ID, input.Name)
	err := myTeamRepository.Create(ctx, myTeam)

	if err != nil {
		return err
	}

	return a.Uow.CommitOrRollback()
}

func (a *AddMyTeamUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")

	if err != nil {
		panic(err)
	}

	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}
