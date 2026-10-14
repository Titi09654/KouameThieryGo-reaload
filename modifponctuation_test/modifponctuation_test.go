package modifponctuationtest

import (
	"modification/modifponctuation"
	"testing"
)

func TestModifPonctuaction(t *testing.T) {
	entrer := []string{
		"( bonjour )",
		`" bonjour "`,
		"' bonjour '",
		"l' homme et j' aime",
		"( bonjour , toi ! )",
	}
	attendu := []string{
	  "(bonjour)",
	  `"bonjour"`,
	  "'bonjour'",
	  "l'homme et j'aime",
	  "(bonjour, toi!)",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifponctuation.Modifponctuation([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu, resultat)
		}
	}
}
