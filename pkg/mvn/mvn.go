package mvn

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"yoreyuan/deploy-maven-localRepository/pkg/config"
	"yoreyuan/deploy-maven-localRepository/pkg/constant"
	"yoreyuan/deploy-maven-localRepository/pkg/utils"
)

type MvnModel struct {
	File       string
	PomFile    string
	Packaging  string //jar,pom,exe etc.
	Classifier string //sources,tests,linux-x86_64 etc.
}

var (
	pomDirSet = utils.NewSet()
	errMvnCmd []string
	conf      *config.Config
)

func Init() {
	errMvnCmd = make([]string, 0)
	conf = config.GetConfig()
}

func executeMvn(pomDir map[string]bool) error {
	var index int = 0
	c := len(pomDir)

	for item := range pomDir {
		if conf.Verbose {
			log.Info().Msgf("Start processing packages in folder %s", item)
		}
		entries, err := os.ReadDir(item)
		if err != nil {
			return err
		}

		pomFile := ""
		itemSplit := strings.Split(item, constant.Separator)
		if len(itemSplit) <= 1 {
			return errors.New("the maven repository path format is incorrect")
		}
		parentDirName := itemSplit[len(itemSplit)-1]

		// Iterate through each entry in the directory
		var j = 0
		for _, entry := range entries {
			fileName := entry.Name()
			if conf.Verbose {
				j++
				if j == len(entries) {
					println(" └ ", fileName)
				} else {
					println(" ├ ", fileName)
				}
			}
			//Only one pom file is currently found in a directory
			if strings.HasSuffix(fileName, ".pom") {
				pomFile = filepath.Join(item, fileName)
			}
		}
		if pomFile == "" {
			log.Error().Msgf("maven pom file not found in folder: %s", item)
			continue
		}

		var itemSet = utils.NewSet()
		for _, entry := range entries {
			fileName := entry.Name()
			if strings.HasSuffix(fileName, ".pom") {
				continue
			}

			// Skip files not included
			c := false
			for _, s := range conf.ExcludeSuffixs {
				if strings.HasSuffix(fileName, s) {
					c = true
					break
				}
			}
			if c {
				continue
			}

			if !strings.Contains(fileName, "-"+parentDirName) {
				continue
			}

			itemSet.Add(fileName)
		}

		if itemSet.Size() == 0 {
			mvnModel := MvnModel{PomFile: pomFile, File: pomFile, Packaging: "pom"}
			err = mvnDeploy(&mvnModel)
			if err != nil {
				log.Err(err).Msg("Please execute the command manually later!")
				//return err
			}
		} else {
			for file := range *itemSet.GetSet() {
				arr1 := strings.Split(file, parentDirName+".")
				if len(arr1) < 1 {
					log.Error().Msgf("incorrect file format: %s", file)
					continue
				}

				mvnModel := MvnModel{PomFile: pomFile, File: filepath.Join(item, file)}
				if len(arr1) == 1 {
					arr2 := strings.Split(file, parentDirName+"-")
					s2 := arr2[len(arr2)-1]
					arr3 := strings.Split(s2, ".")
					if len(arr3) < 2 {
						log.Error().Msgf("incorrect file format: %s", file)
						continue
					}
					mvnModel.Packaging = arr3[len(arr3)-1]
					mvnModel.Classifier = strings.TrimRight(s2, "."+mvnModel.Packaging)
				} else {
					mvnModel.Packaging = arr1[len(arr1)-1]
				}

				err = mvnDeploy(&mvnModel)
				if err != nil {
					log.Err(err).Msg("Please execute the command manually later!")
				}
			}
		}

		//if index > 2 {
		//	break
		//}
		index++
		log.Info().Msgf("progress is %.2f%%", float32(index)/float32(c)*100)
	}

	if len(errMvnCmd) > 1 {
		log.Warn().Msg("An exception occurred when publishing the following package. Please execute it manually.")
		for _, s := range errMvnCmd {
			println(s)
		}
	}

	return nil
}

func mvnDeploy(mvnModel *MvnModel) error {
	if conf.Verbose {
		log.Info().Msg(utils.Obj2JsonStr(mvnModel))
	}
	var cmd *exec.Cmd
	cmdArgs := make([]string, 0)
	cmdArgs = append(cmdArgs,
		"-s", conf.SettingXml,
		"deploy:deploy-file",
		"-Dfile="+mvnModel.File,
		"-DpomFile="+mvnModel.PomFile,
		"-Dclassifier="+mvnModel.Classifier,
		"-Durl="+conf.RepoUrl,
		"-DrepositoryId="+conf.RepoId)
	if conf.MvnDebug {
		cmdArgs = append(cmdArgs, "-X")
	}

	if mvnModel.Packaging == "jar" || mvnModel.Packaging == "" {
		cmdArgs = append(cmdArgs, "-Dpackaging=jar")
	} else {
		cmdArgs = append(cmdArgs, "-Dpackaging="+mvnModel.Packaging)
	}
	cmd = exec.Command(conf.CommandName, cmdArgs...)
	if conf.Verbose {
		log.Info().Msg(cmd.String())
	}

	// Create a buffer to capture the standard output of a command
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run the command and wait for it to complete
	err := cmd.Run()

	// Print the output of a command
	fmt.Println(out.String())

	if err != nil {
		errMvnCmd = append(errMvnCmd, cmd.String())
		return err
	}

	return nil
}
