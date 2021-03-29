package main

import (
	"./helper"
	"fmt"
)

func main(){
	dato, _ :=helper.LeerDat("datos/noche.dat")
	fmt.Printf("%s",dato)

}


