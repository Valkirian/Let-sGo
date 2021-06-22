package cmd

import (
	"fmt"
	"io/ioutil"
)

func SubdomainFinder(filename, domain string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error leyendo el archivo", err)
	}
	// TODO Leer el archivo linea por linea y hacer las peticiones get para comprobar el 200 y que cada subdominio responda.

	for word := range string(data) {
		fmt.Println(string(word), domain)
	}
	return "hola"
}
