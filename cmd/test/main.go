package main

import (
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("whoami").Output()
	if err != nil {
		log.Fatalf("Er: %v", err)
	}

	log.Printf("user: %s", out)
}
