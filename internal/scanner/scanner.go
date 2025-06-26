package scanner

import (
	"fmt"
	"net"
	"time"
)

type ScanResult struct {
	Subdomain string `json:"subdomain"`
	OpenPorts []int  `json:"open_ports"`
}

var commonSubs = []string{
	"www", "mail", "ftp", "blog", "api", "dev", "test", "shop", "cpanel", "webmail",
}

var commonPorts = []int{22, 80, 443, 8080, 8443, 21, 25, 3306}

func ScanPortsConcurrent(host string, ports []int) []int {
	openPorts := []int{}
	results := make(chan int)

	timeout := 500 * time.Millisecond

	for _, port := range ports {
		go func(p int) {
			address := fmt.Sprintf("%s:%d", host, p)
			conn, err := net.DialTimeout("tcp", address, timeout)
			if err == nil {
				results <- p
				conn.Close()
			} else {
				results <- 0
			}
		}(port)
	}

	for range ports {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	return openPorts
}

func ScanSubdomains(domain string) []ScanResult {
	var results []ScanResult

	for _, sub := range commonSubs {
		fullDomain := sub + "." + domain
		_, err := net.LookupHost(fullDomain)
		if err == nil {
			openPorts := ScanPortsConcurrent(fullDomain, commonPorts)
			results = append(results, ScanResult{
				Subdomain: fullDomain,
				OpenPorts: openPorts,
			})
		}
	}

	return results
}
