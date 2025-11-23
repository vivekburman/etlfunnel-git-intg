package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBMicrosoftServerSource interface {
	FetchRecords(param *models.MicrosoftServerSourceFetch) <-chan map[string]any
	GenerateQuery(param *models.MicrosoftServerSourceQuery) (*models.MicrosoftServerSourceQueryTune, error)
	GenerateCDC(param *models.MicrosoftServerSourceCDC) (*models.MicrosoftServerSourceCDCTune, error)
	GenerateServiceBroker(param *models.MicrosoftServerSourceServiceBroker) (*models.MicrosoftServerServiceBrokerTune, error)
}

type IClientDBMicrosoftServerDest interface {
	GenerateQuery(param *models.MicrosoftServerDestQuery) (*models.MicrosoftServerDestQueryTune, error)
}
