package microbio

import (
	"strings"
)

// NucleotideStrand is a slice of Nucleotides
type NucleotideStrand struct {
	buf    []Nucleotide
	offset int
}

// MakeStrandStr is a helper to call MakeStrand and cast string to bytes
func MakeStrandStr(letters string, sugar SugarType) *NucleotideStrand {
	return MakeStrand([]byte(letters), sugar)
}

// MakeStrand builds NucleotideStrand from string of base letters
// TODO: support spaces for ease of test... will mess with capacity?
func MakeStrand(letters []byte, sugar SugarType) *NucleotideStrand {

	strand := new(NucleotideStrand)

	// allocate slice of nucleotides

	strand.buf = make([]Nucleotide, len(letters))

	// loop through letters and create Nucleotide objects for each

	for i := 0; i < len(letters); i++ {

		strand.buf[i] = Nucleotide{letters[i], sugar, nil}

		// if building a DNA strand, will create the complementary
		// Nucleotide objects and pair them

		if sugar == DEOXYRIBOSE {
			complementary := Nucleotide{ComplementaryBase(letters[i], sugar), sugar, nil}
			BondNucleotides(&strand.buf[i], &complementary)
		}
	}

	return strand
}

// Length : number of total nucleotides in strand,
// regardless of position
func (t *NucleotideStrand) Length() int {
	return len(t.buf)
}

// Remaining : number of nucleotides in the strand after current
func (t *NucleotideStrand) Remaining() int {
	return len(t.buf) - t.offset
}

// CurrentNucleotide : return pointer to current nucleotide in strand
func (t *NucleotideStrand) CurrentNucleotide() *Nucleotide {
	if t.Remaining() <= 0 {
		return nil
	}
	return &t.buf[t.offset]
}

// SlideRight : move one nucleotide right along the strand
func (t *NucleotideStrand) SlideRight() {
	if t.offset >= len(t.buf) {
		return
	}
	t.offset++
}

// Codon returns a new NucleotideStrand of consisting of 3 nucleotides
// copied from current strand, starting at current position.
func (t *NucleotideStrand) Codon() *NucleotideStrand {
	if t.Remaining() < 3 {
		return nil
	}

	codon := new(NucleotideStrand)

	// allocate slice of nucleotides

	codon.buf = make([]Nucleotide, 3)

	// loop through letters and create Nucleotide objects for each

	for i := 0; i < 3; i++ {
		codon.buf[i] = t.buf[t.offset+i]
	}
	return codon
}

// Matches : from the current position on strand, will
// check that every nucleotide in pattern strand is equivalent.
// returns false if not a match or if number of nucleotides
//         remaining in strand is less than length of pattern.
func (t *NucleotideStrand) Matches(pattern *NucleotideStrand) bool {
	if t.Remaining() < pattern.Length() {
		return false
	}
	for i := 0; i < pattern.Length(); i++ {
		a := t.buf[t.offset+i]
		b := pattern.buf[i]
		if (a.base != b.base) || (a.sugar != b.sugar) {
			return false
		}
	}
	return true
}

// String() is the standard toString() equivalent
// This will print the nucleotide bases.
// If any of the nucleotides have a pairing, as is the case
// with most DNA strands, this will make 3-line string
// Example DNA:
//   GATACA
//   ||||||
//   CTATGT
//
// Example RNA (no pairs, one line):
//   GUAC
//
// Example Mixed strand. Top DNA, bottom RNA and DNA nucleotides:
//
//   GATACA
//   :::|||
//   CUATGT
func (t *NucleotideStrand) String() string {
	var sb strings.Builder
	var sbLink strings.Builder
	var sbPaired strings.Builder
	numPairs := 0
	sb.Grow(len(t.buf))
	for i := 0; i < len(t.buf); i++ {
		n := &(t.buf)[i]
		sb.WriteByte(n.base)
		if n.paired == nil {
			sbLink.WriteByte(byte(' '))
			sbPaired.WriteByte(byte(' '))
		} else {
			numPairs++
			sbPaired.WriteByte(n.paired.base)
			if n.sugar == RIBOSE || n.paired.sugar == RIBOSE {
				sbLink.WriteByte(byte(':'))
			} else {
				sbLink.WriteByte(byte('|'))
			}
		}
	}
	retval := sb.String()
	if numPairs > 0 {
		retval += "\n" + sbLink.String() + "\n" + sbPaired.String()
	}
	return retval
}
