package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBMongoSource interface {
	GenerateQuery(param *models.MongoSourceQuery) (*models.MongoSourceQueryTune, error)
	GenerateStream(param *models.MongoSourceStreams) (*models.MongoStreamsTune, error)
	GenerateOplogTrailing(param *models.MongoSourceOplog) (*models.MongoSourceOplogTune, error)
	FetchRecords(param *models.MongoSourceFetch) <-chan map[string]any
}
type IClientDBMongoDest interface {
	GenerateQuery(param *models.MongoDestQuery) (*models.MongoDestQueryTune, error)
}
