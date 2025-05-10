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
	for _, n := range notes {
		if n.StartTime() < tb.time {
			return fmt.Errorf("note time %v precedes current time %v", n.StartTime(), tb.time)
		}
		if n.StartTime() != time {
			return fmt.Errorf("notes time are not equal")
		}
	}

	silence := int((time - tb.time) / tb.timeStep)
	tb.addSilence(silence)
	tb.time += tb.timeStep * (float32(silence + 1))

	maxLen := -1

	for _, n := range notes {
		tb.tabStrings[n.StringPosition()].WriteString(n.FretPosition())
		if tb.tabStrings[n.StringPosition()].Len() > maxLen {
			maxLen = tb.tabStrings[n.StringPosition()].Len()
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
