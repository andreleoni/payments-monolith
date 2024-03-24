package broker

import (
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BrokerController struct {
	logger slog.Logger
}

func NewBrokerController(logger slog.Logger) BrokerController {
	return BrokerController{logger: logger}
}

func (bc BrokerController) Enqueue(c *gin.Context) {
	logCorrelationID, logCorrelationIDExists := c.Get("logCorrelationID")

	// Parse body
	bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
	jsonMap := make(map[string]string)
	json.Unmarshal(bodyAsByteArray, &jsonMap)

	id := jsonMap["id"]
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "id not found",
		})

		return
	}

	contextlogger := bc.logger.With("id", id)

	if logCorrelationIDExists {
		contextlogger = bc.logger.With("correlation_id", logCorrelationID)
	}

	JobChannel <- Job{Logger: *contextlogger, ID: id}

	c.JSON(http.StatusOK, "")
}

var JobChannel = make(chan Job)

type Job struct {
	Logger slog.Logger
	ID     string
}

func (BrokerController) Consumer() {
	for job := range JobChannel {
		job.Logger.Info("Job processed",
			"id",
			job.ID)
	}
}
