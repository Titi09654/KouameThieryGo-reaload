package modifuptest

import (
	"modification/modifUp"
	"testing"
)

func TestModifUp(t *testing.T) {
	entrer := []string{
		"it was the worst of times (up)",
	}
	attendu := []string{
		"it was the worst of TIMES",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifUp.ModifUp([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu%q", attendu, resultat)
		}
	}
}
