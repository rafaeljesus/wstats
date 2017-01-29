package worker

type Request struct {
	Payload string
}

type Worker struct {
	Pool    chan chan Request
	Channel chan Request
}

func NewWorker(pool chan chan Request) Worker {
	return Worker{
		Pool:    pool,
		Channel: make(chan Request),
	}
}

func (w Worker) Register(fn func(string)) {
	for {
		w.Pool <- w.Channel
		select {
		case request := <-w.Channel:
			fn(request.Payload)
		}
	}
}
