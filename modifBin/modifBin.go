package modifBin

import (
	"fmt"
	"strconv"
	"strings"
)

func ModifBin(contenu []byte) string {

	mots := strings.Fields(string(contenu))
	for i := 0; i < len(mots); i++ {
		if mots[i] == "(bin)" {
			nombre, err := strconv.ParseInt(mots[i-1], 2, 64)
			if err != nil {
				fmt.Print(err)
				continue
			}
			mots[i] = " "
			mots[i-1] = strconv.FormatInt(nombre, 10)
		}
	}
	return strings.Join(mots, " ")
}
