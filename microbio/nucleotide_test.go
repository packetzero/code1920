package microbio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBondNucleotides(t *testing.T) {
	a := Nucleotide{'A', DEOXYRIBOSE, nil}
	b := Nucleotide{'T', DEOXYRIBOSE, nil}
	BondNucleotides(&a, &b)
	assert.Equal(t, &b, a.paired)
	assert.Equal(t, &a, b.paired)
}

func TestBreakNucleotideBond(t *testing.T) {
	a := Nucleotide{'A', DEOXYRIBOSE, nil}
	b := Nucleotide{'T', DEOXYRIBOSE, nil}
	BondNucleotides(&a, &b)
	BreakNucleotideBond(&a)
	assert.True(t, nil == a.paired)
}
