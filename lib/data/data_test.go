package data

import "testing"

func TestDicts(t *testing.T) {
	for _, dict := range []map[string]bool{OneKSet(), TenKSet(), LargeSet()} {
		if !dict["go"] {
			t.Errorf("Expected %q in dictionary, but was not found", "go")
		}
	}
}

func TestBigram(t *testing.T) {
	want := 4356462 + 37938534 + 5403137 + 4402940 + 144814
	if score := BigramScore("bitlux"); score != want {
		t.Errorf("BigramScore(%q) = %d, wanted %d", "bitlux", score, want)
	}
}
