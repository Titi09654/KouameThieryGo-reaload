package modifhextest

import (
	"modification/modifHex"
	"testing"
)

func TestModifCap(t *testing.T) {

	entrer := []string{
		"Simply add 42 (hex)",
	}

	attendu := []string{
		"Simply add 66",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifHex.ModifHex([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu, resultat)
		}
	}
}
