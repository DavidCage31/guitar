package guitar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteNotes(t *testing.T) {
	tun, _ := GetTuning(StandardTuning, GuitarType)
	tuningNotes := tun.NoteNames()

	testCases := []struct {
		name        string
		notes       [][]NotePositioner
		expectedTab string
		expectError bool
	}{
		{
			name: "single note on E",
			notes: [][]NotePositioner{
				{Note{Name: "E", Octave: 2, Fret: 12, String: 5, Time: 0}},
			},
			expectedTab: "e|---\nB|---\nG|---\nD|---\nA|---\nE|12-\n",
		},
		{
			name: "multiple notes with timing",
			notes: [][]NotePositioner{
				{Note{Name: "E", Fret: 0, String: 5, Time: 0}},
				{Note{Name: "B", Fret: 1, String: 1, Time: 0.2}},
				{Note{Name: "G", Fret: 3, String: 2, Time: 0.4}},
			},
			expectedTab: "e|------\nB|--1---\nG|----3-\nD|------\nA|------\nE|0-----\n",
		},
		{
			name: "invalid note time",
			notes: [][]NotePositioner{
				{Note{Name: "E", Time: 0.5}},
				{Note{Name: "B", Time: 0.3}},
			},
			expectError: true,
		},
		{
			name: "two-digit fret formatting",
			notes: [][]NotePositioner{
				{Note{Name: "E", Fret: 10, String: 5, Time: 0}},
			},
			expectedTab: "e|---\nB|---\nG|---\nD|---\nA|---\nE|10-\n",
		},
		{
			name: "C chord",
			notes: [][]NotePositioner{
				{
					Note{Fret: 3, String: 4},
					Note{Fret: 2, String: 3},
					Note{Fret: 0, String: 2},
					Note{Fret: 1, String: 1},
					Note{Fret: 0, String: 0},
				},
			},
			expectedTab: "e|0-\nB|1-\nG|0-\nD|2-\nA|3-\nE|--\n",
		},
		{
			name: "slide technique",
			notes: [][]NotePositioner{
				{Slide{FretStart: 5, FretEnd: 7, String: 2, Time: 0}},
			},
			expectedTab: "e|----\nB|----\nG|5/7-\nD|----\nA|----\nE|----\n",
		},
		{
			name: "hammer-on technique",
			notes: [][]NotePositioner{
				{HammerOn{FretFrom: 2, FretTo: 4, String: 1, Time: 0}},
			},
			expectedTab: "e|----\nB|2h4-\nG|----\nD|----\nA|----\nE|----\n",
		},
		{
			name: "invalid string position",
			notes: [][]NotePositioner{
				{Note{String: 10, Time: 0}},
			},
			expectError: true,
		},
		{
			name:        "chord",
			notes:       [][]NotePositioner{GetChord(G)},
			expectedTab: "e|1-\nB|0-\nG|0-\nD|0-\nA|2-\nE|3-\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tb, _ := NewTabBuilder(GuitarType, tuningNotes)

			var err error
			for _, n := range tc.notes {
				err = tb.WriteNotes(n...)
				if err != nil {
					break
				}
			}

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedTab, tb.Tab())
		})
	}
}
