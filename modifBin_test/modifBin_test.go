package modifbintest

import (
	"modification/modifBin"
	"testing"
)

func TestModifBin(t *testing.T) {

	entrer := []string{
		"1010 (bin)",
		"11000011010100010 (bin)",
		"1101111111110101111 (bin)",
	}

	attendu := []string{
		"10",
		"100002",
		"458671",
	}

	for i := 0; i < len(entrer); i++ {
		resultat := modifBin.ModifBin([]byte(entrer[i]))

		if resultat != attendu[i] {

			t.Errorf("attendu %q, obtenu %q", attendu[i], resultat)
		}
	}

}
