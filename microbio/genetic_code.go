package microbio

// WaitForTRNAAndGetAttachedAminoAcid emulates tRNA machine
// RNA codon to deliver matching amino acids.
// tRNA have an RNA matching codon (3 nucleotides) on one end
// and an AminoAcid on the other
func WaitForTRNAAndGetAttachedAminoAcid(codon *NucleotideStrand) *AminoAcid {
	return lookupAminoForCodon(codon.String())
}

func lookupAminoForCodon(codonStr string) *AminoAcid {
	return codonToAminoMap[codonStr]
}

// https://www.thoughtco.com/genetic-code-373449

var codonToAminoMap = map[string]*AminoAcid{
	// UUU through UGG
	"UUU": &Phenylalanine, "UUC": &Phenylalanine, "UUA": &Leucine, "UUG": &Leucine,
	"UCU": &Serine, "UCC": &Serine, "UCA": &Serine, "UCG": &Serine,
	"UAU": &Tyrosine, "UAC": &Tyrosine, "UAA": nil, "UAG": nil,
	"UGU": &Cysteine, "UGC": &Cysteine, "UGA": nil, "UGG": &Tryptophan,

	// CUU through CGG
	// TODO : amino acids are defined in amino_acid.go

	// AUU through AGG
	// TODO

	// GUU through GGG
	// TODO
}
