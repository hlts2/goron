package goron

import (
	"sync"
	"time"
)

// DefaultJobCount is the number of job
const DefaultJobCount int = 1000

// const
const (
	Minute = iota
	Hour
	Day
	Month
	Week
)

type (
	// JobHandler represents job handler type for goron
	JobHandler func() error

	// Job represents job info
	Job struct {
		g       *goron
		spec    []string
		handler JobHandler
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
		JobCount() int
		Run()
	}

	goron struct {
		mu   *sync.Mutex
		spec []string
		jobs Jobs
	}
)

// New returns Goron(*goron) object
func New() Goron {
	return &goron{
		mu:   new(sync.Mutex),
		jobs: make(Jobs, 0, DefaultJobCount),
		spec: initSpec(),
	}
}

func initSpec() []string {
	return []string{"*", "*", "*", "*"}
}

func (g *goron) Minute(spec string) Goron {
	if len(g.jobs) < Minute {
		return g
	}

	g.spec[Minute] = spec
	return g
}

func (g *goron) Hour(spec string) Goron {
	if len(g.jobs) < Hour {
		return g
	}

	g.spec[Hour] = spec
	return g
}

func (g *goron) Day(spec string) Goron {
	if len(g.jobs) < Day {
		return g
	}

	g.spec[Day] = spec
	return g
}

func (g *goron) Month(spec string) Goron {
	if len(g.jobs) < Month {
		return g
	}

	g.spec[Month] = spec
	return g
}

func (g *goron) Week(spec string) Goron {
	if len(g.jobs) < Week {
		return g
	}

	g.spec[Week] = spec
	return g
}

func (g *goron) With(handlers ...JobHandler) {
	g.addJob(g.spec, handlers...)
	g.spec = initSpec()
}

func (g *goron) AddJob(spec string, handlers ...JobHandler) {
	g.addJob(g.spec, handlers...)
	g.spec = initSpec()
}

func (g *goron) addJob(spec []string, handlers ...JobHandler) {
	for _, handler := range handlers {
		g.jobs = append(g.jobs, Job{
			g:       g,
			spec:    spec,
			handler: handler,
			finish:  make(chan bool),
		})
	}
}

func (g *goron) JobCount() int {
	return len(g.jobs)
}

func (g *goron) Run() {
	t := time.Now()

	for _, job := range g.jobs {
		go job.run(t)
	}
}

func (j Job) run(t time.Time) {
END_LOOP:
	for {
		select {
		case _ = <-j.finish:
			break END_LOOP
		default:
		}
	}
}
