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
	case 0:
		return s.minutes
	case 1:
		return s.hours
	case 2:
		return s.days
	case 3:
		return s.months
	case 4:
		return s.weeks
	default:
		return Specs{}
	}
}

// SetField sets time spec field
func (s *Schedule) SetField(num int, specs Specs) {
	switch num {
	case 0:
		s.minutes = specs
	case 1:
		s.hours = specs
	case 2:
		s.days = specs
	case 3:
		s.months = specs
	case 4:
		s.weeks = specs
	}
}

func parse(spec string, specs Specs) (Specs, error) {
	return nil, nil
}
