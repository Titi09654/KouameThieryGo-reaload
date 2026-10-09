package modifCap

import "strings"

func ModifCap(contenu []byte) string {
	mots := strings.Fields(string(contenu))
	for i := 0; i < len(mots); i++ {
		if mots[i] == "(cap)" {
			lettres := []rune(mots[i-1])
			if lettres[0] >= 'a' && lettres[0] <= 'z' {
				lettres[0] = lettres[0] - ('a' - 'A')
				mots[i-1] = string(lettres)
			}
			mots[i] = ""
		}
	}
	return strings.Join(mots, " ")
}
