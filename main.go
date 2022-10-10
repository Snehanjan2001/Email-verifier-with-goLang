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
	fmt.Println("domain, hasMX,hasSPF.spfRecord,hasDMARC,dmarcRecord\n")

	//scanning the users input
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error could not be read from input :%v \n", err)
	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error : %v\n", err)

	}

	if len(mxRecords) > 0 {
		hasMX = true

	}

	//Finding the txt records

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error : %v\n", err)

	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break

		}
	}

	//Dmarc records
	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)
	if err != nil {
		log.Printf("Error : %v\n", err)

	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break

		}
	}

	fmt.Printf("%v %v %v %v %v %v", domain, hasMX, hasDMARC, hasSPF, spfRecord, dmarcRecord)

}
