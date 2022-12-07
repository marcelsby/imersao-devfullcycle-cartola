package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/domain/entity"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/internal/usecase"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type ProcessNewAction struct{}

func (p ProcessNewAction) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.AddActionInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}

	actionTable := entity.ActionTable{}
	actionTable.Init()

	addNewActionUsecase := usecase.NewAddActionUseCase(uow, &actionTable)
	err = addNewActionUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
