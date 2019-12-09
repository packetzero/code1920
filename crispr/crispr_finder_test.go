package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopN(t *testing.T) {
	vals := []Section{
		{"aaa", 3},
		{"bbb", 1},
		{"ccc", 10},
		{"d", 1},
		{"e", 7},
		{"f", 2},
		{"g", 3},
		{"h", 1},
	}
	results := TopN(vals)

	// should always return 5 or less
	assert.Equal(t, 5, len(results))

	// results should be in descending order
	assert.Equal(t, "ccc", results[0].str)
	assert.Equal(t, 10, results[0].count)
	assert.Equal(t, "e", results[1].str)
	assert.Equal(t, 7, results[1].count)
}

// test topN with less than 5 items
func TestTopNsmall(t *testing.T) {
	vals := []Section{
		{"aaa", 3},
		{"bbb", 1},
		{"ccc", 10},
	}
	results := TopN(vals)
	assert.Equal(t, 3, len(results))
	assert.Equal(t, "ccc", results[0].str)
	assert.Equal(t, 10, results[0].count)
	assert.Equal(t, "aaa", results[1].str)
	assert.Equal(t, 3, results[1].count)
}

// test TopN with no items
func TestTopNempty(t *testing.T) {
	vals := []Section{}
	results := TopN(vals)
	assert.Equal(t, 0, len(results))
}

// If all values are same, order shouldn't change?
func TestTopNSame(t *testing.T) {
	vals := []Section{
		{"aaa", 1},
		{"bbb", 1},
		{"ccc", 1},
		{"d", 1},
		{"e", 1},
		{"f", 1},
		{"g", 1},
		{"h", 1},
	}
	results := TopN(vals)
	assert.Equal(t, 5, len(results))
	assert.Equal(t, "aaa", results[0].str)
	assert.Equal(t, "bbb", results[1].str)
}

var testDataBytes = []byte("GAATACTGTAGCCAGACCAGAATAGCCAGCAACAGCAGCGTTAGTTTACCCATCCTGCCCCCTGAAAAAC" +
	"GAATACTGCATCCCATGCATCCGAAGACGACTCTACATCCTCTGTTGGGGATACCGCGACAACGCGGGCA" +
	"GAATACTGCTATTTGTCCATTGTTACGTATACCCAGGGCGTGCAGAACATAATCTCATTATTAGTTACGG" +
	"GAATACTGATGAACAGAGGAGACAAGAAAGTACAAATTAGCCCAGTAGCCACATAAACAGTGCGCCAAAC" +
	"GAATACTGACATTTGTCCATTGTTACGTATACCCAGGGCGTGCAGAACATAATCTCATTATTAGTTACGG" +
	"GAATACTGATGAACAGAGGAGACAAGAAAGTACAAATTAGCCCAGTAGCCACATAAACAGTGCGCCAAAC" +
	"GAATACTGTACTGTCATCAGGGTGAAAACAATACTGTAGCGTAGCTTTCCGTCCATCAATGAATGCAGCG" +
	"GAATACTGACCACTACTGCGACAGGCATCAGTGCCAGAAAGAAAGGCCAGGTGTAGATAAAGAAGAACAG")

func TestFindRepeatsSimple(t *testing.T) {
	data := testDataBytes
	// find length 8, min repeat of 4
	results := findRepeats(data, 8, 4)
	//t.Log(results)
	assert.Equal(t, 2, len(results))
	if len(results) > 1 {
		assert.Equal(t, "GAATACTG", results[0].str)
		assert.Equal(t, 8, results[0].count)
		assert.Equal(t, "AATACTGA", results[1].str)
		assert.Equal(t, 4, results[1].count)
	}

	// find length 9, min 4 occurrences
	results = findRepeats(data, 9, 4)
	assert.Equal(t, 1, len(results))
	if len(results) > 0 {
		assert.Equal(t, "GAATACTGA", results[0].str)
		assert.Equal(t, 4, results[0].count)
	}
}
