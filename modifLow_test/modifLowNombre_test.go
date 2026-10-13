package modiflowtest

import (
	"modification/modifLow"
	"testing"
)

func TestModifLowNombre(t *testing.T) {
	entrer := []string{
		"THIS IS A TEST (low, 2)",
	}
	attendu := []string{
		"THIS IS a test",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifLow.ModifLow_nombre([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu, resultat)
		}
	}
}
