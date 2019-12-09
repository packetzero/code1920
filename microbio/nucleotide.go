package microbio

// https://en.m.wikipedia.org/wiki/Nucleoside
// DNA: GTAC cytosine [C], guanine [G], adenine [A] or thymine [T]),
// RNA: GUAC guanine, uracil, adenine, and cytosine

const (
	// ADENINE and other bases
	ADENINE  = byte('A')
	GUANINE  = byte('G')
	CYTOSINE = byte('C')
	THYMINE  = byte('T')
	URACIL   = byte('U')
)

// SugarType of Nucleotides
type SugarType int

// enum of SugarTypes
const (
	RIBOSE      SugarType = 0
	DEOXYRIBOSE SugarType = 1
)

// Nucleotide is a molecule in the double helix DNA
// or single-stranded RNA.  In our simplified representation,
// we have the base, sugar, and an optional paired or bonded
// nucleotide.  The sugar is the main indicator whether the nucleotide
// is a RNA (RIBOSE) or DNA (DEOXYRIBOSE).  The other indicator is
// that RNA uses URACIL, while DNA has THYMINE base.
// We don't model the phosphate or hydroxyl group here.
type Nucleotide struct {
	base   byte
	sugar  SugarType
	paired *Nucleotide
}

// complementary base map
var compBaseDna = map[byte]byte{'A': 'T', 'T': 'A', 'G': 'C', 'C': 'G'}
var compBaseRna = map[byte]byte{'A': 'U', 'U': 'A', 'G': 'C', 'C': 'G'}

// ComplementaryBase returns base letter of matching nucleotide
func ComplementaryBase(base byte, sugar SugarType) byte {
	if sugar == DEOXYRIBOSE {
		return compBaseDna[base]
	}
	return compBaseRna[base]
}

// BondNucleotides sets paired members to each other
func BondNucleotides(a *Nucleotide, b *Nucleotide) {
	a.paired = b
	b.paired = a
}

// BreakNucleotideBond sets paired to nil for both nucleotides
func BreakNucleotideBond(a *Nucleotide) {
	if a.paired == nil {
		return
	}
	a.paired.paired = nil
	a.paired = nil
}
