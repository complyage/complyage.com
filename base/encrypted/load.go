package encrypted

import (
	"encoding/json"

	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/encrypt"
)

//||------------------------------------------------------------------------------------------------||
//|| Load
//||------------------------------------------------------------------------------------------------||

func Load(uuid string) (Encrypted, error) {
	enc := Encrypted{
		UUID: uuid,
	}
	filename := enc.Filename()

	data, err := Storage.Get(filename)
	if err != nil {
		return Encrypted{}, err
	}

	enc.Cipher = data
	return enc, nil
}

//||------------------------------------------------------------------------------------------------||
//|| LoadADDR
//||------------------------------------------------------------------------------------------------||

func LoadADDR(uuid string, privateKey string) (types.Address, error) {
	enc, err := Load(uuid)
	if err != nil {
		return types.Address{}, err
	}
	bytes, dErr := encrypt.DecryptWithPrivateKey(enc.Cipher, privateKey)
	if dErr != nil {
		return types.Address{}, dErr
	}
	var addr types.Address
	err = json.Unmarshal(bytes, &addr)
	if err != nil {
		return types.Address{}, err
	}
	return addr, nil
}

//||------------------------------------------------------------------------------------------------||
//|| SaveCRCD / LoadCRCD
//||------------------------------------------------------------------------------------------------||

func LoadCRCD(uuid string, privateKey string) (types.CreditCard, error) {
	enc, err := Load(uuid)
	if err != nil {
		return types.CreditCard{}, err
	}
	bytes, dErr := encrypt.DecryptWithPrivateKey(enc.Cipher, privateKey)
	if dErr != nil {
		return types.CreditCard{}, dErr
	}
	var cc types.CreditCard
	err = json.Unmarshal(bytes, &cc)
	if err != nil {
		return types.CreditCard{}, err
	}
	return cc, nil
}

//||------------------------------------------------------------------------------------------------||
//|| SaveFACE / LoadFACE
//||------------------------------------------------------------------------------------------------||

func LoadFACE(uuid string, privateKey string) (types.Facial, error) {
	enc, err := Load(uuid)
	if err != nil {
		return types.Facial{}, err
	}
	bytes, dErr := encrypt.DecryptWithPrivateKey(enc.Cipher, privateKey)
	if dErr != nil {
		return types.Facial{}, dErr
	}
	var face types.Facial
	err = json.Unmarshal(bytes, &face)
	if err != nil {
		return types.Facial{}, err
	}
	return face, nil
}

//||------------------------------------------------------------------------------------------------||
//|| SaveIDEN / LoadIDEN
//||------------------------------------------------------------------------------------------------||

func LoadIDEN(uuid string, privateKey string) (types.Identification, error) {
	if uuid == "" || privateKey == "" {
		return types.Identification{}, app.Err("Encrypted").Error("ENCRYPT_LOAD_FAILED")
	}
	enc, err := Load(uuid)
	if err != nil {
		return types.Identification{}, err
	}
	bytes, dErr := encrypt.DecryptWithPrivateKey(enc.Cipher, privateKey)
	if dErr != nil {
		return types.Identification{}, dErr
	}
	var iden types.Identification
	err = json.Unmarshal(bytes, &iden)
	if err != nil {
		return types.Identification{}, err
	}
	return iden, nil
}

//||------------------------------------------------------------------------------------------------||
//|| SaveMAIL / LoadMAIL
//||------------------------------------------------------------------------------------------------||

func LoadMAIL(uuid string, privateKey string) (types.EmailAddress, error) {
	enc, err := Load(uuid)
	if err != nil {
		return types.EmailAddress{}, err
	}
	bytes, dErr := encrypt.DecryptWithPrivateKey(enc.Cipher, privateKey)
	if dErr != nil {
		return types.EmailAddress{}, dErr
	}
	var mail types.EmailAddress
	err = json.Unmarshal(bytes, &mail)
	if err != nil {
		return types.EmailAddress{}, err
	}
	return mail, nil
}

//||------------------------------------------------------------------------------------------------||
//|| SavePHNE / LoadPHNE
//||------------------------------------------------------------------------------------------------||

func LoadPHNE(uuid string, privateKey string) (types.PhoneNumber, error) {
	enc, err := Load(uuid)
	if err != nil {
		return types.PhoneNumber{}, err
	}
	bytes, dErr := encrypt.DecryptWithPrivateKey(enc.Cipher, privateKey)
	if dErr != nil {
		return types.PhoneNumber{}, dErr
	}
	var phone types.PhoneNumber
	err = json.Unmarshal(bytes, &phone)
	if err != nil {
		return types.PhoneNumber{}, err
	}
	return phone, nil
}

//||------------------------------------------------------------------------------------------------||
//|| LoadPreview: Decrypt and return as generic JSON (any type)
//||------------------------------------------------------------------------------------------------||

func LoadPreview(uuid string, privateKey string, dataType string) (map[string]interface{}, error) {
	//|| Validate input
	if uuid == "" || privateKey == "" || dataType == "" {
		return nil, app.Err("Encrypted").Error("LOAD_PREVIEW_INVALID_INPUT")
	}

	//|| Load raw encrypted record
	enc, err := Load(uuid)
	if err != nil {
		return nil, err
	}

	//|| Decrypt cipher
	bytes, dErr := encrypt.DecryptWithPrivateKey(enc.Cipher, privateKey)
	if dErr != nil {
		return nil, dErr
	}

	//|| Decode to a generic JSON map
	var generic map[string]interface{}
	err = json.Unmarshal(bytes, &generic)
	if err != nil {
		return nil, err
	}

	//|| Optionally: include metadata about type or UUID
	preview := map[string]interface{}{
		"uuid": uuid,
		"type": dataType,
		"data": generic,
	}

	return preview, nil
}
