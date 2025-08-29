package verify

import "github.com/ralphferrara/aria/app"

func LogInfo(message string, args ...interface{}) {
	app.Log.Info(message, args...)
}
