package verify

import (
	"github.com/ralphferrara/aria/log"
)

func LogInfo(message string, args ...interface{}) {
	log.Print(log.INFO, "Verify", message, args...)
}
