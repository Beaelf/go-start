package scheduler

import "go-start/crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueueScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(workerQ) > 0 && len(requestQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-q.requestChan:
				// send r to worker
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				// send next request to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()

}
