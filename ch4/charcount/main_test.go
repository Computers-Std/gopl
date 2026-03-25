package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCountData(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantCounts  map[rune]int
		wantUtflen  [utf8.UTFMax + 1]int
		wantInvalid int
	}{
		{
			name:        "Simple ASCII",
			input:       "abc",
			wantCounts:  map[rune]int{'a': 1, 'b': 1, 'c': 1},
			wantUtflen:  [5]int{0, 3, 0, 0, 0},
			wantInvalid: 0,
		},
		{
			name:        "Multi-byte Unicode",
			input:       "Go⌘", // '⌘' is 3 bytes
			wantCounts:  map[rune]int{'G': 1, 'o': 1, '⌘': 1},
			wantUtflen:  [5]int{0, 2, 0, 1, 0},
			wantInvalid: 0,
		},
		{
			name:        "Invalid UTF-8",
			input:       "abc" + string([]byte{0xff}), // 0xff is an invalid start byte
			wantCounts:  map[rune]int{'a': 1, 'b': 1, 'c': 1},
			wantUtflen:  [5]int{0, 3, 0, 0, 0},
			wantInvalid: 1,
		},
		{
			name:        "Empty Input",
			input:       "",
			wantCounts:  map[rune]int{},
			wantUtflen:  [5]int{0, 0, 0, 0, 0},
			wantInvalid: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			counts, utflen, invalid, err := countData(r)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(counts, tt.wantCounts) {
				t.Errorf("counts = %v, want %v", counts, tt.wantCounts)
			}
			if utflen != tt.wantUtflen {
				t.Errorf("utflen = %v, want %v", utflen, tt.wantUtflen)
			}
			if invalid != tt.wantInvalid {
				t.Errorf("invalid = %d, want %d", invalid, tt.wantInvalid)
			}
		})
	}
}
