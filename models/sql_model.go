package models

import (
	"time"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/replication"
)

type PostgresSourceWALTune struct {
	SlotName        string
	OutputPlugin    PostgresCDCOutputPluginType
	Streaming       bool
	PublicationName string
}
type MySQLDestQueryTune struct {
	Query           string
	Value           []any
	RecordsPerBatch int
}
type MariaDestQueryTune struct {
	Query           string
	Value           []any
	RecordsPerBatch int
}
type PostgresDestQueryTune struct {
	Query           string
	Value           []any
	RecordsPerBatch int
}
type OracleDestQueryTune struct {
	Query           string
	Value           []any
	RecordsPerBatch int
}
type MicrosoftServerDestQueryTune struct {
	Query           string
	Value           []any
	RecordsPerBatch int
}
type MySQLSourceQueryTune struct {
	Query string
}
type MicrosoftServerSourceQueryTune struct {
	Query string
}
type OracleSourceQueryTune struct {
	Query           string
	RecordsPerBatch int
	PrefetchSize    int
}
type MariaSourceQueryTune struct {
	Query string
}
type PostgresSourceQueryTune struct {
	Query string
}
type PostgresSourceNotificationTune struct {
	ChannelName string
}
type MySQLSourceBinlogTune struct {
	ServerID uint32
}
type MariaSourceBinlogTune struct {
	ServerID uint32
}
type MysqlConnImpl struct {
	Conn     *client.Conn
	Syncer   *replication.BinlogSyncer
	Streamer *replication.BinlogStreamer
}
type MariaSQLConnImpl struct {
	Conn     *client.Conn
	Syncer   *replication.BinlogSyncer
	Streamer *replication.BinlogStreamer
}
type OracleSourceCDCTune struct {
	SourceTables           []string
	SCNType                string
	StartSCN               uint64
	StartTime              time.Time
	ExtractionMode         string
	IncludeOperations      []string
	BatchSize              int
	PollingInterval        time.Duration
	SessionRefreshMode     string
	SessionRefreshInterval time.Duration
	SessionRefreshCount    int
	MaxRetries             int
	BaseRetryDelayMs       int
	MaxRetryDelayMs        int
	RetryJitter            float64
}

type MicrosoftServerServiceBrokerTune struct {
	QueueName  string
	SchemaName string
	Timeout    int // -1 means never timeout
}

type MicrosoftServerSourceCDCTune struct {
	FromLSN      string
	ToLSN        string
	StartTime    time.Time
	EndTime      time.Time
	UseMinMaxLSN bool
	QueryType    MicrosoftServerCDCQueryType
	InstanceName string
	RowFilter    string
}
