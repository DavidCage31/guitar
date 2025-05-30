package guitar

import (
	"strconv"
	"strings"
)

func ParseChord(chordTab string) []NotePositioner {
	notes := strings.Split(chordTab, " ")
	chord := []NotePositioner{}

	for i, note := range notes {
		num, err := strconv.Atoi(note)
		if err != nil {
			if note != "-" {
				return []NotePositioner{}
			}
			continue
		}
		chord = append(chord, Note{Fret: num, String: i})
	}

	return chord
}
