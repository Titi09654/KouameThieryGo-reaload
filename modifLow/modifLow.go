package modifLow

import (
	"strconv"
	"strings"
)

func ModifLow(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		for j := 0; j < len(mots); j++ {
			if mots[j] == "(low)" {
				if j == 0 {
					continue
				}
				lettres := []rune(mots[j-1])
				for l := 0; l < len(lettres); l++ {
					if lettres[l] >= 'A' && lettres[l] <= 'Z' {
						lettres[l] = lettres[l] + ('a' - 'A')
					}
				}
				mots[j-1] = string(lettres)
				mots = append(mots[:j], mots[j+1:]...)
				j--
			}
		}
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}

func ModifLow_nombre(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		for j := 0; j < len(mots); j++ {
			if mots[j] == "(low," {
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
						if lettres[k] >= 'A' && lettres[k] <= 'Z' {
							lettres[k] = lettres[k] + ('a' - 'A')
						}
					}
					mots[j-l] = string(lettres)
				}
				mots = append(mots[:j], mots[j+2:]...)
				j--
			}
		}
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
