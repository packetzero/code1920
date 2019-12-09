package microbio

// Ribosome : Here we attempt to model the Translation phase
// of gene expression done by the Ribosome
type Ribosome struct {
}

var startCodon = MakeStrand([]byte("AUG"), RIBOSE)

// Translate returns an AminoAcidStrand for a given NucleotideStrand
// Use WaitForTRNAAndGetAttachedAminoAcid(codon) to get AminoAcid
func (t *Ribosome) Translate(nucleotides NucleotideStrand) AminoAcidStrand {

	// TODO: implement
	// First, loop through nucleotides looking for START (matching startCodon)
	// then adding amino acids to an AminoAcidStrand until STOP.
	// A STOP is when WaitForTRNAAndGetAttachedAminoAcid() returns nil.
	// If reach the end of the NucleotideStrand without a STOP,
	// then consider that an error and return an empty AminoAcidStrand.
	return AminoAcidStrand{}
}
