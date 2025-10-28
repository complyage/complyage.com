package encrypted

import (
	"encoding/json"

	"github.com/ralphferrara/aria/base/encrypt"
)

//||------------------------------------------------------------------------------------------------||
//|| MakeCipher: generic encryptor for any struct
//||------------------------------------------------------------------------------------------------||

func MakeCipher(publicKey string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return encrypt.EncryptWithPublicKey(jsonData, publicKey)
}
