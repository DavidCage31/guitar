package guitar

import (
	"errors"
	"fmt"
	"strings"
)

type TuningType int

const (
	StandardTuning TuningType = iota
)

type Tuning []Note

func (t *Tuning) NoteNames() []string {
	names := make([]string, len(*t))
	for i := range *t {
		names[i] = (*t)[i].Name
	}
	names[0] = strings.ToLower(names[0])
	return names
}

func NewTuning(t []Note, instrument InstrumentType) (Tuning, error) {
	for i := range t {
		if !noteIsValid(t[i]) {
			return Tuning{}, fmt.Errorf("invalid note %+v", t[i])
		}
	}

	switch instrument {
	case GuitarType:
		if len(t) != 6 {
			return Tuning{}, fmt.Errorf("invalid tuning for guitar! must be 6 notes")
		}
	default:
		return Tuning{}, fmt.Errorf("unsupported instrument")
	}

	return append(Tuning{}, t...), nil
}

func GetTuning(t TuningType, instrumentType InstrumentType) (Tuning, error) {
	switch instrumentType {
	case GuitarType:
		switch t {
		case StandardTuning:
			return Tuning{
				{
					Name:   "E",
					Octave: 4,
					Fret:   0,
					String: 0,
				},
				{
					Name:   "B",
					Octave: 3,
					Fret:   0,
					String: 1,
				},
				{
					Name:   "G",
					Octave: 3,
					Fret:   0,
					String: 2,
				},
				{
					Name:   "D",
					Octave: 3,
					Fret:   0,
					String: 3,
				},
				{
					Name:   "A",
					Octave: 2,
					Fret:   0,
					String: 4,
				},
				{
					Name:   "E",
					Octave: 2,
					Fret:   0,
					String: 5,
				},
			}, nil
		default:
			return Tuning{}, errors.ErrUnsupported
		}
	default:
		return Tuning{}, errors.ErrUnsupported
	}
}
