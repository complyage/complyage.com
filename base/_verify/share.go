package verify

import (
	"encoding/json"
	"fmt"

	"github.com/ralphferrara/aria/base/encrypt"
	"github.com/ralphferrara/aria/storage"
)

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

func ShareLoad(s *storage.Storage, uuid string, vType DataType, privateKey string) (Data, error) {
	v := Verification{}
	v.UUID = uuid
	v.Type = vType
	//||----------------------------------------------------------------------------------------------||
	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification JSON
	//||------------------------------------------------------------------------------------------------||
	rawData, err := s.Get(v.ObjectName(true))
	if err != nil {
		return Data{}, fmt.Errorf("get %q failed: %w", v.ObjectName(true), err)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Decrypt Data
	//||------------------------------------------------------------------------------------------------||
	fmt.Println("Raw Private:", string(privateKey))
	decryptedData, err := encrypt.DecryptWithPrivateKey(rawData, privateKey)
	fmt.Println("Decrypted Data:", string(decryptedData))
	if err != nil {
		fmt.Println("Decryption Error:", err)
		return Data{}, fmt.Errorf("decrypt %q failed: %w", v.ObjectName(true), err)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Unmarshal
	//||------------------------------------------------------------------------------------------------||
	err = json.Unmarshal(decryptedData, &v.Encrypted)
	if err != nil {
		return Data{}, fmt.Errorf("unmarshal %q failed: %w", v.ObjectName(true), err)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||
	return v.Encrypted.Data, nil
}
