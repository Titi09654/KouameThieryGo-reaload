package modifponctuation

import (
	"strings"
)

func Modifponctuation(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		for j := 0; j < len(mots); j++ {
			if mots[j] == "," || mots[j] == "..." || mots[j] == "?" || mots[j] == "." || mots[j] == "!" {
				if j > 0 {
					lettre := mots[j-1] + mots[j]
					mots[j-1] = lettre
					mots = append(mots[:j], mots[j+1:]...)
					j--
				}
			}
		}
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
