package main

import (
	"fmt"
	"modification/modifBin"
	"modification/modifCap"
	"modification/modifHex"
	"modification/modifLow"
	"modification/modifUp"
	"os"
)

func main() {

	fmt.Println(os.Args[1])
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contenu))

	hexModifie := []byte(modifHex.ModifHex(contenu))
	upModifie := []byte(modifUp.ModifUp(hexModifie))
	capModifie := []byte(modifCap.ModifCap(upModifie))
	binModifie := []byte(modifBin.ModifBin(capModifie))
	lowModifie := []byte(modifLow.ModifLow(binModifie))
	cap_NombModifie := []byte(modifCap.ModifCap_nombre(lowModifie))
	low_NombModifie := []byte(modifLow.ModifLow_nombre(cap_NombModifie))
	up_NombModifie := []byte(modifUp.ModifUp_nombre(low_NombModifie))

	errW := os.WriteFile(os.Args[2], up_NombModifie, 0644)
	if errW != nil {
		fmt.Println(errW)
		return
	}
	fmt.Println(os.Args[2])
	contenu, err = os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(up_NombModifie))
}
