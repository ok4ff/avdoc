package arch

import (
	"fmt"
	"log"
	"os/exec"
)

func addUser(name, workdir, shell string) error {
	log.Print("creating user...")
	out, err := exec.Command(
		"useradd",
		"-mG", "wheel",
		"-s", shell,
		name).Output()
	if err != nil {
		return fmt.Errorf("created user failed: %w: %s", err, out)
	}

	return nil
}
