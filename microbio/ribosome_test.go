package microbio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	ribosome := Ribosome{}
	rna := MakeStrand([]byte("UACGGGAUGUCACCUACGGUAUAACGCGGG"), RIBOSE)
	protein := ribosome.Translate(rna)
	assert.Equal(t, "SER-PRO-THR-VAL", protein.String())
}

func TestTranslateEmpty(t *testing.T) {
	ribosome := Ribosome{}
	rna := MakeStrand([]byte(""), RIBOSE)
	protein := ribosome.Translate(rna)
	assert.Equal(t, "", protein.String())
}

// Make sure Translate doesn't crash when a full 3-base codon
// is not remaining in strand
func TestTranslateInvalidLength(t *testing.T) {
	ribosome := Ribosome{}
	rna := MakeStrand([]byte("UACGGGAUGUCACCUAC"), RIBOSE)
	protein := ribosome.Translate(rna)
	assert.Equal(t, 0, len(protein))
	assert.Equal(t, "", protein.String())

	rna = MakeStrand([]byte("UACGGGAUGUCACCUA"), RIBOSE)
	protein = ribosome.Translate(rna)
	assert.Equal(t, 0, len(protein))
	assert.Equal(t, "", protein.String())
}

// No 'AUG' codon, so Translate should return empty
func TestTranslateNoStart(t *testing.T) {
	ribosome := Ribosome{}
	rna := MakeStrand([]byte("AUAUAUAAAAGGUA"), RIBOSE)
	protein := ribosome.Translate(rna)
	assert.Equal(t, 0, len(protein))
	assert.Equal(t, "", protein.String())
}

// See what happens when no STOP codon found
func TestTranslateNoEnd(t *testing.T) {
	ribosome := Ribosome{}
	rna := MakeStrand([]byte("UACGGGAUG"+"UCACCUACGGUA"+"CGCGGG"), RIBOSE)
	protein := ribosome.Translate(rna)
	assert.Equal(t, 0, len(protein))
	assert.Equal(t, "", protein.String())
}
