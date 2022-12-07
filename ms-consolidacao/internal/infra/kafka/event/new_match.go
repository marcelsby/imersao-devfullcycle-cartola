package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/usecase"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type ProcessNewMatch struct{}

func (p ProcessNewMatch) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.AddMatchInput

	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}

	addNewMatchUsecase := usecase.NewAddMatchUseCase(uow)
	err = addNewMatchUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
