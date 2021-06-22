package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func SubdomainFinder(filename, domain string) []string {
	// List of domains in ok.
	founds := []string{}

	file, err  := os.Open(filename)
	if err != nil {
		fmt.Println("Error al abrir el archivo")
	}
	defer file.Close()
	// Scanning for lines
	scanner := bufio.NewScanner(file)

	// Spliting by new line
	scanner.Split(bufio.ScanLines)
	text := []string{}

	// Getting all lines
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// Printing line by line
	for _, line := range text {
		subdomain := line + "." + domain
		res, err := http.Get("http://"+subdomain+"/")
		if err != nil {
			fmt.Println("Error llamando al subdominio: ", subdomain)
		}
		defer res.Body.Close()
		founds = append(founds, fmt.Sprintf("%v with status code %v", subdomain, res.StatusCode))
	}
	return founds
}
