package goron

import (
	"sync"
)

// DefaultJobCount is the number of job
const DefaultJobCount int = 1000

type (
	// JobHandler represents job handler type for goron
	JobHandler func() error

	// Job represents job info
	Job struct {
		g       *goron
		Spec    string
		Handler JobHandler
		finish  chan bool
	}

	// Jobs is Job slice type
	Jobs []Job
)

type (

	// Goron is base goron interface
	Goron interface {
		Week(spec string) Goron
		Month(spec string) Goron
		Day(spec string) Goron
		Hour(spec string) Goron
		Minute(spec string) Goron
		With(handlers ...JobHandler)
		AddJob(spec string, handlers ...JobHandler)
		Run()
	}

	goron struct {
		mu   *sync.Mutex
		spec string
		jobs Jobs
	}
)

// New returns Goron(*goron) object
func New() Goron {
	return &goron{
		mu:   new(sync.Mutex),
		jobs: make(Jobs, 0, DefaultJobCount),
	}
}

func (g *goron) Week(spec string) Goron {
	return g
}

func (g *goron) Month(spec string) Goron {
	return g
}

func (g *goron) Day(spec string) Goron {
	return g
}

func (g *goron) Hour(spec string) Goron {
	return g
}

func (g *goron) Minute(spec string) Goron {
	return g
}

func (g *goron) With(handlers ...JobHandler) {
	g.mu.Lock()
	g.addJob(g.spec, handlers...)
	g.spec = ""
	g.mu.Unlock()
}

func (g *goron) AddJob(spec string, handlers ...JobHandler) {
	g.mu.Lock()
	g.addJob(spec, handlers...)
	g.spec = ""
	g.mu.Unlock()
}

func (g *goron) addJob(spec string, handlers ...JobHandler) {
	for _, handler := range handlers {
		g.jobs = append(g.jobs, Job{
			Handler: handler,
			Spec:    spec,
			finish:  make(chan bool),
			g:       g,
		})
	}
}

func (g *goron) Run() {
	for _, job := range g.jobs {
		go job.run()
	}
}

func (j Job) run() {
END_LOOP:
	for {
		select {
		case _ = <-j.finish:
			break END_LOOP
		default:
		}
	}
}
