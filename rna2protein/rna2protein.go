package main

import (
	"fmt"
	"os"

	"github.com/packetzero/code1920/microbio"
)

func main() {
	nucleotides := microbio.LoadNucleotidesFromFastaFile(os.Stdin, microbio.RIBOSE)
	//fmt.Println(nucleotides)
	ribosome := microbio.Ribosome{}
	protein := ribosome.Translate(nucleotides)
	fmt.Println(protein.String())
}
