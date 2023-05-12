package workerpool

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkerPool(t *testing.T) {
	pool := NewWorkerPool(3)
	assert.Equal(t, 3, pool.Size())

	var tasks []Task
	for i := 0; i < 10; i++ {
		tasks = append(tasks, &testTask{i})
	}

	pool.AddTasks(tasks)
	pool.Wait()
}

type testTask struct {
	id int
}

func (task *testTask) Do() {
	fmt.Printf("Task %d is running\n", task.id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d is done\n", task.id)
}

func (task *testTask) ID() int {
	return task.id
}

func (task *testTask) String() string {
	return fmt.Sprintf("Task %d", task.id)
}
