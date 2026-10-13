package modifponctuationtest

import (
	"modification/modifponctuation"
	"testing"
)

func TestModifPonctuaction(t *testing.T) {
	entrer := []string{
		"hello my FRIEND  !",
		"I love call of dutty  .",
		"Hello , world ! How are you ?",
	}
	attendu := []string{
		"hello my FRIEND!",
		"I love call of dutty.",
		"Hello, world! How are you?",
	}
	for i := 0; i < len(entrer); i++ {
		resultat := modifponctuation.Modifponctuation([]byte(entrer[i]))
		if resultat != attendu[i] {
			t.Errorf("attendu %q, obtenu %q", attendu, resultat)
		}
	}
}
