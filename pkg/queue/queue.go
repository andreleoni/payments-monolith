package queue

import "log/slog"

type Queue struct {
	jobs chan Job
}

func NewQueue() Queue {
	return Queue{jobs: make(chan Job)}
}

func (q Queue) Enqueue(job Job) {
	q.jobs <- job
}

func (q Queue) Consumer(logger *slog.Logger) {
	logger.Info("Starting consumer...")

	for job := range q.jobs {
		job.Logger().Info(
			"Starting process job",
			"job", job)

		processResult := job.Process()

		if processResult != nil {
			job.Logger().Error(
				"Error given on processing job",
				"job", job)
		} else {
			job.Logger().Info(
				"Processed with success",
				"job", job)
		}
	}
}
