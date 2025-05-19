package guitar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTuning(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      Tuning
		expectedError string
	}{
		{
			name:  "Standard 6-string tuning",
			input: "E4 B3 G3 D3 A2 E2",
			expected: Tuning{
				{Name: "E", Octave: 4, String: 0, Fret: 0},
				{Name: "B", Octave: 3, String: 1, Fret: 0},
				{Name: "G", Octave: 3, String: 2, Fret: 0},
				{Name: "D", Octave: 3, String: 3, Fret: 0},
				{Name: "A", Octave: 2, String: 4, Fret: 0},
				{Name: "E", Octave: 2, String: 5, Fret: 0},
			},
		},
		{
			name:  "Tuning with sharps",
			input: "F#4 C#4 G#3 D#3 A#2 F#2",
			expected: Tuning{
				{Name: "F#", Octave: 4, String: 0, Fret: 0},
				{Name: "C#", Octave: 4, String: 1, Fret: 0},
				{Name: "G#", Octave: 3, String: 2, Fret: 0},
				{Name: "D#", Octave: 3, String: 3, Fret: 0},
				{Name: "A#", Octave: 2, String: 4, Fret: 0},
				{Name: "F#", Octave: 2, String: 5, Fret: 0},
			},
		},
		{
			name:  "Tuning with flats",
			input: "Gb4 Db4 Ab3 Eb3 Bb2 Gb2",
			expected: Tuning{
				{Name: "F#", Octave: 4, String: 0, Fret: 0}, // Нормализовано в F#
				{Name: "C#", Octave: 4, String: 1, Fret: 0}, // Нормализовано в C#
				{Name: "G#", Octave: 3, String: 2, Fret: 0},
				{Name: "D#", Octave: 3, String: 3, Fret: 0},
				{Name: "A#", Octave: 2, String: 4, Fret: 0},
				{Name: "F#", Octave: 2, String: 5, Fret: 0},
			},
		},
		{
			name:  "4-string bass tuning",
			input: "G2 D2 A1 E1",
			expected: Tuning{
				{Name: "G", Octave: 2, String: 0, Fret: 0},
				{Name: "D", Octave: 2, String: 1, Fret: 0},
				{Name: "A", Octave: 1, String: 2, Fret: 0},
				{Name: "E", Octave: 1, String: 3, Fret: 0},
			},
		},
		{
			name:          "Empty input",
			input:         "",
			expectedError: "empty notes",
		},
		{
			name:          "Invalid octave (letter)",
			input:         "EX B3 G3 D3 A2 E2",
			expectedError: "invalid octave at note: EX",
		},
		{
			name:          "Invalid note name",
			input:         "H4 B3 G3 D3 A2 E2",
			expectedError: "invalid note name: H",
		},
		{
			name:          "Missing octave",
			input:         "E B3 G3 D3 A2 E2",
			expectedError: "invalid octave at note: E",
		},
		{
			name:          "Double sharps/flats not supported",
			input:         "F##4 B3 G3 D3 A2 E2",
			expectedError: "invalid note name: F##",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseTuning(tc.input)

			if tc.expectedError != "" {
				assert.ErrorContains(t, err, tc.expectedError)
				assert.Equal(t, Tuning{}, result)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, len(tc.expected), len(result), "Number of strings mismatch")

			for i := range tc.expected {
				assert.Equal(t, tc.expected[i].Name, result[i].Name, "Note name mismatch at string %d", i)
				assert.Equal(t, tc.expected[i].Octave, result[i].Octave, "Octave mismatch at string %d", i)
				assert.Equal(t, tc.expected[i].String, result[i].String, "String number mismatch at string %d", i)
				assert.Equal(t, 0, result[i].Fret, "Fret should be 0 for open string")
			}
		})
	}
}
