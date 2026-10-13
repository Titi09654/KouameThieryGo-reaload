package modifcaptest

import (
	"modification/modifCap"
	"testing"
)

func TestModifCap(t *testing.T) {

	entrer := []string{
		"it (cap) was the best of times",
		"i (cap)",
		"it (cap) is A hair",
	}

	attendu := []string{
		"It was the best of times",
		"I",
		"It is A hair",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifCap.ModifCap([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu[i], resultat)
		}
	}

}
