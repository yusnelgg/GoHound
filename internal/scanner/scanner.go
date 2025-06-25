package scanner

import (
	"fmt"
	"net"
)

var commonSubs = []string{
	"www", "mail", "ftp", "blog", "api", "dev", "test", "shop", "cpanel", "webmail",
}

func ScanSubdomains(domain string) []string {
	var found []string

	for _, sub := range commonSubs {
		fullDomain := sub + "." + domain
		ips, err := net.LookupHost(fullDomain)
		if err == nil {
			fmt.Printf("[+] Found: %s -> %v\n", fullDomain, ips)
			found = append(found, fullDomain)
		}
	}

	return found
}
