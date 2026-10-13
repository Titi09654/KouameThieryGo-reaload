package modiflowtest

import (
	"modification/modifLow"
	"testing"
)

func TestModifLow(t *testing.T) {

	entrer := []string{
		"HELLO WORLD (low)",
	}
	attendu := []string{
		"HELLO world",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifLow.ModifLow([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu, resultat)
		}
	}
}
