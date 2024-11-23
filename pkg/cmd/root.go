package cmd

import "yoreyuan/deploy-maven-localRepository/pkg/mvn"

func Execute() {
	c := initArgs()
	if !c {
		return
	}

	mvn.Run()
}
