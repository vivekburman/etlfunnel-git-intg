package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBPostgresSource interface {
	FetchRecords(param *models.PostgresSourceFetch) <-chan map[string]any
	GenerateQuery(param *models.PostgresSourceQuery) (*models.PostgresSourceQueryTune, error)
	GenerateNotification(param *models.PostgresSourceNotification) (*models.PostgresSourceNotificationTune, error)
	GenerateWAL(param *models.PostgresSourceWAL) (*models.PostgresSourceWALTune, error)
}

type IClientDBPostgresDest interface {
	GenerateQuery(param *models.PostgresDestQuery) (*models.PostgresDestQueryTune, error)
}
