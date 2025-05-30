package main

import (
	"fmt"

	"github.com/er-davo/guitar"
)

func main() {
	// Create a tab builder in Standard Tuning
	tuning, _ := guitar.ParseTuning(guitar.StandardTuning)
	tab, _ := guitar.NewTabWriter(tuning.NoteNames())

	// Add an A minor chord
	tab.WriteNotes(guitar.ParseChord("0 1 2 2 0 -", 0)...)

	// Add a slide from fret 5 to 7 on the G string
	tab.WriteNotes(guitar.Slide{FretStart: 5, FretEnd: 7, String: 2, Time: 0.5})

	fmt.Println(tab.Tab())

	fb, _ := guitar.NewFingerBoard(tuning, 12) // 12-fret board
	notes := fb.FindNotes("C#", 3)             // Find all C#3 notes
	// Returns [{C# 3 4 4} {C# 3 9 5}]
	fmt.Println(notes)
}
