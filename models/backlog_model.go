package models

type BacklogProps struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	FailureStage       FailureStage
	Records            []map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type BacklogTune struct {
	Action PipelineAction
}
