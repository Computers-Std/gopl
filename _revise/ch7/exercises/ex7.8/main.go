package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column width and print table
}

type multiTier struct {
	ts        []*Track
	primary   string
	secondary string
	third     string
}

func (m *multiTier) Len() int { return len(m.ts) }
func (m *multiTier) Less(i, j int) bool {
	key := m.primary
	for k := range 3 {
		switch key {
		case "Title":
			if m.ts[i].Title != m.ts[j].Title {
				return m.ts[i].Title < m.ts[j].Title
			}
		case "Year":
			if m.ts[i].Year != m.ts[j].Year {
				return m.ts[i].Year < m.ts[j].Year
			}
		case "Length":
			if m.ts[i].Length != m.ts[j].Length {
				return m.ts[i].Length < m.ts[j].Length
			}
		}

		switch k {
		case 0:
			key = m.secondary
		case 1:
			key = m.third
		}
	}
	return false
}

func (m *multiTier) Swap(i, j int) { m.ts[i], m.ts[j] = m.ts[j], m.ts[i] }

func changePrimary(m *multiTier, p string) {
	m.primary, m.secondary, m.third = p, m.primary, m.secondary
}

func SetPrimary(m sort.Interface, p string) {
	switch m := m.(type) {
	case *multiTier:
		changePrimary(m, p)
	}
}

func NewMultiTier(ts []*Track, p, s, th string) sort.Interface {
	return &multiTier{
		ts:        ts,
		primary:   p,
		secondary: s,
		third:     th,
	}
}

func main() {
	fmt.Println("\nMultiTier:")
	multi := NewMultiTier(tracks, "Title", "", "")
	sort.Sort(multi)
	printTracks(tracks)

	fmt.Println()
	SetPrimary(multi, "Year")
	sort.Sort(multi)
	printTracks(tracks)

	fmt.Println()
	SetPrimary(multi, "Length")
	sort.Sort(multi)
	printTracks(tracks)

	fmt.Println()
	SetPrimary(multi, "Title")
	sort.Sort(multi)
	printTracks(tracks)
}
