package core

type Pipeline struct {
	core *Core
}

func NewPipeline(core *Core) *Pipeline {
	pipeline := &Pipeline{
		core: core,
	}
	return pipeline
}

func (pipeline *Pipeline) RunPipeline() error {
	pipeline.runApiWorkflowMonitor()
	pipeline.runWorkflowMonitor()
	pipeline.runTaskMonitor()
	return nil
}

func (pipeline *Pipeline) runTaskMonitor() {
	go pipeline.core.Task.taskMonitor()
}

func (pipeline *Pipeline) runApiWorkflowMonitor() {
	go pipeline.core.Task.workApiMonitor()
}
func (pipeline *Pipeline) runWorkflowMonitor() {
	go pipeline.core.Task.workflowMonitor()
}
