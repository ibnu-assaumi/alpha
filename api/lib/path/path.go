package path

import (
	"os"
	"regexp"
)

// Root : path folder inside api project
func Root() string {
	re := regexp.MustCompile(`^(.*api)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	return string(rootPath)
}
