package arch

import "log"

type (
	Config struct{}
)

func Init(c Config) error {
	log.Print("Arch linux prepare...")
	log.Print("Arch linux ready")

	return nil
}
