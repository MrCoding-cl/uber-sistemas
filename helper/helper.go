package helper

import (
	"bufio"
	"os"
)

func LeerDat(nombrearchivo string) ([]byte, error) {
	archivo, err := os.Open(nombrearchivo)

	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	stats, statsErr := archivo.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(archivo)
	_,err = bufr.Read(bytes)

	return bytes, err
}