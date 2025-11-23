package models

type TransformerProps struct {
	Ctx                IPipelineContextContract
	Logger             ILoggerContract
	Record             map[string]any
	SourceDBConn       IDatabaseEngine
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
