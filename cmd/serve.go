package main

import (
	"context"

	"github.com/BigDwarf/sahtian/internal/app"
	"github.com/BigDwarf/sahtian/internal/config"

	lib_errors "github.com/BigDwarf/sahtian/internal/errors"
	"github.com/BigDwarf/sahtian/internal/log"
	"github.com/BigDwarf/sahtian/version"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var serverCmd = &cobra.Command{
	Use:           "serve",
	Short:         "sahtian backend server",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {

		conf, err := config.Init(cfgFile)
		if err != nil {
			log.Error("Init config failed",
				zap.Error(err),
				zap.String(log.FieldKeyErrorStack, lib_errors.StackTrace(err)),
			)

			return errors.Wrap(err, "failed to init config")
		}

		err = log.Init(conf.Logger)
		if err != nil {
			log.Error("Init logger failed",
				zap.Error(err),
				zap.String(log.FieldKeyErrorStack, lib_errors.StackTrace(err)),
			)

			return errors.Wrap(err, "failed to init logger")
		}

		log.Info("Starting server... ")

		defer log.Sync()

		log.Sugar().Infof("Using config file: %s", viper.ConfigFileUsed())

		app := app.NewServerApplication(conf)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go waitSignalExit(cancel)

		err = app.Run(ctx)
		if err != nil && !errors.Is(err, context.Canceled) {
			log.Error("Run application failed",
				zap.Error(err),
				zap.String(log.FieldKeyErrorStack, lib_errors.StackTrace(err)),
			)

			return err
		}

		log.Info("Started server... ")
		<-ctx.Done()

		app.Shutdown()

		return nil
	},
}
