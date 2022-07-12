package main

/** Useful links
https://pkg.go.dev/os
https://pkg.go.dev/os/exec
https://www.consul.io/docs/discovery/checks#check-scripts
**/

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "google.com", "-c1")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", out)
		os.Exit(0)
	}
}
