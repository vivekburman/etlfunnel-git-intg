package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBRabbitMQSource interface {
	GenerateQueue(param *models.RabbitMQSourceQueue) (*models.RabbitMQSourceQueueTune, error)
	GenerateExchange(param *models.RabbitMQSourceExchange) (*models.RabbitMQSourceExchangeTune, error)
	FetchRecords(param *models.RabbitMQSourceFetch) <-chan map[string]any
}

type IClientDBRabbitMQDest interface {
	GenerateQuery(param *models.RabbitMQDestQuery) (*models.RabbitMQDestQueryTune, error)
}
