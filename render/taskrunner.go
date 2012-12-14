package render

type TaskRunner struct {
	NumRoutines int
	Close       chan int

	TasksDone chan int
	tasks     chan Task
}

func NewTaskRunner(numRoutines int) TaskRunner {
	tasks := make(chan Task)

	return TaskRunner{numRoutines, make(chan int, numRoutines), make(chan int, numRoutines), tasks}
}

func (t TaskRunner) Start() {
	for i := 0; i < t.NumRoutines; i++ {
		go t.receiveTasks()
	}
}

func (t TaskRunner) Stop() {
	close(t.tasks)
}

func (t TaskRunner) Wait() {
	for i := 0; i < t.NumRoutines; i++ {
		<-t.Close
	}
	close(t.TasksDone)
}

func (t TaskRunner) AddTask(task Task) {
	t.tasks <- task
}

func (t TaskRunner) receiveTasks() {
	for task := range t.tasks {
		task.Run()
		t.TasksDone <- 1
	}
	t.Close <- 1
}
