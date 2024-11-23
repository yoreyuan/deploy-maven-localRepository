package cmd

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	c "yoreyuan/deploy-maven-localRepository/pkg/constant"
)

func initArgs() bool {
	s := flag.String("s", "", "Alternate path for the user settings file")
	repo := flag.String("repo", "", "The path to the local repository maven will use to store artifacts.")
	url := flag.String("url", "", "The URL of the repository maven.")
	repoId := flag.String("repoId", "", "The ID of the repository.")

	verbose := flag.Bool("verbose", false, "Enable verbose mode")
	mvnX := flag.Bool("X", false, "Produce execution debug output")
	help := flag.Bool("help", false, "Display help information")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return false
	}
	if *verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		c.Verbose = true
	}
	if *mvnX {
		c.MvnDebug = true
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Err(err).Msg("The home directory is obtained abnormally. The home directory will be set to the root directory")
		homeDir = "/root"
	}

	c.SettingXml = *s
	c.LocalRepository = *repo

	if c.SettingXml == "" {
		c.SettingXml = fmt.Sprintf("%s/.m2/settings.xml", homeDir)
	}
	log.Info().Msgf("setting xml is %s", c.SettingXml)

	if c.LocalRepository == "" {
		c.LocalRepository = fmt.Sprintf("%s/.m2/repository", homeDir)
	}
	log.Info().Msgf("local repository is %s", c.LocalRepository)

	if *url != "" {
		c.RepoUrl = *url
	}
	log.Info().Msgf("repository url is %s", c.RepoUrl)

	if *repoId != "" {
		c.RepoId = *repoId
	}
	log.Info().Msgf("repository id is %s", c.RepoId)

	if strings.HasSuffix(c.LocalRepository, c.Separator) {
		c.LocalRepository = strings.TrimRight(c.LocalRepository, c.Separator)
	}

	return true
}
