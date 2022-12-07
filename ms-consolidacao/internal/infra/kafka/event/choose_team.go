package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/usecase"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type ProcessChooseTeam struct{}

func (p ProcessChooseTeam) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.ChooseMyTeamPlayersInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}

	chooseMyTeamPlayersUsecase := usecase.NewChooseMyTeamPlayersUseCase(uow)
	err = chooseMyTeamPlayersUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
