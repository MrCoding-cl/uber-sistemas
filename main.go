package main

import (
	"fmt"
	"./helper"
)

func main(){
	dato, err :=helper.LeerDat("datos/noche.dat")

	if err != nil {
		fmt.Println(err)
		//Si ocurre un error este se imprime
	}
	fmt.Println(dato[0]+dato[1])
}


