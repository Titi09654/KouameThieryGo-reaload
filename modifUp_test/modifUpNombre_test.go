package modifuptest

import (
	"modification/modifUp"
	"testing"
)

func TestModifUpNombre(t *testing.T) {
	entrer := []string{
		"We are learning go (up, 4)",
		"I love call of dutty (up, 5)",
	}
	attendu := []string{
		"WE ARE LEARNING GO",
		"I LOVE CALL OF DUTTY",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifUp.ModifUp_nombre([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu, resultat)
		}
	}
}
