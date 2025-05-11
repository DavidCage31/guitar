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

const (
	// A
	//   e|0
	//   B|2
	//   G|2
	//   D|2
	//   A|0
	//   E|-
	A = iota
	// A7
	//   e|0
	//   B|2
	//   G|0
	//   D|2
	//   A|0
	//   E|-
	A7
	// Am
	//   e|0
	//   B|1
	//   G|2
	//   D|2
	//   A|0
	//   E|-
	Am
	// Am7
	//   e|3
	//   B|1
	//   G|2
	//   D|2
	//   A|0
	//   E|-
	Am7
	// Ab
	//   e|1
	//   B|1
	//   G|2
	//   D|3
	//   A|0
	//   E|-
	Ab

	// B
	//   e|1
	//   B|3
	//   G|3
	//   D|3
	//   A|1
	//   E|1
	B
	// B7
	//   e|2
	//   B|0
	//   G|2
	//   D|1
	//   A|2
	//   E|0
	B7
	// Bm
	//   e|1
	//   B|2
	//   G|3
	//   D|3
	//   A|1
	//   E|0
	Bm
	// Bm7-5
	//   e|1
	//   B|3
	//   G|2
	//   D|0
	//   A|2
	//   E|0
	Bm7_5

	// C
	//   e|0
	//   B|1
	//   G|0
	//   D|2
	//   A|3
	//   E|-
	C
	// C7
	//   e|0
	//   B|1
	//   G|3
	//   D|2
	//   A|3
	//   E|-
	C7
	// C6
	//   e|0
	//   B|1
	//   G|2
	//   D|0
	//   A|3
	//   E|-
	C6
	// C#m
	//   e|1
	//   B|2
	//   G|3
	//   D|1
	//   A|0
	//   E|-
	Cshm

	// D
	//   e|2
	//   B|3
	//   G|2
	//   D|0
	//   A|-
	//   E|-
	D
	// D7
	//   e|2
	//   B|1
	//   G|2
	//   D|0
	//   A|-
	//   E|-
	D7
	// Dm
	//   e|1
	//   B|3
	//   G|2
	//   D|0
	//   A|-
	//   E|-
	Dm
	// Dm7
	//   e|1
	//   B|1
	//   G|2
	//   D|0
	//   A|-
	//   E|-
	Dm7
	// Dm7
	//   e|1
	//   B|0
	//   G|2
	//   D|0
	//   A|-
	//   E|-
	Dm6
	// Dsus4
	//   e|3
	//   B|3
	//   G|2
	//   D|0
	//   A|-
	//   E|-
	Dsus4

	// E
	//   e|0
	//   B|0
	//   G|1
	//   D|2
	//   A|2
	//   E|0
	E
	// E7
	//   e|0
	//   B|3
	//   G|1
	//   D|0
	//   A|2
	//   E|0
	E7
	// Em
	//   e|0
	//   B|0
	//   G|0
	//   D|2
	//   A|2
	//   E|0
	Em
	// Esus4
	//   e|0
	//   B|0
	//   G|2
	//   D|2
	//   A|2
	//   E|0
	Esus4
	// E7 jazz
	//   e|0
	//   B|2
	//   G|3
	//   D|1
	//   A|0
	//   E|0
	E7Jazz

	// F
	//   e|1
	//   B|1
	//   G|2
	//   D|3
	//   A|3
	//   E|1
	F
	// F7
	//   e|1
	//   B|1
	//   G|2
	//   D|1
	//   A|3
	//   E|1
	F7
	// Fm
	//   e|1
	//   B|1
	//   G|1
	//   D|3
	//   A|3
	//   E|1
	Fm
	// Fmaj7
	//   e|0
	//   B|1
	//   G|2
	//   D|3
	//   A|0
	//   E|-
	Fmaj7
	// F#m7-5
	//   e|2
	//   B|2
	//   G|2
	//   D|1
	//   A|0
	//   E|0
	Fshm7_5
	// F#
	//   e|1
	//   B|0
	//   G|2
	//   D|3
	//   A|3
	//   E|1
	Fsh

	// G
	//   e|1
	//   B|0
	//   G|0
	//   D|0
	//   A|2
	//   E|3
	G
	// G7
	//   e|1
	//   B|0
	//   G|0
	//   D|0
	//   A|2
	//   E|3
	G7
	// Gm
	//   e|1
	//   B|1
	//   G|1
	//   D|3
	//   A|3
	//   E|1
	Gm
	// Gm6
	//   e|0
	//   B|3
	//   G|3
	//   D|0
	//   A|0
	//   E|3
	Gm6
	// G9
	//   e|1
	//   B|3
	//   G|3
	//   D|0
	//   A|0
	//   E|3
	G9
	// Gmaj7
	//   e|2
	//   B|0
	//   G|0
	//   D|0
	//   A|0
	//   E|3
	Gmaj7

	// C5
	//   e|-
	//   B|-
	//   G|-
	//   D|5
	//   A|3
	//   E|-
	PowerC5
	// D5
	//   e|-
	//   B|-
	//   G|-
	//   D|7
	//   A|5
	//   E|-
	PowerD5
	// Db5
	//   e|-
	//   B|-
	//   G|-
	//   D|6
	//   A|4
	//   E|-
	PowerDb5
	// Eb5
	//   e|-
	//   B|-
	//   G|-
	//   D|3
	//   A|1
	//   E|-
	PowerEb5
	// E5
	//   e|-
	//   B|-
	//   G|-
	//   D|-
	//   A|2
	//   E|0
	PowerE5
	// F5
	//   e|-
	//   B|-
	//   G|-
	//   D|-
	//   A|3
	//   E|1
	PowerF5
	// Gb5
	//   e|-
	//   B|-
	//   G|-
	//   D|-
	//   A|4
	//   E|2
	PowerGb5
	// G5
	//   e|-
	//   B|-
	//   G|-
	//   D|-
	//   A|5
	//   E|3
	PowerG5
	// A5
	//   e|-
	//   B|-
	//   G|-
	//   D|-
	//   A|7
	//   E|5
	PowerA5
	// Bb5
	//   e|-
	//   B|-
	//   G|-
	//   D|3
	//   A|1
	//   E|-
	PowerBb5
	// B5
	//   e|-
	//   B|-
	//   G|-
	//   D|4
	//   A|2
	//   E|-
	PowerB5
)

