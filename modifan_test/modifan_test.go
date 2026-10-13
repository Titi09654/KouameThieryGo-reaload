package modifan_test

import (
	"modification/modifan"
	"testing"
)

func TestModifan(t *testing.T) {

	entrer := []string{
		"a apple",
		"a village",
		"a house",
		"a orange",
	}

	attendu := []string{
		"an apple",
		"a village",
		"a house",
		"an orange",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifan.Modifan([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu[i], resultat)
		}
	}

}
