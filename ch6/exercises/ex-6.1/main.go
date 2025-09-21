package main

// An IntSet is a set of small non-negative integers
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Len returns the number of elements in the set (count of set bits).
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for word != 0 {
			// Increment count if the least significant bit is 1
			if word&1 == 1 {
				count++
			}
			// Shift word right by 1 to check the next bit
			word >>= 1
		}
	}
	return count
}

// Remove removes an element from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64

	// Check if the index is valid (i.e., within range)
	if word < len(s.words) {
		// Clear the corresponding bit (set it to 0)
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() // remove all elements from the set

func (s *IntSet) Copy() *IntSet // return a copy of the set

// NOTE: Lacking the domain knowledge
