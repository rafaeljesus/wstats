package worker

var RequestQueue chan Request

type Dispatcher struct {
	workerPool chan chan Request
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	RequestQueue = make(chan Request)

	return &Dispatcher{
		workerPool: make(chan chan Request, maxWorkers),
		maxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Run(fn func(string)) {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.workerPool)
		go worker.Register(fn)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case request := <-RequestQueue:
			go func(request Request) {
				requestChannel := <-d.workerPool
				requestChannel <- request
			}(request)
		}
	}
}
