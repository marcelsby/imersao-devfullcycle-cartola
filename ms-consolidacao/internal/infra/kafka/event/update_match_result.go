package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/usecase"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type ProcessUpdateMatchResult struct{}

func (p ProcessUpdateMatchResult) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.UpdateMatchResultInput

	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}

	updateMatchResultUsecase := usecase.NewUpdateMatchResultUseCase(uow)
	err = updateMatchResultUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
