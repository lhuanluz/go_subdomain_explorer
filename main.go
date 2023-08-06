package main

import (
    "fmt"
    "bufio"
    "os"
    "net"
    "sync"
)

func resolveSubdomain(domain, subdomain string, wg *sync.WaitGroup) {
	defer wg.Done()

	fullDomain := fmt.Sprintf("%s.%s", subdomain, domain)
	addresses, err := net.LookupHost(fullDomain)
	if err == nil && len(addresses) > 0 {
		fmt.Println(fullDomain)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: subdomain_explorer <domain> <subdomain-wordlist>")
		return
	}

	domain := os.Args[1]
	wordlist := os.Args[2]

	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Printf("Error opening wordlist: %s\n", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		subdomain := scanner.Text()
		wg.Add(1)
		go resolveSubdomain(domain, subdomain, &wg)
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading wordlist: %s\n", err)
	}
}
