package task

import (
	"github.com/bocchi-the-cache/stargazer/internal/conf"
	"github.com/bocchi-the-cache/stargazer/internal/dao"
	"github.com/bocchi-the-cache/stargazer/internal/entity"
	"github.com/bocchi-the-cache/stargazer/pkg/logger"
	"reflect"
	"time"
)

var defaultManager *Manager

type Manager struct {
	// NOTE: add/delete huge number of tasks may cause memory leak, but it's ok for now
	// TODO: graceful shutdown
	ws map[string]*Worker
}

func (m *Manager) Init() {
	m.ws = make(map[string]*Worker)
}

func (m *Manager) Start() {
	// Refresh every 5 seconds
	go func() {
		for {
			time.Sleep(5 * time.Second)
			ts, err := dao.GetTasks()
			if err != nil {
				logger.Errorf("get tasks from database failed: %v", err)
				continue
			}
			m.Refresh(ts)
		}
	}()
	// Clear data logs
	go func() {
		for {
			cleanBefore := conf.Cfg.Service.LogCleanBefore
			if cleanBefore == 0 {
				cleanBefore = 24 * 7
			}
			t := time.Now().Add(time.Hour * time.Duration(cleanBefore*-1))
			err := dao.DeleteDataLogBeforeTime(int(t.Unix()))
			if err != nil {
				logger.Errorf("clear data logs failed: %v", err)
			}
			logger.Infof("clear log success! before: %v", t)
			time.Sleep(15 * time.Minute)
		}
	}()
}

func (m *Manager) Refresh(ts []*entity.Task) {
	// add new
	for _, t := range ts {
		if _, ok := m.ws[t.Name]; !ok {
			w, err := NewWorker(t)
			if err != nil {
				logger.Errorf("create worker failed: %v", err)
				continue
			}
			m.ws[t.Name] = w
			m.ws[t.Name].Start()
		}
	}
	// update exist
	for _, t := range ts {
		if w, ok := m.ws[t.Name]; ok {
			if deepEqual(t, w.t) {
				continue
			}
			w.Stop()
			nw, err := NewWorker(t)
			if err != nil {
				logger.Errorf("create worker failed: %v", err)
				continue
			}
			m.ws[t.Name] = nw
			m.ws[t.Name].Start()
		}
	}
	// remove old
	for k, v := range m.ws {
		var found bool
		for _, t := range ts {
			if t.Name == k {
				found = true
				break
			}
		}
		if !found {
			v.Stop()
			delete(m.ws, k)
		}
	}
}

func deepEqual(t1, t2 *entity.Task) bool {
	return reflect.DeepEqual(t1, t2)
}

func Init() {
	defaultManager = &Manager{}
	defaultManager.Init()
	defaultManager.Start()
}
