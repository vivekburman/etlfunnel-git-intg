package coreinterface

import (
	"etlfunnel/execution/models"
)

type IClientDBRedisSource interface {
	GenerateKeys(param *models.RedisSourceKeys) (*models.RedisSourceKeysTune, error)
	GenerateStreams(param *models.RedisSourceStreams) (*models.RedisSourceStreamsTune, error)
	GenerateKeyspace(param *models.RedisSourceKeyspace) (*models.RedisSourceKeySpacesTune, error)
	FetchRecords(param *models.RedisSourceFetch) <-chan map[string]any
}

type IClientDBRedisDest interface {
	GenerateQuery(param *models.RedisDestQuery) (*models.RedisDestQueryTune, error)
}
