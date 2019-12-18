# rna2protein - Translation

This is a program that uses the code you implemented to translate a nucleotide sequence to a protein.

## Build

Inside vscode there's a panel on the bottom with a tab named 'Terminal'. 
 - Select Terminal tab
 - `cd rna2protein`
 - `go build`

There should now be a rna2protein executable in that directory.

## Run

There should also be a file named `lct.fna` that contains a DNA
sequence for the LCT gene.
It was downloaded from [https://www.ncbi.nlm.nih.gov/nuccore/XM_017004088.2](ncbi.nlm.nih.gov).

```
# Windows:
type lct.fna | .\rna2protein.exe

# MacOS or Unix:
cat lct.fna | ./rna2protein
```

## Results

You should see a long sequence of three letter amino-acid chain that starts with
`GLU-LEU-SER-TRP-HIS-VAL-VAL-PHE-ILE-ALA-LEU-LEU-SER-PHE-SER-CYS-TRP-GLY-SER-ASP-TRP-`

Send results to `ctf@bluesand.org` for credit.
