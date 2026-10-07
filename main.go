package main

import (
	"fmt"
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


	errW := os.WriteFile(os.Args[2], contenu, 0644)
	if errW != nil {
		fmt.Println(errW)
		return
	}
	fmt.Println(os.Args[2])
	contenu, err = os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(contenu))
}
