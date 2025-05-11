package guitar

import (
	"errors"
)

type FingerBoard struct {
	tuning Tuning
	frets  int
}

func NewFingerBoard(tun Tuning, frets int) (*FingerBoard, error) {
	if frets < 0 {
		return nil, errors.New("frets value can not be negative")
	}

	return &FingerBoard{
		tuning: tun,
		frets:  frets,
	}, nil
}

func (fb *FingerBoard) GetTuningNotes() []string {
	return fb.tuning.NoteNames()
}

// TODO
// fix error
// returns empty list for # notes
func (fb FingerBoard) GetNotes(targetNote string, targetOctave int) Notes {
	notes := Notes{}
	currentNote := Note{}

	for i := range fb.tuning {
		currentNote = fb.tuning[i]

		for fret := 0; fret < fb.frets; fret++ {
			if currentNote.Name == targetNote && currentNote.Octave == targetOctave {
				notes = append(notes, currentNote)
			}
			currentNote.AddFret()
		}
	}

	return notes
}
