package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBElasticSource interface {
	FetchRecords(param *models.ElasticSourceFetch) <-chan map[string]any
	GenerateQuery(request *models.ElasticSourceQuery) (*models.ElasticQueryTune, error)
}

type IClientDBElasticDest interface {
	GenerateQuery(param *models.ElasticDestQuery) (*models.ElasticDestQueryTune, error)
}
