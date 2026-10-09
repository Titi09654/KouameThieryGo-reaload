package modifLow

import "strings"

func ModifLow(contenu []byte) string {
	mots := strings.Fields(string(contenu))
	for i := 0; i < len(mots); i++ {
		if i == 0 {
			continue
		}
		if mots[i] == "(low)" {
			lettres := []rune(mots[i-1])
			for j := 0; j < len(lettres); j++ {
				if lettres[j] >= 'A' && lettres[j] <= 'Z' {
					lettres[j] = lettres[j] + ('a' - 'A')
				}
			}
			mots[i-1] = string(lettres)
			mots[i] = ""
		}
	}
	return strings.Join(strings.Fields(strings.Join(mots, " ")), " ")
}
