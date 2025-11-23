package cast

import (
	"etlfunnel/execution/models"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CastAsRedisDBConnection(engine models.IDatabaseEngine) (*redis.Client, error) {
	return castDatabaseEngine[redis.Client](engine)
}
func CastAsMongoDBConnection(engine models.IDatabaseEngine) (*mongo.Client, error) {
	return castDatabaseEngine[mongo.Client](engine)
}
func CastAsElasticsearchConnection(engine models.IDatabaseEngine) (*elasticsearch.Client, error) {
	return castDatabaseEngine[elasticsearch.Client](engine)
}
