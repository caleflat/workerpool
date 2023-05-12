package workerpool

// WorkerPool is a worker pool
type WorkerPool struct {
	size    int
	workers []*Worker
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(size int) *WorkerPool {
	pool := &WorkerPool{
		size:    size,
		workers: make([]*Worker, size),
	}

	for i := 0; i < size; i++ {
		pool.workers[i] = NewWorker()
	}

	return pool
}

// Size returns the size of the worker pool
func (pool *WorkerPool) Size() int {
	return pool.size
}

// AddTask adds a task to the worker pool
func (pool *WorkerPool) AddTask(task Task) {
	worker := pool.getWorker()
	worker.AddTask(task)
}

// AddTasks adds tasks to the worker pool
func (pool *WorkerPool) AddTasks(tasks []Task) {
	for _, task := range tasks {
		pool.AddTask(task)
	}
}

// Wait waits for all tasks to be completed
func (pool *WorkerPool) Wait() {
	for _, worker := range pool.workers {
		worker.Wait()
	}
}

// getWorker returns a worker with the least number of tasks
func (pool *WorkerPool) getWorker() *Worker {
	worker := pool.workers[0]
	for _, w := range pool.workers {
		if w.NumTasks() < worker.NumTasks() {
			worker = w
		}
	}
	return worker
}

// Task is a task
type Task interface {
	Do()
}

// Worker is a worker
type Worker struct {
	tasks chan Task
	done  chan bool

	numTasks int
}

// NewWorker creates a new worker
func NewWorker() *Worker {
	worker := &Worker{
		tasks: make(chan Task),
		done:  make(chan bool),
	}

	go worker.run()
	return worker
}

// AddTask adds a task to the worker
func (worker *Worker) AddTask(task Task) {
	worker.tasks <- task
	worker.numTasks++
}

// NumTasks returns the number of tasks in the worker
func (worker *Worker) NumTasks() int {
	return worker.numTasks
}

// Wait waits for all tasks to be completed
func (worker *Worker) Wait() {
	<-worker.done
}

// run runs the worker
func (worker *Worker) run() {
	for task := range worker.tasks {
		task.Do()
		worker.numTasks--
	}

	worker.done <- true
}
