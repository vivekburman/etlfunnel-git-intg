package models

type OrchestratorProps struct {
	Name               string
	SourceDBConn       IDatabaseEngine
	DestDBConn         IDatabaseEngine
	AuxiliaryDBConnMap map[string]IDatabaseEngine
}
type PipelineOrchestratorProps struct {
	Pipelines []OrchestratorProps
}
type FlowOrchestratorProps struct {
	Flows []OrchestratorProps
}
type FlowOrchestratorTune struct {
	ParentName   string
	ReplicaName  string
	ReplicaProps map[string]any
}
type PipelineOrchestratorTune struct {
	ParentName   string
	ReplicaName  string
	ReplicaProps map[string]any
}
