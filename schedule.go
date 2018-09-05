package goron

// DefaultSpecCount represents the default spec count
const DefaultSpecCount = 5

// const value for date
const (
	Minute = iota
	Hour
	Day
	Month
	Week
)

type (
	// Spec is spec for goron
	Spec struct {
		min      int
		max      int
		wildcard string
	}

	// Specs is spec slice type
	Specs []Spec

	// Schedule represents time schedule
	Schedule struct {
		minutes Specs
		hours   Specs
		days    Specs
		months  Specs
		weeks   Specs
	}
)

// NewSchedule returns *Schedule object
func NewSchedule(specs []string) (*Schedule, error) {
	s := &Schedule{
		minutes: make(Specs, 0, DefaultSpecCount),
		hours:   make(Specs, 0, DefaultSpecCount),
		days:    make(Specs, 0, DefaultSpecCount),
		months:  make(Specs, 0, DefaultSpecCount),
		weeks:   make(Specs, 0, DefaultSpecCount),
	}

	return s, nil
}

func parse(spec string, specs Specs) (Specs, error) {
	return nil, nil
}
