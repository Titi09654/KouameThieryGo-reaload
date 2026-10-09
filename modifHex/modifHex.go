package modifHex

import (
	"fmt"
	"strconv"
	"strings"
)

func ModifHex(contenu []byte) string {
	mots := strings.Fields(string(contenu))
	for i := 0; i < len(mots); i++ {
		if mots[i] == "(hex)" {
			nombre, err := strconv.ParseInt(mots[i-1], 16, 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			mots[i] = ""
			mots[i-1] = strconv.FormatInt(nombre, 10)

		}

	}
	return strings.Join(mots, " ")
}
