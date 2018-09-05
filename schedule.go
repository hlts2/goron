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
func NewSchedule(strSpecs []string) (*Schedule, error) {
	s := &Schedule{
		minutes: make(Specs, 0, DefaultSpecCount),
		hours:   make(Specs, 0, DefaultSpecCount),
		days:    make(Specs, 0, DefaultSpecCount),
		months:  make(Specs, 0, DefaultSpecCount),
		weeks:   make(Specs, 0, DefaultSpecCount),
	}

	for i, strSpec := range strSpecs {
		spec, err := parse(strSpec, s.Field(i))
		if err != nil {
			return nil, err
		}

		s.SetField(i, spec)
	}

	return s, nil
}

// TODO 実装方法を改善する

// Field returns time spec field
func (s *Schedule) Field(num int) Specs {
	switch num {
	case Minute:
		return s.minutes
	case Hour:
		return s.hours
	case Day:
		return s.days
	case Month:
		return s.months
	case Week:
		return s.weeks
	default:
		return Specs{}
	}
}

// SetField sets time spec field
func (s *Schedule) SetField(num int, specs Specs) {
	switch num {
	case Minute:
		s.minutes = specs
	case Hour:
		s.hours = specs
	case Day:
		s.days = specs
	case Month:
		s.months = specs
	case Week:
		s.weeks = specs
	}
}
