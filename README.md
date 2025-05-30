# 🎸 Guitar Tab Library for Go
[![Tests](https://github.com/er-davo/guitar/actions/workflows/go.yaml/badge.svg)](https://github.com/er-davo/guitar/actions/workflows/go.yaml)

A Go library for generating guitar tabs, working with chords, notes, and tunings.

## Installation
```bash
go get github.com/er-davo/guitar
```
## Features
- Tab Generation: Build ASCII tabs from notes/chords.
- Tuning Support: Standard, Drop D, and custom tunings.
- Advanced Techniques: Slides (5/7), hammer-ons (2h4), pull-offs(5p3). Harmonics (<12>).
- Note Calculations: Find closest fret positions, handle enharmonics (e.g., Gb → F#).

# Quick start
## 1. Generate tab
```go
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
	tab.WriteNotes(guitar.GetChord(guitar.Am)...)

	// Add a slide from fret 5 to 7 on the G string
	tab.WriteNotes(guitar.Slide{
		FretStart: 5,
		FretEnd: 7,
		String: 2, // G string
		Time: 0.5,
	})

	fmt.Println(tab.Tab())
}
```
Output
```
e|0------
B|1------
G|2--5/7-
D|2------
A|0------
E|-------
```
## 2. Find Notes on the Fretboard
```go
fb, _ := guitar.NewFingerBoard(tuning, 12) // 12-fret board
notes := fb.FindNotes("C#", 3)              // Find all C#3 notes
// Returns [{C# 3 4 4} {C# 3 9 5}]
```
## 3. Parse Custom Chords
```go
chord := guitar.ParseChord("0 2 2 2 0 -") // A major chord
// Returns [0, 2, 2, 2, 0] (frets for each string)
```