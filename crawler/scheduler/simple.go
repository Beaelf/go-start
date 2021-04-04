package scheduler

import "go-start/crawler/engine"

type SimpleScheduer struct {
	workChannel chan engine.Request
}

func (s *SimpleScheduer) Submit(request engine.Request) {
	go func() {
		s.workChannel <- request
	}()
}

func (s *SimpleScheduer) WorkerChan() chan engine.Request {
	return s.workChannel
}

func (s *SimpleScheduer) WorkerReady(request chan engine.Request) {
}

func (s *SimpleScheduer) Run() {
	s.workChannel = make(chan engine.Request)
}
