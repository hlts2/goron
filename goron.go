package goron

import (
	"strings"
	"time"
)

// DefaultJobCount represents the default job count
const DefaultJobCount int = 1000

type (
	// JobHandler represents job handler type for goron
	JobHandler func() error

	// Job represents job info
	Job struct {
		g        *goron
		schedule *Schedule
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
		Week(strSpec string) Goron

		// Month sets spec of month
		Month(strSpec string) Goron

		// Day sets sepec of day
		Day(strSpec string) Goron

		// Hour sets spec of hour
		Hour(strSpec string) Goron

		// Minute sets spec of minute
		Minute(strSpec string) Goron

		// With register job handler
		With(handlers ...JobHandler)

		// AddJob register job handler with spec
		AddJob(strSpec string, handlers ...JobHandler)

		// JobCount returns the number of job
		JobCount() int

		// Run starts deamon
		Run()
	}

	goron struct {
		strSpecs []string
		jobs     Jobs
	}
)

// New returns Goron(*goron) object
func New() Goron {
	return &goron{
		strSpecs: initSpec(),
		jobs:     make(Jobs, 0, DefaultJobCount),
	}
}

func initSpec() []string {
	return []string{"*", "*", "*", "*"}
}

func (g *goron) Minute(strSpec string) Goron {
	if len(g.jobs) < Minute {
		return g
	}

	g.strSpecs[Minute] = strSpec
	return g
}

func (g *goron) Hour(strSpec string) Goron {
	if len(g.jobs) < Hour {
		return g
	}

	g.strSpecs[Hour] = strSpec
	return g
}

func (g *goron) Day(strSpec string) Goron {
	if len(g.jobs) < Day {
		return g
	}

	g.strSpecs[Day] = strSpec
	return g
}

func (g *goron) Month(strSpec string) Goron {
	if len(g.jobs) < Month {
		return g
	}

	g.strSpecs[Month] = strSpec
	return g
}

func (g *goron) Week(strSpec string) Goron {
	if len(g.jobs) < Week {
		return g
	}

	g.strSpecs[Week] = strSpec
	return g
}

func (g *goron) With(handlers ...JobHandler) {
	err := g.addJob(g.strSpecs, handlers...)
	if err != nil {
		panic(err)
	}

	g.strSpecs = initSpec()
}

func (g *goron) AddJob(strSpec string, handlers ...JobHandler) {
	err := g.addJob(strings.Split(strSpec, " "), handlers...)
	if err != nil {
		panic(err)
	}

	g.strSpecs = initSpec()
}

func (g *goron) addJob(strSpecs []string, handlers ...JobHandler) error {
	schedule, err := NewSchedule(strSpecs)
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
