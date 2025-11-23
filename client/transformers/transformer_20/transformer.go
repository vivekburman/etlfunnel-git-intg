package client_transformer_20

import (
	"etlfunnel/execution/models"
)
func Transformer(param *models.TransformerProps) (map[string]any, error) {
	// code here, attempt conflict #2

	return param.Record, nil
}