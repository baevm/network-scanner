package port

import (
	"net"
	"strconv"
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

func (s Scanner) ScanHost() []ScanResult {

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
