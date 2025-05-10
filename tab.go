package guitar

import (
	"errors"
	"fmt"
	"strings"
)

type InstrumentType int

const (
	GuitarType InstrumentType = iota
)

type tabBuilder struct {
	time     float32
	timeStep float32

	instrumentType InstrumentType
	tabStrings     []strings.Builder
}

type NotePositioner interface {
	FretPosition() string
	StringPosition() int
	StartTime() float32
}

func NewTabBuilder(instrumentType InstrumentType, tuningNotes []string) (*tabBuilder, error) {
	defaultTimeStep := float32(0.2)
	switch instrumentType {
	case GuitarType:
		tb := tabBuilder{
			time:           0.0,
			timeStep:       defaultTimeStep,
			instrumentType: instrumentType,
			tabStrings:     make([]strings.Builder, 6),
		}
		tb.addNotes(tuningNotes)

		return &tb, nil
	default:
		return nil, errors.ErrUnsupported
	}
}

func (tb *tabBuilder) Tab() string {
	tab := strings.Builder{}

	for i := range len(tb.tabStrings) {
		tab.WriteString(tb.tabStrings[i].String() + "\n")
	}

	return tab.String()
}

func (tb *tabBuilder) WriteNotes(notes ...NotePositioner) error {
	time := notes[0].StartTime()
	for i, n := range notes {
		if n.StartTime() < tb.time {
			return fmt.Errorf("note time %v precedes current time %v iteration %d note %+v",
				n.StartTime(), tb.time, i, n)
		}
		if n.StartTime() != time {
			return fmt.Errorf("notes time are not equal")
		}
		if n.StringPosition() >= len(tb.tabStrings) {
			return fmt.Errorf("invalid string index %d, in tab builder only %d strings",
				n.StringPosition(), len(tb.tabStrings))
		}
	}

	silence := int((time - tb.time) / tb.timeStep)
	tb.addSilence(silence)
	tb.time += tb.timeStep * (float32(silence + 1))

	maxLen := -1

	for _, n := range notes {
		stringPos := n.StringPosition()
		tb.tabStrings[stringPos].WriteString(n.FretPosition())
		if tb.tabStrings[stringPos].Len() > maxLen {
			maxLen = tb.tabStrings[stringPos].Len()
		}
	}

	for i := range tb.tabStrings {
		if tb.tabStrings[i].Len() < maxLen {
			diffLen := maxLen - tb.tabStrings[i].Len()
			for j := 0; j < diffLen; j++ {
				tb.tabStrings[i].WriteString("-")
			}
		}
	}

	// to escape situations like:
	// E|-3--123-----
	tb.addSilence(1)

	return nil
}

// TODO
// func (tb *tabBuilder) WriteChord() {
//
// }

func (tb *tabBuilder) addNotes(notes []string) error {
	if len(notes) != len(tb.tabStrings) {
		return fmt.Errorf("invalid tuning notes count")
	}

	for i := range notes {
		tb.tabStrings[i].WriteString(notes[i] + "|")
	}
	return nil
}

func (tb *tabBuilder) addSilence(n int) {
	for range n {
		for i := range len(tb.tabStrings) {
			tb.tabStrings[i].WriteString("-")
		}
	}
}
