package utils

import (
	"fmt"
	"os"
	"strings"
)

func Replace(str, r string, files ...string) {
	for _, f := range files {
		datos, err := os.ReadFile(f)
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			str := strings.Replace(string(datos), str, r, -1)
			d := []byte(str)
			err := os.WriteFile(f, d, 0644)
			if err == nil {
				fmt.Println(f + " Guardado correctamente")
			} else {
				fmt.Println("Ocurrio un problema")
			}
		}
	}
}
