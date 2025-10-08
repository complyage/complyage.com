package sites

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Struct for Full Verification Response
//||------------------------------------------------------------------------------------------------||

func LogoHash(siteId, accountId int) string {
	hashInput := fmt.Sprintf("website-logo-%d-%d", siteId, accountId)
	hashBytes := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hashBytes[:]) + ".webp"
}
