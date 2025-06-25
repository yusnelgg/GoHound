package main

import (
	"fmt"
	"os"

	"github.com/yusnelgg/gohound/internal/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gohound scan <domain>")
		return
	}

	var oneArgument = os.Args[1]
	var twoArguments = os.Args[2]

	if oneArgument == "scan" {
		fmt.Println("Scanning domain:", twoArguments)
		scanner.ScanSubdomains(twoArguments)
	}
}
