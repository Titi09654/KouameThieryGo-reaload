package modifcaptest

import (
	"modification/modifCap"
	"testing"
)

func TestModifCapNombre(t *testing.T) {

	entrer := []string{
		"it was the best of times (cap, 3)",
		"it is A hair (cap, 4)",
	}

	attendu := []string{
		"it was the Best Of Times",
		"It Is A Hair",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifCap.ModifCap_nombre([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu, resultat)
		}
	}
}
