package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func modifUp(contenu []byte) string {
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
func modifBin(contenu []byte) string {

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

func modifHex(contenu []byte) string {
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

func modifCap(contenu []byte) string {
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

func modifLow(contenu []byte) string {
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

func main() {

	fmt.Println(os.Args[1])
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contenu))

	hexModifie := []byte(modifHex(contenu))
	binModifie := []byte(modifBin(hexModifie))
	capModifie := []byte(modifCap(binModifie))
	upModifie := []byte(modifUp(capModifie))
	lowModifie := []byte(modifLow(upModifie))

	errW := os.WriteFile(os.Args[2], lowModifie, 0644)
	if errW != nil {
		fmt.Println(errW)
		return
	}
	fmt.Println(os.Args[2])
	contenu, err = os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(lowModifie))
}
