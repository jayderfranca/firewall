package main

import (
	"os"
	"fmt"
)

func main() {

	iptables := FindExecutable("iptables")
	ulogd := FindExecutable("ulogd")

	fmt.Printf("IPTables: %s\n", iptables)
	fmt.Printf("ULogd: %s\n", ulogd)

	// finaliza o programa
	os.Exit(0)
}
