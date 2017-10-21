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

	variables := GetVariables("./_examples/config")
	fmt.Println(variables)

	// finaliza o programa
	os.Exit(0)
}
