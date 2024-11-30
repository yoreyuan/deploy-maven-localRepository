package cmd

import (
	"yoreyuan/deploy-maven-localRepository/pkg/mvn"
)

func Execute() {
	err := initArgs()
	if err != nil {
		return
	}

	mvn.Init()
	mvn.Run()
}
