package modifCap

import (
	"strconv"
	"strings"
)

func ModifCap(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		for j := 0; j < len(mots); j++ {
			if mots[j] == "(cap)" {
				if j < 0 {
					mots = append(mots[:j], mots[j+1:]...)
					j--
					continue
				}
				lettres := []rune(mots[j-1])
				if lettres[0] >= 'a' && lettres[0] <= 'z' {
					lettres[0] = lettres[0] - ('a' - 'A')
				}
				mots[j-1] = string(lettres)
				mots = append(mots[:j], mots[j+1:]...)
				j--
			}
			lignes[i] = strings.Join(mots, " ")
		}
	}
	return strings.Join(lignes, "\n")
}
func ModifCap_nombre(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		for j := 0; j < len(mots); j++ {
			if mots[j] == "(cap," {
				if j+1 >= len(mots) {
					continue
				}
				motsModif := strings.TrimSuffix(mots[j+1], ")")
				nombre, err := strconv.Atoi(motsModif)
				if err != nil {
					continue
				}
				for l := 1; l <= nombre; l++ {
					if j-l < 0 {
						break
					}
					lettres := []rune(mots[j-l])
					if len(lettres) == 0 {
						continue
					}
					for k := 0; k < len(lettres); k++ {
						if lettres[0] >= 'a' && lettres[0] <= 'z' {
							lettres[0] = lettres[0] - ('a' - 'A')
						}
					}
					mots[j-l] = string(lettres)
				}
				mots = append(mots[:j], mots[j+2:]...)
				j--
			}
			lignes[i] = strings.Join(mots, " ")
		}
	}
	return strings.Join(lignes, "\n")
}
