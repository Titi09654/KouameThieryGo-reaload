package modifBin

import (
	"fmt"
	"strconv"
	"strings"
)

func ModifBin(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		for j := 0; j < len(mots); j++ {
			if mots[j] == "(bin)" {
				nombre, err := strconv.ParseInt(mots[j-1], 2, 64)
				if err != nil {
					fmt.Print(err)
					continue
				}
				mots[j] = " "
				mots[j-1] = strconv.FormatInt(nombre, 10)
			}
		}
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}
