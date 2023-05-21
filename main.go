package main

import (
	"flag"
	"log"
	"os"
	"network-scanner/network"
	"network-scanner/port"
)

var scanOption int

var hostname string
var protocol string
var portCount int

func main() {
	flag.IntVar(&scanOption, "o", 1, "Select scanning option: \n 1. Port scanner \n 2. Network scanner")
	flag.StringVar(&hostname, "a", "localhost", "Address to scan ports")
	flag.StringVar(&protocol, "p", "tcp", "Protocol to scan ports")
	flag.IntVar(&portCount, "c", 1024, "Port count")
	flag.Parse()

	switch scanOption {
	case 1:
		port.Scan(hostname, protocol, portCount)
	case 2:
		network.StartScan()
	default:
		log.Println("No such option.")
		os.Exit(1)
	}
}
