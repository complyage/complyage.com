package verify_test

import (
	"os"
	"testing"

	"github.com/ralphferrara/aria/app"
)

func TestMain(m *testing.M) {
	// ✅ Initialize your app before running tests
	app.Init("../../../config.json")

	// ✅ Run tests
	code := m.Run()

	os.Exit(code)
}
