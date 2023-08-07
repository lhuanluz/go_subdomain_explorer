package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func resolveSubdomain(domain, subdomain string, wg *sync.WaitGroup, sem chan struct{}, outputFile *os.File, ticker *time.Ticker) {
	defer wg.Done()

	if ticker != nil {
		<-ticker.C
	}

	sem <- struct{}{}
	defer func() { <-sem }()

	fullDomain := fmt.Sprintf("%s.%s", subdomain, domain)
	addresses, err := net.LookupHost(fullDomain)
	if err == nil && len(addresses) > 0 {
		if outputFile != nil {
			outputFile.WriteString(fullDomain + "\n")
		} else {
			fmt.Println(fullDomain)
		}
	}
}

func main() {
	concurrencyPtr := flag.Int("c", 10, "The concurrency level (number of goroutines that can run in parallel)")
	outputFilePtr := flag.String("f", "", "Output file to write the results")
	rateLimitPtr := flag.Float64("r", 0, "Rate limit in requests per second")

	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("Usage: subdomain-enumerator <domain> <subdomain-wordlist> [-c <concurrency-level>] [-f <output-file>] [-r <rate-per-second>]")
		return
	}

	if *concurrencyPtr < 1 {
		fmt.Println("Invalid concurrency level. It should be a positive integer greater than 0.")
		return
	}

	var ticker *time.Ticker
	if *rateLimitPtr > 0 {
		ticker = time.NewTicker(time.Second / time.Duration(*rateLimitPtr))
		defer ticker.Stop()
	}

	domain := flag.Args()[0]
	wordlist := flag.Args()[1]

	sem := make(chan struct{}, *concurrencyPtr)

	var outputFile *os.File
	if *outputFilePtr != "" {
		var err error
		outputFile, err = os.Create(*outputFilePtr)
		if err != nil {
			fmt.Printf("Error creating output file: %s\n", err)
			return
		}
		defer outputFile.Close()
	}

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
		go resolveSubdomain(domain, subdomain, &wg, sem, outputFile, ticker)
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading wordlist: %s\n", err)
	}
}

