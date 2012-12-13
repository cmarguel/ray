package render

type TaskRunner struct {
	NumRoutines int

	tasks chan Task
}

func NewTaskRunner(numRoutines int) TaskRunner {
	tasks := make(chan Task)
	return TaskRunner{numRoutines, tasks}
}

func (t TaskRunner) Start() {
	for i := 0; i < t.NumRoutines; i++ {
		go t.receiveTasks()
	}
}

func (t TaskRunner) Stop() {
	close(t.tasks)
}

func (t TaskRunner) AddTask(task Task) {
	t.tasks <- task
}

func (t TaskRunner) receiveTasks() {
	for task := range t.tasks {
		task.Run()
	}
}
