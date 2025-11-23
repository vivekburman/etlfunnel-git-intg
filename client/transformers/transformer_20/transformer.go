package client_transformer_20

import (
	"etlfunnel/execution/models"
)
func Transformer(param *models.TransformerProps) (map[string]any, error) {
	// code here, something is fishy here bor
	return param.Record, nil
}