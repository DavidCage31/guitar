package guitar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseChord(t *testing.T) {
	testCases := []struct {
		name        string
		chordTab    string
		expected    []NotePositioner
		expectError bool
	}{
		{
			name:     "Standard A chord",
			chordTab: "0 2 2 2 0 -",
			expected: []NotePositioner{
				Note{Fret: 0, String: 0},
				Note{Fret: 2, String: 1},
				Note{Fret: 2, String: 2},
				Note{Fret: 2, String: 3},
				Note{Fret: 0, String: 4},
			},
		},
		{
			name:     "Power chord E5",
			chordTab: "- - - 2 0 -",
			expected: []NotePositioner{
				Note{Fret: 2, String: 3},
				Note{Fret: 0, String: 4},
			},
		},
		{
			name:     "Chord with skipped strings (mixed)",
			chordTab: "3 - 0 1 - 3",
			expected: []NotePositioner{
				Note{Fret: 3, String: 0},
				Note{Fret: 0, String: 2},
				Note{Fret: 1, String: 3},
				Note{Fret: 3, String: 5},
			},
		},
		{
			name:        "Invalid character (letters)",
			chordTab:    "a b c 1 2 3",
			expectError: true,
		},
		{
			name:        "Empty input",
			chordTab:    "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ParseChord(tc.chordTab, 0)

			if tc.expectError {
				assert.Empty(t, result, "Expected error but got notes")
				return
			}

			assert.Equal(t, len(tc.expected), len(result), "Number of parsed notes mismatch")
			for i, note := range tc.expected {
				assert.Equal(t, note.TabSymbol(), result[i].TabSymbol(), "Fret mismatch")
				assert.Equal(t, note.StringNumber(), result[i].StringNumber(), "String mismatch")
			}
		})
	}
}
