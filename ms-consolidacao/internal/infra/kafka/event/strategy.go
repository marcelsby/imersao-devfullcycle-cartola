package event

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcelsby/imersao-devfullcycle-cartola-consolidacao/pkg/uow"
)

type ProcessEventStrategy interface {
	Process(cts context.Context, msg *kafka.Message, uow uow.UowInterface) error
}
