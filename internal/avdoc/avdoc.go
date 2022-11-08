package avdoc

import (
	"avdoc/internal/arch"
	"avdoc/internal/av"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	log.Printf("Container started...")

	if err := arch.Init(arch.Config{}); err != nil {
		log.Fatal(err)
	}

	if err := av.Init(av.Config{}); err != nil {
		log.Fatal(err)
	}

	log.Print("Container ready!")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	s := <-interrupt
	log.Printf("syscall: %s. Container stoped..", s.String())
}
