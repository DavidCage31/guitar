package guitar

import "fmt"

type Harmonic struct {
	Fret int

	String int

	Time float32
}

func (h Harmonic) FretPosition() string {
	return fmt.Sprintf("<%d>", h.Fret)
}

func (h Harmonic) StringPosition() int {
	return h.String
}

func (h Harmonic) StartTime() float32 {
	return h.Time
}

type Slide struct {
	FretStart int
	FretEnd   int

	String int

	Time float32
}

func (s Slide) FretPosition() string {
	if s.FretStart == -1 {
		return fmt.Sprintf("/%d", s.FretEnd)
	}
	return fmt.Sprintf("%d/%d", s.FretStart, s.FretEnd)
}

func (s Slide) StringPosition() int {
	return s.String
}

func (s Slide) StartTime() float32 {
	return s.Time
}

type HammerOn struct {
	FretFrom int
	FretTo   int

	String int

	Time float32
}

func (h HammerOn) FretPosition() string {
	return fmt.Sprintf("%dh%d", h.FretFrom, h.FretTo)
}

func (h HammerOn) StringPosition() int {
	return h.String
}

func (h HammerOn) StartTime() float32 {
	return h.Time
}

type PullOff struct {
	FretFrom int
	FretTo   int

	String int

	Time float32
}

func (p PullOff) FretPosition() string {
	return fmt.Sprintf("%dp%d", p.FretFrom, p.FretTo)
}

func (p PullOff) StringPosition() int {
	return p.String
}

func (p PullOff) StartTime() float32 {
	return p.Time
}
