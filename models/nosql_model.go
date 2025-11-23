package models

import "time"

type RedisSourceKeysTune struct {
	SpecificKeyList []string
	KeyPatterns     []string
	ScanCount       int
}
type RedisSourceStreamsTune struct {
	StreamNames     []string
	ConsumerGroup   string
	ConsumerName    string
	SpecificStartId string
	StartFrom       string
	BatchSize       int
	BlockTime       int
	AutoAck         bool
	ClaimMinIdle    int
}

type RedisSourceKeySpacesTune struct {
	NotificationTypes []string
	KeyPatterns       []string
	Database          int
	SubscriptionMode  string
}
type RedisDestQueryTune struct {
	Operation       string
	Key             string
	Value           any
	Expiration      time.Duration
	RecordsPerBatch int
}
type MongoDestQueryTune struct {
	Operation       MongoWriteOperationType
	Query           any
	Payload         any
	Options         MongoDBWriteOptions
	RecordsPerBatch int
}
type ElasticQueryTune struct {
	QueryType     DBElasticsearchQueryType
	Index         string
	Body          any
	ScrollTimeout time.Duration
	DocumentID    string
}
type ElasticDestQueryTune struct {
	Index           string
	DocID           string
	Operation       DBElasticWriteOperationType
	Document        map[string]any
	Upsert          map[string]any
	RefreshPolicy   string
	Script          string
	ScriptParams    map[string]any
	RecordsPerBatch int
}

type RabbitMQSourceQueueTune struct {
	QueueName     string
	ConsumerTag   string
	AutoAck       bool
	PrefetchCount int
}

type RabbitMQSourceExchangeTune struct {
	ExchangeName  string
	ExchangeType  string
	QueueName     string
	RoutingKeys   []string
	AutoAck       bool
	PrefetchCount int
}

type RabbitMQDestQueryTune struct {
	ExchangeName    string
	RoutingKey      string
	QueueName       string
	Mandatory       bool
	Immediate       bool
	ContentType     string
	DeliveryMode    uint8
	Priority        uint8
	Expiration      string
	Headers         map[string]any
	Body            []byte
	RecordsPerBatch int
}
