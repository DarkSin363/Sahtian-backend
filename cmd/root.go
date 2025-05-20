package main

import (
	"fmt"
	"github.com/BigDwarf/sahtian/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "",
		Long:          "Sahtian server : " + version.Version,
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       fmt.Sprintf("%s %s", version.Version, version.CommitHash),
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/default.yml)")
	viper.SetDefault("license", "none")

	cmd.AddCommand(serverCmd)
	return cmd
}

func Execute() (err error) {
	return rootCmd().Execute()
}
