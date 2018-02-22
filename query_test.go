package main

import "testing"

func TestQueryWord(t *testing.T) {
	word := query("book")

	if len(word.explanations) == 0 {
		t.Error("Query was incorrect, cannot find explanations for book")
	}
}
