package cast

import (
	"database/sql"
	"etlfunnel/execution/models"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/jackc/pgx/v5"
)

func CastAsPostgresDBConnection(engine models.IDatabaseEngine) (*pgx.Conn, error) {
	return castDatabaseEngine[pgx.Conn](engine)
}

func CastAsMySQLDBConnection(engine models.IDatabaseEngine) (*client.Conn, error) {
	return castDatabaseEngine[client.Conn](engine)
}

func CastAsMariaDBConnection(engine models.IDatabaseEngine) (*client.Conn, error) {
	return castDatabaseEngine[client.Conn](engine)
}

func CastAsMicrosoftServerDBConnection(engine models.IDatabaseEngine) (*sql.DB, error) {
	return castDatabaseEngine[sql.DB](engine)
}

func CastAsOracleDBConnection(engine models.IDatabaseEngine) (*sql.DB, error) {
	return castDatabaseEngine[sql.DB](engine)
}
