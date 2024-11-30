package config

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Verbose         bool   `mapstructure:"verbose,default=false"`
	LogLevel        string `yaml:"logLevel"`
	LocalRepository string `mapstructure:"localRepository,default=~/.m2/repository"`
	MvnRepo         `mapstructure:"repo"`
	Clean           RepoClean `mapstructure:"clean"`
}

type MvnRepo struct {
	CommandName    string `mapstructure:"commandName,default=mvn"`
	SettingXml     string `mapstructure:"settingXml,default=~/.m2/settings.xml"`
	RepoId         string `mapstructure:"id,default=yore_nexus"`
	RepoUrl        string `mapstructure:"url,default=http://nexus.yore.cn/repository/maven-releases"`
	MvnDebug       bool   `mapstructure:"debug,default=false"`
	ExcludeSuffixs `mapstructure:"excludeSuffixs"`
}

type ExcludeSuffixs []string

type RepoClean struct {
	Enable       bool `mapstructure:"enable,default=false"`
	CleanSuffixs `mapstructure:"suffixs"`
}

type CleanSuffixs []string

func NewDefaultConfig() *Config {
	return &Config{
		Verbose:         false,
		LogLevel:        "INFO",
		LocalRepository: "",
		MvnRepo: MvnRepo{
			CommandName:    "mvn",
			SettingXml:     "",
			RepoId:         "yore_nexus",
			RepoUrl:        "http://nexus.yore.cn/repository/maven-releases",
			MvnDebug:       false,
			ExcludeSuffixs: ExcludeSuffixs{},
		},
		Clean: RepoClean{
			Enable:       false,
			CleanSuffixs: CleanSuffixs{},
		},
	}
}

var globalConf = &Config{}

func GetConfig() *Config {
	return globalConf
}

func SetConfig(conf *Config) {
	globalConf = conf
}

func Parse() (string, error) {
	defaultConfigFile, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	var configFileName string
	pflag.StringVar(&configFileName,
		"config",
		filepath.Join(defaultConfigFile, "config.yaml"),
		"explicitly defines the path, name and extension of the config file")
	pflag.Parse()
	return configFileName, err
}

func Load(fileName string) (*Config, error) {
	viper.SetConfigFile(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn().Err(err).Msg("viper read in config failed, use env or default")
		return nil, err
	}

	conf := NewDefaultConfig()
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	globalConf = conf
	logLevel, err := zerolog.ParseLevel(conf.LogLevel)
	if err != nil {
		return nil, err
	}
	zerolog.SetGlobalLevel(logLevel)

	return conf, nil
}