func GetChord(chordNum int) []NotePositioner {
	switch chordNum {
	case A:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 2, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 0, String: 4},
		}
	case A7:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 2, String: 1},
			Note{Fret: 0, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 0, String: 4},
		}
	case Am:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 0, String: 4},
		}
	case Am7:
		return []NotePositioner{
			Note{Fret: 3, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 0, String: 4},
		}
	case Ab:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 3, String: 3},
			Note{Fret: 0, String: 4},
		}
	case B:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 3, String: 1},
			Note{Fret: 3, String: 2},
			Note{Fret: 3, String: 3},
			Note{Fret: 1, String: 4},
			Note{Fret: 1, String: 5},
		}
	case B7:
		return []NotePositioner{
			Note{Fret: 2, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 1, String: 3},
			Note{Fret: 2, String: 4},
			Note{Fret: 0, String: 5},
		}
	case Bm:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 2, String: 1},
			Note{Fret: 3, String: 2},
			Note{Fret: 3, String: 3},
			Note{Fret: 1, String: 4},
			Note{Fret: 0, String: 5},
		}
	case Bm7_5:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 3, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 0, String: 3},
			Note{Fret: 2, String: 4},
			Note{Fret: 0, String: 5},
		}
	case C:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 0, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 3, String: 4},
		}
	case C7:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 3, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 3, String: 4},
		}
	case C6:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 0, String: 3},
			Note{Fret: 3, String: 4},
		}
	case Cshm:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 2, String: 1},
			Note{Fret: 3, String: 2},
			Note{Fret: 1, String: 3},
			Note{Fret: 0, String: 4},
		}
	case D:
		return []NotePositioner{
			Note{Fret: 2, String: 0},
			Note{Fret: 3, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 0, String: 3},
		}
	case D7:
		return []NotePositioner{
			Note{Fret: 2, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 0, String: 3},
		}
	case Dm:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 3, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 0, String: 3},
		}
	case Dm7:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 0, String: 3},
		}
	case Dm6:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 0, String: 3},
		}
	case Dsus4:
		return []NotePositioner{
			Note{Fret: 3, String: 0},
			Note{Fret: 3, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 0, String: 3},
		}
	case E:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 1, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 2, String: 4},
			Note{Fret: 0, String: 5},
		}
	case E7:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 3, String: 1},
			Note{Fret: 1, String: 2},
			Note{Fret: 0, String: 3},
			Note{Fret: 2, String: 4},
			Note{Fret: 0, String: 5},
		}
	case Em:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 0, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 2, String: 4},
			Note{Fret: 0, String: 5},
		}
	case Esus4:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 2, String: 3},
			Note{Fret: 2, String: 4},
			Note{Fret: 0, String: 5},
		}
	case E7Jazz:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 2, String: 1},
			Note{Fret: 3, String: 2},
			Note{Fret: 1, String: 3},
			Note{Fret: 0, String: 4},
			Note{Fret: 0, String: 5},
		}
	case F:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 3, String: 3},
			Note{Fret: 3, String: 4},
			Note{Fret: 1, String: 5},
		}
	case F7:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 1, String: 3},
			Note{Fret: 3, String: 4},
			Note{Fret: 1, String: 5},
		}
	case Fm:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 1, String: 2},
			Note{Fret: 3, String: 3},
			Note{Fret: 3, String: 4},
			Note{Fret: 1, String: 5},
		}
	case Fmaj7:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 3, String: 3},
			Note{Fret: 0, String: 4},
		}
	case Fshm7_5:
		return []NotePositioner{
			Note{Fret: 2, String: 0},
			Note{Fret: 2, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 1, String: 3},
			Note{Fret: 0, String: 4},
			Note{Fret: 0, String: 5},
		}
	case Fsh:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 2, String: 2},
			Note{Fret: 3, String: 3},
			Note{Fret: 3, String: 4},
			Note{Fret: 1, String: 5},
		}
	case G:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 0, String: 2},
			Note{Fret: 0, String: 3},
			Note{Fret: 2, String: 4},
			Note{Fret: 3, String: 5},
		}
	case G7:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 0, String: 2},
			Note{Fret: 0, String: 3},
			Note{Fret: 2, String: 4},
			Note{Fret: 3, String: 5},
		}
	case Gm:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 1, String: 1},
			Note{Fret: 1, String: 2},
			Note{Fret: 3, String: 3},
			Note{Fret: 3, String: 4},
			Note{Fret: 1, String: 5},
		}
	case Gm6:
		return []NotePositioner{
			Note{Fret: 0, String: 0},
			Note{Fret: 3, String: 1},
			Note{Fret: 3, String: 2},
			Note{Fret: 0, String: 3},
			Note{Fret: 0, String: 4},
			Note{Fret: 3, String: 5},
		}
	case G9:
		return []NotePositioner{
			Note{Fret: 1, String: 0},
			Note{Fret: 3, String: 1},
			Note{Fret: 3, String: 2},
			Note{Fret: 0, String: 3},
			Note{Fret: 0, String: 4},
			Note{Fret: 3, String: 5},
		}
	case Gmaj7:
		return []NotePositioner{
			Note{Fret: 2, String: 0},
			Note{Fret: 0, String: 1},
			Note{Fret: 0, String: 2},
			Note{Fret: 0, String: 3},
			Note{Fret: 0, String: 4},
			Note{Fret: 3, String: 5},
		}
	case PowerC5:
		return []NotePositioner{
			Note{Fret: 5, String: 3},
			Note{Fret: 3, String: 4},
		}
	case PowerD5:
		return []NotePositioner{
			Note{Fret: 7, String: 3},
			Note{Fret: 5, String: 4},
		}
	case PowerDb5:
		return []NotePositioner{
			Note{Fret: 6, String: 3},
			Note{Fret: 4, String: 4},
		}
	case PowerEb5:
		return []NotePositioner{
			Note{Fret: 3, String: 3},
			Note{Fret: 1, String: 4},
		}
	case PowerE5:
		return []NotePositioner{
			Note{Fret: 2, String: 4},
			Note{Fret: 0, String: 5},
		}
	case PowerF5:
		return []NotePositioner{
			Note{Fret: 3, String: 4},
			Note{Fret: 1, String: 5},
		}
	case PowerGb5:
		return []NotePositioner{
			Note{Fret: 4, String: 4},
			Note{Fret: 2, String: 5},
		}
	case PowerG5:
		return []NotePositioner{
			Note{Fret: 5, String: 4},
			Note{Fret: 3, String: 5},
		}
	case PowerA5:
		return []NotePositioner{
			Note{Fret: 7, String: 4},
			Note{Fret: 5, String: 5},
		}
	case PowerBb5:
		return []NotePositioner{
			Note{Fret: 3, String: 3},
			Note{Fret: 1, String: 4},
		}
	case PowerB5:
		return []NotePositioner{
			Note{Fret: 4, String: 3},
			Note{Fret: 2, String: 4},
		}
	default:
		return []NotePositioner{}
	}
}
