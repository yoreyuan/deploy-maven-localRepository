package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"yoreyuan/deploy-maven-localRepository/pkg/config"
	"yoreyuan/deploy-maven-localRepository/pkg/constant"
)

func initArgs() error {
	configFile := flag.String("config", "", "Configuration file path")

	s := flag.String("s", "", "Alternate path for the user settings file")
	repo := flag.String("repo", "", "The path to the local repository maven will use to store artifacts")
	url := flag.String("url", "", "The URL of the repository maven")
	repoId := flag.String("repoId", "", "The ID of the repository")

	verbose := flag.Bool("verbose", false, "Enable app verbose mode")
	mvnX := flag.Bool("X", false, "Produce execution debug output")
	help := flag.Bool("help", false, "Display help information")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return errors.New("print help")
	}

	var conf *config.Config

	if *configFile != "" {
		err := initConfig()
		if err != nil {
			return err
		}
		conf = config.GetConfig()
	} else {
		conf = config.NewDefaultConfig()

		if *verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			conf.Verbose = true
		}
		if *mvnX {
			conf.MvnDebug = true
		}

		homeDir := getUserHomeDir()
		conf.SettingXml = *s
		conf.LocalRepository = *repo
		if conf.SettingXml == "" {
			conf.SettingXml = fmt.Sprintf("%s/.m2/settings.xml", homeDir)
		}
		log.Info().Msgf("setting xml is %s", conf.SettingXml)

		if conf.LocalRepository == "" {
			conf.LocalRepository = fmt.Sprintf("%s/.m2/repository", homeDir)
		}
		log.Info().Msgf("local repository is %s", conf.LocalRepository)

		if *url != "" {
			conf.RepoUrl = *url
		}
		log.Info().Msgf("repository url is %s", conf.RepoUrl)

		if *repoId != "" {
			conf.RepoId = *repoId
		}
		log.Info().Msgf("repository id is %s", conf.RepoId)

		config.SetConfig(conf)
	}

	if strings.HasSuffix(conf.LocalRepository, constant.Separator) {
		conf.LocalRepository = strings.TrimRight(conf.LocalRepository, constant.Separator)
	}
	if strings.HasPrefix(conf.LocalRepository, "~") {
		newLocalRepository := strings.Replace(conf.LocalRepository, "~", getUserHomeDir(), 1)
		conf.LocalRepository = newLocalRepository
	}
	if strings.HasPrefix(conf.SettingXml, "~") {
		newSettingXml := strings.Replace(conf.SettingXml, "~", getUserHomeDir(), 1)
		conf.SettingXml = newSettingXml
	}

	log.Info().Interface("config", conf).Msg("print config")

	return nil
}

func getUserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Err(err).Msg("The home directory is obtained abnormally. The home directory will be set to the root directory")
		homeDir = "/root"
	}
	return homeDir
}
