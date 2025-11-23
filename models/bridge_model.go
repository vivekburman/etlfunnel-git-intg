package models

import (
	"database/sql"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-mysql-org/go-mysql/client"
	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/writeconcern"
)

// Models related to parameter to be passed to each core connector type for client connector
// Destination
type PostgresDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *pgx.Conn
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MySQLDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *client.Conn
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MariaDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *client.Conn
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MicrosoftServerDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *sql.DB
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type OracleDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *sql.DB
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type RedisDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *redis.Client
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type RabbitMQDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *amqp.Connection
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type ElasticDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *elasticsearch.Client
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MongoDestQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         *mongo.Client
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}

// Source
type MySQLSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *client.Conn
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}

type MySQLSourceQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *client.Conn
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MySQLSourceBinlog struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *client.Conn
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}

type MariaSourceBinlog struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *client.Conn
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MariaSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *client.Conn
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}

type MariaSourceQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *client.Conn
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}
type PostgresSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *pgx.Conn
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}
type PostgresSourceQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *pgx.Conn
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}

type PostgresSourceNotification struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *pgx.Conn
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type PostgresSourceWAL struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *pgx.Conn
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}

type MicrosoftServerSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *sql.DB
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}
type MicrosoftServerSourceQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *sql.DB
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MicrosoftServerSourceCDC struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *sql.DB
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MicrosoftServerSourceServiceBroker struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *sql.DB
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}

type OracleSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *sql.DB
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}
type OracleSourceQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *sql.DB
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type OracleSourceCDC struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *sql.DB
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type RedisSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *redis.Client
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}
type RabbitMQSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *amqp.Connection
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}
type MongoSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *mongo.Client
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}
type MongoSourceQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *mongo.Client
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MongoSourceStreams struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *mongo.Client
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MongoSourceOplog struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *mongo.Client
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type MongoSourceQueryTune struct {
	CommandDoc bson.D
}
type MongoStreamsTune struct {
	ChangeStreamOptions options.ChangeStreamOptionsBuilder
	Pipeline            []bson.M
}
type MongoWriteOperationType string

const (
	MongoWriteInsertOne  MongoWriteOperationType = "INSERT_ONE"
	MongoWriteInsertMany MongoWriteOperationType = "INSERT_MANY"
	MongoWriteUpdateOne  MongoWriteOperationType = "UPDATE_ONE"
	MongoWriteUpdateMany MongoWriteOperationType = "UPDATE_MANY"
	MongoWriteReplaceOne MongoWriteOperationType = "REPLACE_ONE"
	MongoWriteDeleteOne  MongoWriteOperationType = "DELETE_ONE"
	MongoWriteDeleteMany MongoWriteOperationType = "DELETE_MANY"
	MongoWriteBulkWrite  MongoWriteOperationType = "BULK_WRITE"
)

type MongoDBWriteOptions struct {
	Upsert           bool
	WriteConcern     *writeconcern.WriteConcern
	Ordered          bool
	BypassValidation bool
	Comment          any
	ArrayFilters     []any
	Hint             any
	Sort             any
	Let              any
	Collation        *options.Collation
}
type MongoSourceOplogTune struct {
	Filter  bson.M
	Options options.FindOptionsBuilder
}
type RedisSourceKeys struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *redis.Client
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type RedisSourceKeyspace struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *redis.Client
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type RedisSourceStreams struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *redis.Client
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}

type ElasticSourceQuery struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *elasticsearch.Client
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type ElasticSourceFetch struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *elasticsearch.Client
	AuxiliaryDBConnMap map[string]IDatabaseEngine
	DestDBConn         IDatabaseEngine
}
type DBElasticsearchQueryType string

const (
	ElasticsearchQueryTypeSearch   DBElasticsearchQueryType = "SEARCH"
	ElasticsearchQueryTypeScroll   DBElasticsearchQueryType = "SCROLL"
	ElasticsearchQueryTypeGet      DBElasticsearchQueryType = "GET"
	ElasticsearchQueryTypeMultiGet DBElasticsearchQueryType = "MGET"
	ElasticsearchQueryTypeSQL      DBElasticsearchQueryType = "SQL"
)

type DBElasticWriteOperationType string

const (
	ElasticWriteIndex  DBElasticWriteOperationType = "INDEX"
	ElasticWriteCreate DBElasticWriteOperationType = "CREATE"
	ElasticWriteUpdate DBElasticWriteOperationType = "UPDATE"
	ElasticWriteDelete DBElasticWriteOperationType = "DELETE"
)

type PostgresCDCOutputPluginType string

const (
	PostgresCDCTypePGOutput PostgresCDCOutputPluginType = "PG_OUTPUT"
	PostgresCDCTypeWAL2JSON PostgresCDCOutputPluginType = "WAL2JSON"
)

type MicrosoftServerCDCQueryType string

const (
	MicrosoftServerCDCTypeAllChanges MicrosoftServerCDCQueryType = "ALL_CHANGES"
	MicrosoftServerCDCTypeNetChanges MicrosoftServerCDCQueryType = "NET_CHANGES"
)

type RabbitMQSourceQueue struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *amqp.Connection
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type RabbitMQSourceExchange struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	SourceDBConn       *amqp.Connection
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
