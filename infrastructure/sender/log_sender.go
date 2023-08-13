package sender

import (
	"log"

	"github.com/google/uuid"
)

type LogSender struct{}

func (ln *LogSender) Send(id uuid.UUID, total float64) error {
	log.Printf("sending the bill to the customer %s for a total amounf of %0.0f", id, total)

	return nil
}
