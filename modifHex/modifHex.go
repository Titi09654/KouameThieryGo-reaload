package modifHex

import (
	"fmt"
	"strconv"
	"strings"
)

func ModifHex(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")

	for i, ligne := range lignes {
		mots := strings.Fields(ligne)

		for j := 0; j < len(mots); j++ {
			if mots[j] == "(hex)" {
				if j > 0 {
					nombre, err := strconv.ParseInt(mots[j-1], 16, 64)
					if err != nil {
						fmt.Println(err)
						continue
					}
					mots[j-1] = strconv.FormatInt(nombre, 10)
				}
				mots = append(mots[:j], mots[j+1:]...)
				j--
			}
		}
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
