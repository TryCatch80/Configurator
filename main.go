package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Parse the -d flag to get the parent domain
	var parentDomain string
	flag.StringVar(&parentDomain, "d", "", "The parent domain to match against")
	flag.Parse()

	// If the parent domain is not supplied with the -d flag, print an error message and exit
	if parentDomain == "" {
		fmt.Println("Error: -d flag is required to specify the parent domain")
		return
	}

	// Read subdomains from standard input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		subdomain := strings.ToLower(scanner.Text())
		if subdomain == parentDomain || strings.HasSuffix(subdomain, "."+parentDomain) {
			fmt.Println(subdomain)
		}
	}
}
