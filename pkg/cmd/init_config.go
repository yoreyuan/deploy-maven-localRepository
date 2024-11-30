package cmd

import (
	"github.com/rs/zerolog/log"
	"yoreyuan/deploy-maven-localRepository/pkg/config"
)

func initConfig() error {
	configFile, err := config.Parse()
	if err != nil {
		log.Error().Err(err).Msg("get config file failed")
		return err
	}
	log.Info().Msgf("config file is %s", configFile)

	_, err = config.Load(configFile)
	if err != nil {
		log.Error().Err(err).Msg("load config form file failed")
		return err
	}

	return nil
}
