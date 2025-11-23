package models

import (
	"context"

	"go.uber.org/zap"
)

type IPipelineContextContract interface {
	GetName() string
	GetContext() context.Context
	CancelContext()
	GetFlowReplicaProps() map[string]any
	GetReplicaProps() map[string]any
}
type IFlowContextContract interface {
	GetName() string
	GetContext() context.Context
	CancelContext()
	GetReplicaProps() map[string]any
}
type ILoggerContract interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	ErrorWithNotify(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}
