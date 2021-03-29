package helper

import (
	"io/ioutil"
)

func LeerDat(nombrearchivo string) ([]byte,error){
	data, err :=ioutil.ReadFile(nombrearchivo)
	return data,err
}