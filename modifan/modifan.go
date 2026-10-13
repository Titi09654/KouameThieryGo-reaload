package modifan

import "strings"

func Modifan(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")

	exceptions := []string{"hour", "honest", "honor", "heir", "honestly", "honesty", "honors", "honorable", "honours", "heiress", "heirloom"}

	for i, ligne := range lignes {
		mots := strings.Fields(ligne)

		for j := 0; j < len(mots); j++ {
			if j+1 >= len(mots) {
				continue
			}
			if mots[j] == "a" || mots[j] == "A" {
				lettre := mots[j+1][0]

				estExceptions := false

				for _, mot := range exceptions {

					if strings.EqualFold(mots[j+1], mot) {
						estExceptions = true
						break
					}

				}
				if strings.Contains("aeiou", strings.ToLower(string(lettre))) || estExceptions {
					if mots[j] == "a" {
						mots[j] = "an"
					} else {
						mots[j] = "An"
					}
				}
			}

		}
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
