package verify

import (
	"github.com/ralphferrara/aria/log"
)

func logInfo(message string, args ...interface{}) {
	log.Print(log.INFO, "Verify", message, args...)
}
