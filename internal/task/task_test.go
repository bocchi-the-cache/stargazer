package task

import (
	"github.com/bocchi-the-cache/stargazer/internal/db"
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"testing"
	"time"
)

func getMockTasks() []*entity.Task {
	// generate tasks
	task1 := &entity.Task{
		Name:     "task1",
		Type:     "http",
		Target:   "www.baidu.com",
		Interval: 1000,
		Timeout:  1000,
	}
	task2 := &entity.Task{
		Name:     "task2",
		Type:     "https",
		Target:   "www.baidu.com",
		Interval: 1000,
		Timeout:  1000,
	}
	task3 := &entity.Task{
		Name:     "task3",
		Type:     "ping",
		Target:   "1.1.1.1",
		Interval: 1000,
		Timeout:  1000,
	}
	ts := []*entity.Task{task1, task2, task3}
	return ts
}

func TestManager(t *testing.T) {

	_ = db.Init("/tmp/stargazer.db")

	m := Manager{}
	m.Init()
	ts1 := getMockTasks()
	m.Refresh(ts1)

	ts2 := getMockTasks()
	ts2[0].Name = "task1_new" //new & delete
	ts2[1].Type = "ping"      //modify
	m.Refresh(ts2)
	time.Sleep(10 * time.Second)
}
