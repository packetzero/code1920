# Molecular Biology Challenge

Molecular Biology is full of functions and machines.  I have created some classes that abstract nucleotides, amino acids, nucleotide strands, and have implemented the unit tests you need to validate your code.

## 1. Finish Genetic Code mapping

Open up `genetic_code.go` and see there are 3 TODO comments there.  You have to finish up the mapping for codons to amino acids defined in `amino_acids.go`.  The table is defined at [thoughtco page](https://www.thoughtco.com/genetic-code-373449) among other places on the web.  Follow the UCAG ordering from that page.  When you have this correct, running the unit tests in `genetic_code_test.go` will pass.  This should be simple and afterwards, you should get familiar with some of the names.  Notice that there are 4^3 = 64 different combinations for the RNA nucleotide codons, but only 20 amino acids.  Some of them map to a STOP code (which we represent as `nil`).

## 2. Implement Ribosome.Translate()

The ribosome can process a strand of RNA and build proteins (specific amino acid strands).  The process is called [Translation](https://en.wikipedia.org/wiki/Translation_(biology)).  Your job is to implement this in `ribosome.go`.  When you get it right, the tests in `ribosome_tests.go` will pass.
