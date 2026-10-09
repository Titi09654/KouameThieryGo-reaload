package modifUp

import "strings"

func ModifUp(contenu []byte) string {
	mots := strings.Fields(string(contenu))
	for i := 0; i < len(mots); i++ {
		if mots[i] == "0" {
			continue
		}
		if mots[i] == "(up)" {
			lettres := []rune(mots[i-1])
			for j := 0; j < len(lettres); j++ {
				if lettres[j] >= 'a' && lettres[j] <= 'z' {
					lettres[j] = lettres[j] - ('a' - 'A')
				}
			}
			mots[i] = ""
			mots[i-1] = string(lettres)
		}
	}
	return strings.Join(strings.Fields(strings.Join(mots, " ")), " ")
}
