package avdoc

import (
	"strings"

	"github.com/spf13/viper"
)

func initConfig() error {
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	// DEFAULT CONFIG
	viper.SetDefault("user.name", "av")
	viper.SetDefault("user.shell", "/usr/bin/zsh")
	viper.SetDefault("user.workdir", "/home/av")

	viper.SetDefault("arch.package", "mc, zsh")

	return nil
}
