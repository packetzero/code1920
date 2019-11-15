package microbio

type Ribosome struct {
}

var startCodon = MakeStrand([]byte("AUG"), RIBOSE)

// Translate returns an AminoAcidStrand for a given NucleotideStrand
// for each codon after start , call WaitForTRNAAndGetAttachedAminoAcid()
// to get amino acid and add to amino strand
func (t *Ribosome) Translate(nucleotides *NucleotideStrand) *AminoAcidStrand {
	protein := AminoAcidStrand{}

	// TODO

	return &protein
}
