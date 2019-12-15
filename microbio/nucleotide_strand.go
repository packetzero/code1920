package microbio

import (
	"bufio"
	"io"
	"strings"
)

// NucleotideStrand is a slice of Nucleotides with a
// position offset.  The position changes with calls to SlideRight()
type NucleotideStrand struct {
	buf      []Nucleotide
	position int
}

// MakeStrandStr is a helper to call MakeStrand and cast string to bytes
func MakeStrandStr(letters string, sugar SugarType) NucleotideStrand {
	return MakeStrand([]byte(letters), sugar)
}

// MakeStrand builds NucleotideStrand from string of base letters
func MakeStrand(letters []byte, sugar SugarType) NucleotideStrand {

	strand := NucleotideStrand{}

	// allocate slice of nucleotides

	strand.buf = make([]Nucleotide, len(letters))

	// loop through letters and create Nucleotide objects for each

	for i := 0; i < len(letters); i++ {

		c := letters[i]

		// To make it easier to deal with FASTA files
		// which all seem to have DNA nucleotides:
		// map T to U and vice-versa depending on sugar

		if sugar == RIBOSE && c == THYMINE {
			c = URACIL
		} else if sugar == DEOXYRIBOSE && c == URACIL {
			c = THYMINE
		}

		strand.buf[i] = Nucleotide{c, sugar, nil}

		// if building a DNA strand, will create the complementary
		// Nucleotide objects and pair them

		if sugar == DEOXYRIBOSE {
			complementary := Nucleotide{ComplementaryBase(letters[i], sugar), sugar, nil}
			BondNucleotides(&strand.buf[i], &complementary)
		}
	}

	return strand
}

// LoadNucleotidesFromFastaFile will load entire contents
// of infile to a byte buffer, call MakeStrand and return
// resulting strand.
// prefix should normally be empty, but can contain a sequence
// of nucleotides to prepend to the beginning of the strand
// read from file.
func LoadNucleotidesFromFastaFile(infile io.Reader, sugar SugarType, prefix []byte) NucleotideStrand {
	buf := []byte{}
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		line := scanner.Bytes()

		// skip FASTA header line
		if line[0] == byte('>') {
			continue
		}

		// NOTE: MakeStrand() will convert T to U depending on sugar

		buf = append(buf, line...)
	}

	if len(prefix) > 0 {
		buf = append(prefix, buf...)
	}

	return MakeStrand(buf, sugar)
}

// Length : number of total nucleotides in strand,
// regardless of position
func (t *NucleotideStrand) Length() int {
	return len(t.buf)
}

// Remaining : number of nucleotides in the strand after current
func (t *NucleotideStrand) Remaining() int {
	return len(t.buf) - t.position
}

// CurrentNucleotide : return pointer to current nucleotide in strand
func (t *NucleotideStrand) CurrentNucleotide() *Nucleotide {
	if t.Remaining() <= 0 {
		return nil
	}
	return &t.buf[t.position]
}

// SlideRight : move one nucleotide right along the strand
func (t *NucleotideStrand) SlideRight() {
	if t.position >= len(t.buf) {
		return
	}
	t.position++
}

// Codon returns a new NucleotideStrand of consisting of 3 nucleotides
// copied from current strand, starting at current position.
func (t *NucleotideStrand) Codon() *NucleotideStrand {
	if t.Remaining() < 3 {
		return nil
	}

	codon := NucleotideStrand{}

	// allocate slice of nucleotides

	codon.buf = make([]Nucleotide, 3)

	// loop through letters and create Nucleotide objects for each

	for i := 0; i < 3; i++ {
		codon.buf[i] = t.buf[t.position+i]
	}
	return &codon
}

// Matches : from the current position on strand, will
// check that every nucleotide in pattern strand matches
// corresponding nucleotides in this strand.
// returns false if not a match or if number of nucleotides
//         remaining in strand is less than length of pattern.
func (t *NucleotideStrand) Matches(pattern NucleotideStrand) bool {
	if t.Remaining() < pattern.Length() {
		return false
	}
	for i := 0; i < pattern.Length(); i++ {
		a := t.buf[t.position+i]
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
