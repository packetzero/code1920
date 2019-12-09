# CRISPR Challenge

[CRISPR](https://en.wikipedia.org/wiki/CRISPR) stands for *clustered regularly interspaced short palindromic repeats* .

These are unique sections typically between 25 and 35 nucleotides long that repeat in a short span.  For some bacteria, they only repeat 5 times, while in others they can repeat over 30 times.

*For this exercise, we don't care about the palindromic aspect.  The 'palindromic' term is a misnomer if you ask me.  Words that are [palindromes](https://en.wikipedia.org/wiki/Palindrome) are exactly mirrored from the center ('madam','hannah'), while in a strand of RNA some of the nucleotides are complementary (U-A, C-G).  In any case, these RNA sections have a type of symmetry that leads to folding and joining of the complementary nucleotides such that there is a hairpin... but not usually in the exact middle.*

## 1. TopN

Look inside crispr_finder.go and you will see `func TopN()`.  There are some **TODO** comments that you have to replace with code to make the function work correctly.  When implemented correctly, all the 'TopN' tests in `crispr_finder_test.go` should pass.  This function is needed for step 2.

```
// TopN : sorts the vals slice in descending order ,then takes
// the top 5 entries and returns them
func TopN(vals []Section) []Section {
	// TODO: use sort.Slice and provide a function to
	// sort vals in descending order of count

	// TODO: only take top 5 items from vals

	return vals
}
```

## 2. FindRepeats

Implement the **TODO** parts of findRepeats function.  When implemented correctly, all the tests in crispr_finder_test.go should pass.

```
func findRepeats(buf []byte, repeatLen int, minCount int) []Section {

	m := map[string]Section{}
	size := len(buf)

	// make a map of all substrings of repeatLen to occurrances

	for left := 0; left < size-repeatLen; left++ {

		// TODO: get string of length repeatLen at current position
		// if already in map, increment
		// else, add to map
	}

	// make a slice of values with min count

	vals := []Section{}
	// TODO: iterate through map and add items with at least
	// minCount occurrences to vals

	// take topN and return
	vals = TopN(vals)
	return vals
}
```

## 3. Finding these repeats in actual genomes

TODO