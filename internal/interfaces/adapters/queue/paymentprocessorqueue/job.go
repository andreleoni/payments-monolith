package paymentprocessorqueue

import (
	"fmt"
	"log/slog"
)

type Job struct {
	logger *slog.Logger
	id     string
}

func (j Job) Logger() *slog.Logger {
	return j.logger
}

func (j Job) Process() error {
	fmt.Println(j.id)

	return nil
}
