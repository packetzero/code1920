package microbio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var RNABASES = []string{"U", "C", "A", "G"}

const expectedAminos = "PHE,PHE,LEU,LEU," +
	"SER,SER,SER,SER," +
	"TYR,TYR,nil,nil," +
	"CYS,CYS,nil,TRP," +
	// C row
	"LEU,LEU,LEU,LEU," +
	"PRO,PRO,PRO,PRO," +
	"HIS,HIS,GLN,GLN," +
	"ARG,ARG,ARG,ARG," +
	// A row
	"ILE,ILE,ILE,MET," +
	"THR,THR,THR,THR," +
	"ASN,ASN,LYS,LYS," +
	"SER,SER,ARG,ARG," +
	// G row
	"VAL,VAL,VAL,VAL," +
	"ALA,ALA,ALA,ALA," +
	"ASP,ASP,GLU,GLU," +
	"GLY,GLY,GLY,GLY,"

func TestLookupAminoForCodon(t *testing.T) {
	str := ""
	for _, a := range RNABASES {
		for _, b := range RNABASES {
			for _, c := range RNABASES {
				// codonStr will be "UAG" , etc.
				codonStr := a + b + c
				pAmino := lookupAminoForCodon(codonStr)
				if pAmino == nil {
					str += "nil,"
				} else {
					str += pAmino.id + ","
				}
			}
		}
	}
	assert.Equal(t, expectedAminos, str)
}
