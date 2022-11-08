package avdoc

import (
	"avdoc/internal/arch"
	"avdoc/internal/av"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

func Run() {
	log.Printf("Container started...")

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	if err := arch.Init(arch.Config{
		Username: viper.GetString("user.name"),
		Workdir:  viper.GetString("user.workdir"),
		Shell:    viper.GetString("user.shell"),
		Package:  viper.GetString("arch.package"),
	}); err != nil {
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
