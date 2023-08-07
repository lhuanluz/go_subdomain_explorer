# Subdomain Enumerator in Go

The Subdomain Enumerator is a simple CLI tool developed in Go to discover subdomains for a specified domain using wordlists and DNS resolution. It aids in the reconnaissance phase of a penetration test or bug bounty hunt, helping to uncover potential entry points and assets associated with a target domain.

## Features

-   Uses wordlists for potential subdomains.
-   Concurrent DNS resolution using Go's lightweight goroutines.
-   Displays found subdomains in real-time.

## Prerequisites

-   Go installed on your system.
-   A wordlist file containing potential subdomains (e.g., `subdomains.txt`).

## Installation

Clone the repository or copy the code into a file named `main.go`:

`git clone https://github.com/your-username/subdomain-enumerator.git
cd subdomain-enumerator` 

## Usage

To run the subdomain enumerator, use the following command:

`go run .\main.go <target-domain> <path-to-wordlist> [-c <concurrency-level>] [-f <output-file>]` 

For example:

`go run main.go example.com subdomains.txt` 

The output will display any discovered subdomains that exist.

To run with a concurrency level of 20:
`go run main.go -c 20 example.com subdomains.txt`

The -c flag specifies the concurrency level, indicating the number of goroutines that can run in parallel. Adjust this value based on your system's capabilities and the network conditions.

To save the output to a file use:

`go run .\main.go -f output.txt example.com subdomains.txt`

To limit the amount of requests per second use:

`go run .\main.go -r 5 example.com subdomains.txt`

This will ensure 5 requests per second.

## Contributing

Feel free to contribute to this project by submitting pull requests or opening issues with suggestions and improvements.

----------

Remember, always use tools responsibly and ethically. Ensure you have proper permissions when scanning or probing domains or servers.
