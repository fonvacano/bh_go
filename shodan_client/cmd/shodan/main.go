package main

import (
	"fmt"
	"log"
	"os"
	"shodan/shodan.go/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}

	key := os.Getenv("SHODAN_API_KEY")
	client := shodan.New(key)
	info, err := client.APIInfo()
	if err != nil {
		log.Fatalln("Cant get api info")
	}

	fmt.Printf("Query Credits:%d\nScan Credits:%d\n", info.QueryCredits, info.ScanCredits)

	search, err := client.HostSearch(os.Args[1])
	if err != nil {
		log.Fatalln("Cant get search result")
	}

	for _, host := range search.Matches {
		fmt.Printf("%18s%8s\n", host.IPString, host.Port)
	}
}
