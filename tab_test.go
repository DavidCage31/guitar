package guitar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteNotes(t *testing.T) {
	tun, _ := ParseTuning(StandardTuning)
	tuningNotes := tun.NoteNames()

	testCases := []struct {
		name        string
		notes       [][]Playable
		expectedTab string
		expectError bool
	}{
		{
			name: "single note on E",
			notes: [][]Playable{
				{Note{Name: "E", Octave: 2, Fret: 12, String: 5, Time: 0}},
			},
			expectedTab: "e|---\nB|---\nG|---\nD|---\nA|---\nE|12-\n",
		},
		{
			name: "multiple notes with timing",
			notes: [][]Playable{
				{Note{Name: "E", Fret: 0, String: 5, Time: 0}},
				{Note{Name: "B", Fret: 1, String: 1, Time: 0.2}},
				{Note{Name: "G", Fret: 3, String: 2, Time: 0.4}},
			},
			expectedTab: "e|------\nB|--1---\nG|----3-\nD|------\nA|------\nE|0-----\n",
		},
		{
			name: "invalid note time",
			notes: [][]Playable{
				{Note{Name: "E", Time: 0.5}},
				{Note{Name: "B", Time: 0.3}},
			},
			expectError: true,
		},
		{
			name: "two-digit fret formatting",
			notes: [][]Playable{
				{Note{Name: "E", Fret: 10, String: 5, Time: 0}},
			},
			expectedTab: "e|---\nB|---\nG|---\nD|---\nA|---\nE|10-\n",
		},
		{
			name: "C chord",
			notes: [][]Playable{
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
			notes: [][]Playable{
				{Slide{FretStart: 5, FretEnd: 7, String: 2, Time: 0}},
			},
			expectedTab: "e|----\nB|----\nG|5/7-\nD|----\nA|----\nE|----\n",
		},
		{
			name: "hammer-on technique",
			notes: [][]Playable{
				{HammerOn{FretFrom: 2, FretTo: 4, String: 1, Time: 0}},
			},
			expectedTab: "e|----\nB|2h4-\nG|----\nD|----\nA|----\nE|----\n",
		},
		{
			name: "invalid string position",
			notes: [][]Playable{
				{Note{String: 10, Time: 0}},
			},
			expectError: true,
		},
		{
			name:        "chord",
			notes:       [][]Playable{ParseChord("1 0 0 0 2 3", 0)},
			expectedTab: "e|1-\nB|0-\nG|0-\nD|0-\nA|2-\nE|3-\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tb, _ := NewTabWriter(tuningNotes)

			var err error
			for _, n := range tc.notes {
				err = tb.WriteNotes(n)
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

func TestWriteNote(t *testing.T) {
	tun, _ := ParseTuning(StandardTuning)
	tuningNotes := tun.NoteNames()

	testCases := []struct {
		name        string
		notes       []Playable
		expectedTab string
		expectError bool
	}{
		{
			name: "single note on E",
			notes: []Playable{
				Note{Name: "E", Octave: 2, Fret: 12, String: 5, Time: 0},
			},
			expectedTab: "e|---\nB|---\nG|---\nD|---\nA|---\nE|12-\n",
		},
		{
			name: "sequence of notes with timing",
			notes: []Playable{
				Note{Name: "E", Fret: 0, String: 5, Time: 0},
				Note{Name: "B", Fret: 1, String: 1, Time: 0.2},
				Note{Name: "G", Fret: 3, String: 2, Time: 0.4},
			},
			expectedTab: "e|------\nB|--1---\nG|----3-\nD|------\nA|------\nE|0-----\n",
		},
		{
			name: "invalid note time",
			notes: []Playable{
				Note{Name: "E", Time: 0.5},
				Note{Name: "B", Time: 0.3},
			},
			expectError: true,
		},
		{
			name: "two-digit fret formatting",
			notes: []Playable{
				Note{Name: "E", Fret: 10, String: 5, Time: 0},
			},
			expectedTab: "e|---\nB|---\nG|---\nD|---\nA|---\nE|10-\n",
		},
		{
			name: "slide technique",
			notes: []Playable{
				Slide{FretStart: 5, FretEnd: 7, String: 2, Time: 0},
			},
			expectedTab: "e|----\nB|----\nG|5/7-\nD|----\nA|----\nE|----\n",
		},
		{
			name: "hammer-on technique",
			notes: []Playable{
				HammerOn{FretFrom: 2, FretTo: 4, String: 1, Time: 0},
			},
			expectedTab: "e|----\nB|2h4-\nG|----\nD|----\nA|----\nE|----\n",
		},
		{
			name: "invalid string position",
			notes: []Playable{
				Note{String: 10, Time: 0},
			},
			expectError: true,
		},
		{
			name: "single note from chord",
			notes: []Playable{
				ParseChord("1 0 0 0 2 3", 0)[0],
			},
			expectedTab: "e|1-\nB|--\nG|--\nD|--\nA|--\nE|--\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tb, _ := NewTabWriter(tuningNotes)

			var err error
			for _, n := range tc.notes {
				err = tb.WriteNote(n)
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
