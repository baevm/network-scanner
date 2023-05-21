package port

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
	"time"
)

type Scanner struct {
	Hostname  string
	Protocol  string
	PortCount int
}

type ScanResult struct {
	Port     int
	State    string
	Protocol string
}

const (
	closed = "closed"
	open   = "open"
)

func Scan(hostname, protocol string, portCount int) {
	fmt.Printf("Starting scanning on %s \n", hostname)

	scanner := Scanner{
		Hostname:  hostname,
		Protocol:  protocol,
		PortCount: portCount,
	}

	scanResults := scanner.scanHost()

	if len(scanResults) == 0 {
		fmt.Printf("Not found open ports at %s", hostname)
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

func (s Scanner) scanPort(port int, ch chan ScanResult) {
	res := ScanResult{Port: port, Protocol: s.Protocol}

	address := s.Hostname + ":" + strconv.Itoa(port)

	con, err := net.DialTimeout(s.Protocol, address, time.Second)

	if err != nil {
		res.State = closed
		ch <- res
	} else {
		res.State = open
		ch <- res
		defer con.Close()
	}
}

func (s Scanner) scanHost() []ScanResult {
	var res []ScanResult
	ch := make(chan ScanResult, s.PortCount)

	for i := 1; i <= s.PortCount; i++ {
		go s.scanPort(i, ch)
	}

	for i := 1; i <= s.PortCount; i++ {
		scanRes := <-ch

		if scanRes.State == open {
			res = append(res, scanRes)
		}
	}

	close(ch)

	return res
}
