package models

type CheckpointProps struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Records            []map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type CheckpointTune struct {
	Action PipelineAction
}
