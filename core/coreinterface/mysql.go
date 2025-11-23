package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBMySQLSource interface {
	FetchRecords(param *models.MySQLSourceFetch) <-chan map[string]any
	GenerateQuery(param *models.MySQLSourceQuery) (*models.MySQLSourceQueryTune, error)
	GenerateBinLog(pram *models.MySQLSourceBinlog) (*models.MySQLSourceBinlogTune, error)
}

type IClientDBMySQLDest interface {
	GenerateQuery(param *models.MySQLDestQuery) (*models.MySQLDestQueryTune, error)
}
