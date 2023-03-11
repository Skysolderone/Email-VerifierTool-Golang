/*
Create time at 2023年3月11日0011下午 12:33:25
Create User at Administrator
*/

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
	fmt.Printf("domain,hasMx,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")
	for scanner.Scan() {
		checkdomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error:could not read scanner:", err)
	}

}
func checkdomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}
	if len(mxRecord) > 0 {
		hasMx = true
	}
	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=dmarc") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v,%v,%v,%v,%v,%v", domain, hasMx, hasSPF, hasDMARC, spfRecord, dmarcRecord)
}
