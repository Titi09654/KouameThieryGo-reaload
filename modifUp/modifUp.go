package modifUp

import (
	"strconv"
	"strings"
)

func ModifUp(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		for j := 0; j < len(mots); j++ {
			if mots[j] == "0" {
				continue
			}
			if mots[j] == "(up)" {
				lettres := []rune(mots[j-1])
				for j := 0; j < len(lettres); j++ {
					if lettres[j] >= 'a' && lettres[j] <= 'z' {
						lettres[j] = lettres[j] - ('a' - 'A')
					}
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

func ModifUp_nombre(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		for j := 0; j < len(mots); j++ {
			if mots[j] == "(up," {
				if j+1 >= len(mots) {
					continue
				}
				motsModif := strings.TrimSuffix(mots[j+1], ")")
				nombre, err := strconv.Atoi(motsModif)
				if err != nil {
					continue
				}
				for l := 0; l <= nombre; l++ {
					if j-l < 0 {
						break
					}
					lettres := []rune(mots[j-l])
					if len(lettres) == 0 {
						continue
					}
					for k := 0; k < len(lettres); k++ {
						if lettres[k] >= 'a' && lettres[k] <= 'z' {
							lettres[k] = lettres[k] - ('a' - 'A')
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
