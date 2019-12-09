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

// codonToAminoMap maps 3-nucleotide codon strings to the corresponding
// amino acid or stop signal(nil).
// TODO: The amino acids are defined for you in amino_acid.go file.
// It's important that you follow the UCAG pattern as demonstrated below,
// as the Unit Test depends on exact order.
// One line for each of the second letter U,C,A,G and each line
// has a mapping where the third letter changes from U,C,A,G.
// https://www.thoughtco.com/genetic-code-373449

var codonToAminoMap = map[string]*AminoAcid{
	// UUU through UGG
	"UUU": &Phenylalanine, "UUC": &Phenylalanine, "UUA": &Leucine, "UUG": &Leucine,
	"UCU": &Serine, "UCC": &Serine, "UCA": &Serine, "UCG": &Serine,
	"UAU": &Tyrosine, "UAC": &Tyrosine, "UAA": nil, "UAG": nil,
	"UGU": &Cysteine, "UGC": &Cysteine, "UGA": nil, "UGG": &Tryptophan,

	// CUU through CGG
	// TODO: Add 16 entries

	// AUU through AGG
	// TODO: Add 16 entries

	// GUU through GGG
	// TODO: Add 16 entries
}
