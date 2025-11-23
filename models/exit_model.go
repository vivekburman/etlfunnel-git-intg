package models

import (
	"time"
)

type TerminateRuleProps struct {
	Ctx    IPipelineContextContract
	Logger ILoggerContract
}

type CustomTerminateRuleCheckProps struct {
	Ctx           IPipelineContextContract
	Logger        ILoggerContract
	TotalMessages uint64
	LastMessageAt time.Time
	StartTime     time.Time
}
type TerminateRuleActionTune struct {
	Action PipelineAction
	Reason string
}
type TerminateRuleTune struct {
	MaxRecords           *uint64
	IdleTimeout          *time.Duration
	MaxPipelineTime      *time.Duration
	UserDefinedCheckFunc func(*CustomTerminateRuleCheckProps) (*TerminateRuleActionTune, error)
	CheckInterval        time.Duration
}
