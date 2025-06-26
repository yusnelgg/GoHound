package main

import (
	"fmt"
	"os"

	"github.com/yusnelgg/gohound/internal/scanner"
	"github.com/yusnelgg/gohound/internal/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gohound <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  scan <domain>   - Scan domain from CLI")
		fmt.Println("  serve           - Start web server")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "scan":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gohound scan <domain>")
			return
		}
		domain := os.Args[2]
		fmt.Println("🔍 Scanning domain:", domain)
		scanner.ScanSubdomains(domain)
	case "serve":
		server.Start()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
