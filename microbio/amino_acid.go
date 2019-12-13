package microbio

import "strings"

// AminoAcid is simply a 3-letter id and name for our purposes
type AminoAcid struct {
	id   string
	name string
}

// AminoAcidStrand is an array of AminoAcid pointers.
// Also named polypeptide and with the right combination a protein
type AminoAcidStrand []*AminoAcid

func (t *AminoAcidStrand) String() string {
	var sb strings.Builder
	for i := 0; i < len(*t); i++ {
		if i > 0 {
			sb.WriteString("-")
		}
		aa := (*t)[i]
		sb.WriteString(aa.id)
	}
	return sb.String()
}

// Define the 20 proteins used in GeneticCode
// https://www.thoughtco.com/genetic-code-373449

var Alanine = AminoAcid{"ALA", "Alanine"}
var Arginine = AminoAcid{"ARG", "Arginine"}
var Asparagine = AminoAcid{"ASN", "Asparagine"}
var AsparticAcid = AminoAcid{"ASP", "Aspartic Acid"}

var Cysteine = AminoAcid{"CYS", "Cysteine"}
var GlutamicAcid = AminoAcid{"GLU", "Glutamic Acid"}
var Glutamine = AminoAcid{"GLN", "Glutamine"}
var Glycine = AminoAcid{"GLY", "Glycine"}

var Histidine = AminoAcid{"HIS", "Histidine"}
var Isoleucine = AminoAcid{"ILE", "Isoleucine"}
var Leucine = AminoAcid{"LEU", "Leucine"}
var Lysine = AminoAcid{"LYS", "Lysine"}

var Methionine = AminoAcid{"MET", "Methionine"}
var Phenylalanine = AminoAcid{"PHE", "Phenylalanine"}
var Proline = AminoAcid{"PRO", "Proline"}
var Serine = AminoAcid{"SER", "Serine"}

var Threonine = AminoAcid{"THR", "Threonine"}
var Tryptophan = AminoAcid{"TRP", "Tryptophan"}
var Tyrosine = AminoAcid{"TYR", "Tyrosine"}
var Valine = AminoAcid{"VAL", "Valine"}
