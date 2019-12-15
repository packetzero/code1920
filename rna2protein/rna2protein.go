package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/packetzero/code1920/microbio"
)

func main() {

	prefix := []byte{}
	pPrependAUGFlag := flag.Bool("prependAUG", false, "If true, prepand AUG start sequence to nucleotide strand")
	pVerboseFlag := flag.Bool("verbose", false, "If true, prints out more debugging information")

	flag.Parse()

	if *pPrependAUGFlag {
		prefix = []byte("AUG")
	}

	nucleotides := microbio.LoadNucleotidesFromFastaFile(os.Stdin, microbio.RIBOSE, prefix)
	if *pVerboseFlag {
		fmt.Println(nucleotides.String())
	}
	ribosome := microbio.Ribosome{}

	protein := ribosome.Translate(nucleotides)
	fmt.Println(protein.String())
}
