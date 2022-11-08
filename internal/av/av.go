package av

import "log"

type Config struct{}

func Init(c Config) error {
	log.Print("Astra vim prepare...")
	log.Print("Astra vim ready")
	return nil
}
