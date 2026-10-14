package modifponctuation


import (
	"strings"
	"unicode"
)

func Modifponctuation(contenu []byte) string {
	lignes := strings.Split(string(contenu), "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		var resultat []string

		for _, mot := range mots {
			if mot == "," || mot == "..." || mot == "?" || mot == "." || mot == "!" || mot == ";" || mot == ":" {
				if len(resultat) > 0 {
					resultat[len(resultat)-1] += mot
				} else {
					resultat = append(resultat, mot)
				}
			} else {
				resultat = append(resultat, mot)
			}
		}

		lignes[i] = strings.Join(resultat, " ")
	}

	texte := strings.Join(lignes, "\n")
	runes := []rune(texte)

	type Delimiteur struct {
		position int
		symbole  rune
	}

	var pile []Delimiteur
	espacesASupprimer := make([]bool, len(runes))

	for i, c := range runes {
		if c == '\'' && i > 0 && i+1 < len(runes) &&
    (unicode.IsLetter(runes[i-1]) || unicode.IsDigit(runes[i-1])) {

    j := i + 1
    for j < len(runes) && runes[j] == ' ' {
        espacesASupprimer[j] = true
        j++
    }

    continue
}
		if c == '(' {
			pile = append(pile, Delimiteur{i, c})
			continue
		}
		if c == ')' {
			if len(pile) == 0 {
				continue
			}
			dernier := pile[len(pile)-1]
			if dernier.symbole != '(' {
				continue
			}
			for j := dernier.position + 1; j < i &&
				runes[j] == ' '; j++ {
				espacesASupprimer[j] = true
			}
			for j := i - 1; j > dernier.position &&
				runes[j] == ' '; j-- {
				espacesASupprimer[j] = true
			}
			pile = pile[:len(pile)-1]
			continue
		}
		if c == '"' || c == '\'' {
			if len(pile) > 0 &&
				pile[len(pile)-1].symbole == c {

				dernier := pile[len(pile)-1]

				for j := dernier.position + 1; j < i &&
					runes[j] == ' '; j++ {
					espacesASupprimer[j] = true
				}

				for j := i - 1; j > dernier.position &&
					runes[j] == ' '; j-- {
					espacesASupprimer[j] = true
				}

				pile = pile[:len(pile)-1]
			} else {
				pile = append(pile, Delimiteur{i, c})
			}
		}
	}
	var sortie strings.Builder

	for i, c := range runes {
		if !espacesASupprimer[i] {
			sortie.WriteRune(c)
		}
	}

	return sortie.String()
}


