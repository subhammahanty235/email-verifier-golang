package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("domain , hasMX , hasSPF , sprRecord , hasDMARC , dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error : could not read from input ", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var sprRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error occured")
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtrecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error occured")
	}

	for _, record := range txtrecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			sprRecord = record
		}
	}

	dmarcrecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error occured")
	}

	for _, record := range dmarcrecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}

	fmt.Printf("%v %v %v %v %v %v", domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord)

}
