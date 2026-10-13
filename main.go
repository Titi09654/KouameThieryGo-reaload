package main

import (
	"fmt"
	"modification/modifBin"
	"modification/modifCap"
	"modification/modifHex"
	"modification/modifLow"
	"modification/modifUp"
	"modification/modifan"
	"modification/modifponctuation"
	"os"
)

func main() {

	fmt.Println(os.Args[1])
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	hexModifie := []byte(modifHex.ModifHex(contenu))
	upModifie := []byte(modifUp.ModifUp(hexModifie))
	capModifie := []byte(modifCap.ModifCap(upModifie))
	binModifie := []byte(modifBin.ModifBin(capModifie))
	lowModifie := []byte(modifLow.ModifLow(binModifie))
	cap_NombModifie := []byte(modifCap.ModifCap_nombre(lowModifie))
	low_NombModifie := []byte(modifLow.ModifLow_nombre(cap_NombModifie))
	up_NombModifie := []byte(modifUp.ModifUp_nombre(low_NombModifie))
	a_Modifie := []byte(modifan.Modifan(up_NombModifie))
	ponctuaction := []byte(modifponctuation.Modifponctuation(a_Modifie))

	errW := os.WriteFile(os.Args[2], ponctuaction, 0644)
	if errW != nil {
		fmt.Println(errW)
		return
	}
	fmt.Println(os.Args[2])
	contenu, err = os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
}
