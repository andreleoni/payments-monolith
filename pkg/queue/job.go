package queue

import "log/slog"

type Job interface {
	Logger() *slog.Logger
	Process() error
}
