package arch

import "log"

type (
	Config struct {
		Username string
		Workdir  string
		Shell    string
		Package  string
	}
)

func Init(c Config) error {
	log.Printf("Arch linux prepare, config: %+v", c)

	if err := instalPackage(c.Package); err != nil {
		return err
	}

	if err := addUser(c.Username, c.Workdir, c.Shell); err != nil {
		return err
	}

	log.Print("Arch linux ready")

	return nil
}
