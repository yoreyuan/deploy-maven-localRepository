package mvn

import (
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

func doClean(dirs map[string]bool) error {
	for item := range dirs {
		entries, err := os.ReadDir(item)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			fileName := entry.Name()
			if hashCleanSuffixs(fileName) {
				cleanFile := filepath.Join(item, fileName)
				err = os.Remove(cleanFile)
				if err != nil {
					log.Error().Err(err).Msgf("error deleting file: %s", cleanFile)
					return err
				}
				log.Info().Msgf("file deletion succeeded: %s", cleanFile)
			}
		}
	}
	return nil
}

func hashCleanSuffixs(fileName string) bool {
	for _, s := range conf.Clean.CleanSuffixs {
		if strings.HasSuffix(fileName, s) {
			return true
		}
	}
	return false
}
