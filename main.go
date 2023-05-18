package main

import (
	"flag"
	"fmt"
	"os"
	"port-scanner/port"
	"sort"
	"text/tabwriter"
)

func main() {
	var address string
	var protocol string
	var portCount int

	flag.StringVar(&address, "a", "localhost", "Address to scan ports")
	flag.StringVar(&protocol, "p", "tcp", "Protocol")
	flag.IntVar(&portCount, "c", 1024, "Port count")
	flag.Parse()

	fmt.Println("Starting scanning....")

	scanner := port.Scanner{
		Hostname:  address,
		Protocol:  protocol,
		PortCount: portCount,
	}

	scanResults := scanner.ScanHost()

	if len(scanResults) == 0 {
		fmt.Printf("Not found open ports at %s", address)
		return
	}

	sort.Slice(scanResults, func(i, j int) bool {
		return scanResults[i].Port < scanResults[j].Port
	})

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "%s\t%v\t%v\t\n", "PORT", "STATUS", "PROTOCOL")

	for _, res := range scanResults {
		fmt.Fprintf(w, "%d\t%v\t%v\t\n", res.Port, res.State, res.Protocol)
	}

	w.Flush()
}
