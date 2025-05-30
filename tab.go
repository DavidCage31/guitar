package guitar

import (
	"fmt"
	"sort"
	"strings"
)

type TabWriter struct {
	time     float32
	timeStep float32

	tabStrings []strings.Builder
}

type Playable interface {
	TabSymbol() string
	StringNumber() int
	StartTime() float32
}

func NewTabWriter(tuningNotes []string, opts ...TabOption) (*TabWriter, error) {
	tb := &TabWriter{
		time:       0,
		timeStep:   0.2,
		tabStrings: make([]strings.Builder, len(tuningNotes)),
	}

	for _, opt := range opts {
		opt(tb)
	}

	if err := tb.addNotes(tuningNotes); err != nil {
		return nil, err
	}

	return tb, nil
}

func (tb *TabWriter) Tab() string {
	tab := strings.Builder{}

	for i := range len(tb.tabStrings) {
		tab.WriteString(tb.tabStrings[i].String() + "\n")
	}

	return tab.String()
}

func (tb *TabWriter) WriteNotes(notes ...Playable) error {

	sort.Slice(notes, func(i, j int) bool { return notes[i].StartTime() < notes[j].StartTime() })

	time := notes[0].StartTime()
	if time < tb.time {
		return fmt.Errorf("note time %v precedes current time %v",
			time, tb.time)
	}

	for _, n := range notes {
		if n.StringNumber() >= len(tb.tabStrings) {
			return fmt.Errorf("invalid string index %d, in tab builder only %d strings",
				n.StringNumber(), len(tb.tabStrings))
		}
	}

	for i := 0; i < len(notes); i++ {
		time = notes[i].StartTime()
		maxLen := -1
		minLen := tb.tabStrings[0].Len()

		silence := int((time - tb.time) / tb.timeStep)
		tb.addSilence(silence)
		tb.time += tb.timeStep * (float32(silence + 1))

		for i < len(notes) && time == notes[i].StartTime() {
			stringPos := notes[i].StringNumber()

			if tb.tabStrings[stringPos].Len() != minLen {
				return fmt.Errorf("can not write more 2 or more notes with equal time: %f to 1 string", time)
			}

			tb.tabStrings[stringPos].WriteString(notes[i].TabSymbol())
			if tb.tabStrings[stringPos].Len() > maxLen {
				maxLen = tb.tabStrings[stringPos].Len()
			}
			i++
		}

		for i := range tb.tabStrings {
			if tb.tabStrings[i].Len() < maxLen {
				diffLen := maxLen - tb.tabStrings[i].Len()
				for j := 0; j < diffLen; j++ {
					tb.tabStrings[i].WriteString("-")
				}
			}
		}
	}

	// to escape situations like:
	// E|-3--123-----
	tb.addSilence(1)

	return nil
}

func (tb *TabWriter) addNotes(notes []string) error {
	if len(notes) != len(tb.tabStrings) {
		return fmt.Errorf("invalid tuning notes count")
	}

	for i := range notes {
		tb.tabStrings[i].WriteString(notes[i] + "|")
	}
	return nil
}

func (tb *TabWriter) addSilence(n int) {
	for range n {
		for i := range len(tb.tabStrings) {
			tb.tabStrings[i].WriteString("-")
		}
	}
}

type TabOption func(*TabWriter)

func WithTimeStep(step float32) TabOption {
	return func(tb *TabWriter) {
		tb.timeStep = step
	}
}

func WithDefaultTimeStep() TabOption {
	return func(tb *TabWriter) {
		tb.timeStep = 0.2
	}
}
