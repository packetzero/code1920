package microbio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeStrandDna(t *testing.T) {
	s := []byte("GATCCTAG")

	b := *MakeStrand(s, DEOXYRIBOSE)
	assert.Equal(t, len(s), b.Length())
	assert.Equal(t, GUANINE, b.buf[0].base)
	assert.Equal(t, ADENINE, b.buf[1].base)
	assert.Equal(t, THYMINE, b.buf[2].base)
	assert.Equal(t, CYTOSINE, b.buf[3].base)
	assert.Equal(t, CYTOSINE, b.buf[4].base)
	assert.Equal(t, THYMINE, b.buf[5].base)
	assert.Equal(t, ADENINE, b.buf[6].base)
	assert.Equal(t, GUANINE, b.buf[7].base)

	for i := 0; i < b.Length(); i++ {
		assert.Equal(t, DEOXYRIBOSE, b.buf[i].sugar)
		// should be paired
		assert.True(t, b.buf[i].paired != nil)
	}

}

func TestMakeStrandRna(t *testing.T) {
	s := []byte("GAUCCUAG")

	b := *MakeStrand(s, RIBOSE)
	assert.Equal(t, len(s), b.Length())
	assert.Equal(t, GUANINE, b.buf[0].base)
	assert.Equal(t, ADENINE, b.buf[1].base)
	assert.Equal(t, URACIL, b.buf[2].base)
	assert.Equal(t, CYTOSINE, b.buf[3].base)
	assert.Equal(t, byte('C'), b.buf[4].base)
	assert.Equal(t, byte('U'), b.buf[5].base)
	assert.Equal(t, byte('A'), b.buf[6].base)
	assert.Equal(t, byte('G'), b.buf[7].base)

	for i := 0; i < b.Length(); i++ {
		assert.Equal(t, RIBOSE, b.buf[i].sugar)
		// should not be paired
		assert.True(t, b.buf[i].paired == nil)
	}
}

func TestMatch(t *testing.T) {
	strand := MakeStrandStr("GUACA", RIBOSE)
	uac := MakeStrandStr("UAC", RIBOSE)

	assert.False(t, strand.Matches(uac))
	strand.SlideRight()
	assert.True(t, strand.Matches(uac))
	strand.SlideRight()
	assert.False(t, strand.Matches(uac))
}

func TestPrintRna(t *testing.T) {
	s := []byte("GAUCCUAG")

	b := *MakeStrand(s, RIBOSE)
	assert.Equal(t, "GAUCCUAG", b.String())
}

func TestPrintDna(t *testing.T) {
	s := []byte("GATCCTAG")

	b := *MakeStrand(s, DEOXYRIBOSE)
	assert.Equal(t, "GATCCTAG\n||||||||\nCTAGGATC", b.String())
}

func TestPrintDnaRnaMix(t *testing.T) {
	s := []byte("GATCCTAG")

	b := *MakeStrand(s, DEOXYRIBOSE)

	BreakNucleotideBond(&b.buf[0])
	BreakNucleotideBond(&b.buf[1])

	a := Nucleotide{'C', RIBOSE, nil}
	u := Nucleotide{'U', RIBOSE, nil}

	BondNucleotides(&b.buf[0], &a)
	BondNucleotides(&b.buf[1], &u)

	assert.Equal(t, "GATCCTAG\n::||||||\nCUAGGATC", b.String())
}
