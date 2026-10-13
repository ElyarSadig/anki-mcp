package anki

import (
	"encoding/json"
	"testing"
)

// Optional fields must distinguish omission from explicit clearing or disabling.
func TestAPIPatchEncoding(t *testing.T) {
	var noMemory *MemoryState
	var noDeck *Reference
	for _, test := range []struct {
		name  string
		value any
		want  string
	}{
		{"omitted memory", RawCardPatch{}, `{}`},
		{"cleared memory", RawCardPatch{FsrsMemory: &noMemory}, `{"fsrsMemory":null}`},
		{"supplied memory", RawCardPatch{FsrsMemory: new(&MemoryState{Difficulty: 5, StabilityDays: 10})}, `{"fsrsMemory":{"difficulty":5,"stabilityDays":10}}`},
		{"omitted tags", NotePatch{}, `{}`},
		{"cleared tags", NotePatch{Tags: new([]string{})}, `{"tags":[]}`},
		{"empty CSS", NoteTypePatch{CSS: new("")}, `{"css":""}`},
		{"zero and false settings", PresetPatch{NewPerDay: new(0), ShowTimer: new(false)}, `{"newPerDay":0,"showTimer":false}`},
		{"empty learning steps", PresetPatch{LearnStepsMinutes: new([]float64{})}, `{"learnStepsMinutes":[]}`},
		{"cleared template deck", TemplatePatch{Deck: &noDeck}, `{"deck":null}`},
		{"empty media", MediaSource{Data: new("")}, `{"data":""}`},
		{"zero structure index", StructureEdit{Op: "moveField", Name: "Front", Index: new(0)}, `{"op":"moveField","name":"Front","index":0}`},
		{"empty selection", SelectBrowserParams{CardIDs: new([]string{})}, `{"cardIds":[]}`},
		{"empty browser query", OpenBrowserParams{Query: new("")}, `{"query":""}`},
		{"disabled media sync", SyncCollectionParams{IncludeMedia: new(false)}, `{"includeMedia":false}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			encoded, err := json.Marshal(test.value)
			if err != nil {
				t.Fatal(err)
			}
			if string(encoded) != test.want {
				t.Fatalf("got %s, want %s", encoded, test.want)
			}
		})
	}
}
