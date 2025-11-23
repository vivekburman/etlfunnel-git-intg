package client_transformer_20

import (
	"etlfunnel/execution/models"
)
func Transformer(param *models.TransformerProps) (map[string]any, error) {
	// code here
	return param.Record, nil
}
