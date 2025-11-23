package cast

import (
	"etlfunnel/execution/models"
	"fmt"
	"reflect"
)

const connectorInstance = "ConnectorInstance"

// GetConnector extracts the typed connector from a database engine
// Works for any client type: *pgx.Conn, *client.Conn, *sql.DB, *redis.Client, etc.
func castDatabaseEngine[T any](engine models.IDatabaseEngine) (*T, error) {
	// Handle nil input
	if engine == nil {
		return nil, fmt.Errorf("database engine is nil")
	}

	// Get the actual value that the interface pointer points to
	engineValue := reflect.ValueOf(engine)
	if engineValue.Kind() != reflect.Ptr || engineValue.IsNil() {
		return nil, fmt.Errorf("invalid database engine: must be a non-nil pointer")
	}

	// Get the struct that the pointer points to
	structValue := engineValue.Elem()

	// Look for ConnectorInstance field
	field := structValue.FieldByName(connectorInstance)
	if !field.IsValid() || field.IsNil() {
		return nil, fmt.Errorf("invalid database engine: missing or nil ConnectorInstance")
	}

	// Try to convert to *T
	client, ok := field.Interface().(*T)
	if !ok {
		return nil, fmt.Errorf("failed to cast ConnectorInstance to %T", new(T))
	}

	return client, nil
}
