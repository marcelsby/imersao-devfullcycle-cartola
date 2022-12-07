package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/usecase"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type ProcessNewPlayer struct {
	ProcessEventStrategy
}

func (p ProcessNewPlayer) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.AddPlayerInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}

	addNewPlayerUseCase := usecase.NewAddPlayerUseCase(uow)

	err = addNewPlayerUseCase.Execute(ctx, input)

	if err != nil {
		return err
	}

	return nil
}
