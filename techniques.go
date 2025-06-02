package guitar

import "fmt"

type Harmonic struct {
	Fret int

	String int

	Time float32
}

func (h Harmonic) TabSymbol() string {
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

func (s Slide) TabSymbol() string {
	if s.FretStart == -1 {
		return fmt.Sprintf("/%d", s.FretEnd)
	}
	return fmt.Sprintf("%d/%d", s.FretStart, s.FretEnd)
}

func (s Slide) StringNumber() int {
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

func (h HammerOn) TabSymbol() string {
	return fmt.Sprintf("%dh%d", h.FretFrom, h.FretTo)
}

func (h HammerOn) StringNumber() int {
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

func (p PullOff) TabSymbol() string {
	return fmt.Sprintf("%dp%d", p.FretFrom, p.FretTo)
}

func (p PullOff) StringNumber() int {
	return p.String
}

func (p PullOff) StartTime() float32 {
	return p.Time
}
