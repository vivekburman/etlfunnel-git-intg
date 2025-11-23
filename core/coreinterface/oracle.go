package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBOracleSource interface {
	FetchRecords(param *models.OracleSourceFetch) <-chan map[string]any
	GenerateQuery(param *models.OracleSourceQuery) (*models.OracleSourceQueryTune, error)
	GenerateCDC(param *models.OracleSourceCDC) (*models.OracleSourceCDCTune, error)
}

type IClientDBOracleDest interface {
	GenerateQuery(param *models.OracleDestQuery) (*models.OracleDestQueryTune, error)
}
