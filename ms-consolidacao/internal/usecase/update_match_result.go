package usecase

import (
	"context"
	"strconv"
	"strings"

	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/entity"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/repository"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type UpdateMatchResultInput struct {
	ID     string `json:"match_id"`
	Result string `json:"result"`
}

type UpdateMatchResultUseCase struct {
	Uow uow.UowInterface
}

func NewUpdateMatchResultUseCase(uow uow.UowInterface) *UpdateMatchResultUseCase {
	return &UpdateMatchResultUseCase{
		Uow: uow,
	}
}

func (u *UpdateMatchResultUseCase) Execute(ctx context.Context, input UpdateMatchResultInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		matchRepository := u.getMatchRepository(ctx)
		match, err := matchRepository.FindByID(ctx, input.ID)
		if err != nil {
			return err
		}

		matchResult := strings.Split(input.Result, "-")
		teamAResult, _ := strconv.Atoi(matchResult[0])
		teamBResult, _ := strconv.Atoi(matchResult[1])

		match.Result = *entity.NewMatchResult(teamAResult, teamBResult)
		err = matchRepository.Update(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *UpdateMatchResultUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := u.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}
