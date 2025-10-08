package version

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var VERSION = "1.0.1"
var VERSION_FLOAT = 1.0

func init() {
	dir, _ := os.Getwd()
	for {
		versionFile := filepath.Join(dir, "VERSION")
		if _, err := os.Stat(versionFile); err == nil {
			data, _ := os.ReadFile(versionFile)
			VERSION = string(data)
			parts := strings.Split(VERSION, ".")
			if len(parts) >= 2 {
				major, _ := strconv.Atoi(parts[0])
				minor, _ := strconv.Atoi(parts[1])
				// Build float like 1.0, 1.1, 2.5
				VERSION_FLOAT, _ = strconv.ParseFloat(fmt.Sprintf("%d.%d", major, minor), 64)
			}
			break
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			panic("VERSION file not found")
		}
		dir = parent
	}
}
