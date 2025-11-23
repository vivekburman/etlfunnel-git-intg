package models

import (
	"context"
)

type IConnectionConfig any
type Record struct {
	Data map[string]any // User-facing data that goes through transformations
	Meta map[string]any // Internal metadata preserved throughout pipeline
}
type IReadByImpl struct {
	Channel    <-chan *Record
	Cleanup    func() error
	CommitHook func(records []*Record) PipelineAction
}

type IDatabaseEngine interface {
	GetName() string
	Connect(ctx context.Context, connectionConfig IConnectionConfig) error
	IsConnectionError(err error) bool
	Close(ctx context.Context) error
}

type PipelineAction int

const (
	ActionContinue PipelineAction = iota
	ActionStop
)

type FailureStage int

const (
	FailureStageNone FailureStage = iota
	FailureStageTransform
	FailureStageDestination
)
