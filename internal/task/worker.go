package task

import (
	"github.com/sptuan/stargazer/internal/dao"
	"github.com/sptuan/stargazer/internal/entity"
	"github.com/sptuan/stargazer/internal/model"
	"github.com/sptuan/stargazer/pkg/detector"
	"github.com/sptuan/stargazer/pkg/logger"
	"time"
)

type Worker struct {
	t           *entity.Task
	d           detector.Detector
	enableDbLog bool
	exit        chan struct{}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case <-w.exit:
				return
			case <-time.After(time.Duration(w.t.Interval) * time.Millisecond):
				msg, err := w.d.Detect()
				if err != nil {
					logger.Errorf("detect failed, task: %v, err: %v, msg: %v", w.t.Name, err, msg.Message)
				} else {
					logger.Infof("detect success, task: %v, msg: %v", w.t.Name, msg.Message)
				}
				msg.TaskId = int(w.t.ID)
				err = dao.AddDataLog(&msg)
				if err != nil {
					logger.Errorf("add data log failed, err: %v", err)
				}
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.exit <- struct{}{}
}

func NewWorker(t *entity.Task) (*Worker, error) {
	var d detector.Detector
	var err error
	switch model.ProbeType(t.Type) {
	case model.HTTP, model.HTTPS:
		d, err = detector.NewHttpDetector(t)
		if err != nil {
			return nil, err
		}
	case model.PING:
		d, err = detector.NewPingDetector(t)
		if err != nil {
			return nil, err
		}
	}
	return &Worker{
		t:    t,
		d:    d,
		exit: make(chan struct{}),
	}, nil
}
