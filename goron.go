package goron

import (
	"strings"
	"time"
)

// DefaultJobCount is the number of job
const DefaultJobCount int = 1000

type (
	// JobHandler represents job handler type for goron
	JobHandler func() error

	// Job represents job info
	Job struct {
		g        *goron
		schedule Schedule
		handler  JobHandler
		finish   chan bool
	}

	// Jobs is Job slice type
	Jobs []Job
)

type (

	// Goron is base goron interface
	Goron interface {

		// Week sets spec of week
		Week(spec string) Goron

		// Month sets spec of month
		Month(spec string) Goron

		// Day sets sepec of day
		Day(spec string) Goron

		// Hour sets spec of hour
		Hour(spec string) Goron

		// Minute sets spec of minute
		Minute(spec string) Goron

		// With register job handler
		With(handlers ...JobHandler)

		// AddJob register job handler with spec
		AddJob(spec string, handlers ...JobHandler)

		// JobCount returns the number of job
		JobCount() int

		// Run starts deamon
		Run()
	}

	goron struct {
		spec []string
		jobs Jobs
		err  error
	}
)

// New returns Goron(*goron) object
func New() Goron {
	return &goron{
		spec: initSpec(),
		jobs: make(Jobs, 0, DefaultJobCount),
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
	if g.err != nil {
		panic(g.err)
	}

	err := g.addJob(g.spec, handlers...)
	if err != nil {
		panic(err)
	}

	g.spec = initSpec()
	g.err = nil
}

func (g *goron) AddJob(spec string, handlers ...JobHandler) {
	if g.err != nil {
		panic(g.err)
	}

	err := g.addJob(strings.Split(spec, " "), handlers...)
	if err != nil {
		panic(err)
	}

	g.spec = initSpec()
	g.err = nil
}

func (g *goron) addJob(spec []string, handlers ...JobHandler) error {
	schedule, err := parse(spec)
	if err != nil {
		return err
	}

	for _, handler := range handlers {
		g.jobs = append(g.jobs, Job{
			g:        g,
			schedule: schedule,
			handler:  handler,
			finish:   make(chan bool),
		})
	}

	return nil
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
