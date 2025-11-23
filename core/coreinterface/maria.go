package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBMariaSource interface {
	FetchRecords(param *models.MariaSourceFetch) <-chan map[string]any
	GenerateQuery(param *models.MariaSourceQuery) (*models.MariaSourceQueryTune, error)
	GenerateBinLog(param *models.MariaSourceBinlog) (*models.MariaSourceBinlogTune, error)
}
type IClientDBMariaDest interface {
	GenerateQuery(param *models.MariaDestQuery) (*models.MariaDestQueryTune, error)
}
