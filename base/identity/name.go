package identity

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Filename - Returns filename for storing identity
//||------------------------------------------------------------------------------------------------||

func (i Identity) Filename() string {
	salt := app.Config.App.Salt
	input := fmt.Sprintf("%s:%d", salt, i.ID)
	hash := sha256.Sum256([]byte(input))
	hashStr := hex.EncodeToString(hash[:])[:16] // shorten to 16 chars
	return fmt.Sprintf("%s.json", hashStr)
}
