package encrypted

import (
	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/log"
)

//||------------------------------------------------------------------------------------------------||
//|| Save
//||------------------------------------------------------------------------------------------------||

func (e Encrypted) Save() error {
	filename := e.Filename()
	app.Log.Data("Saving Encrypted to Storage: " + filename)
	log.PrettyPrint(e)

	err := Storage.Put(filename, e.Cipher)
	if err != nil {
		return err
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| SaveADDR / LoadADDR
//||------------------------------------------------------------------------------------------------||

func SaveADDR(publicKey string, uuid string, addr types.Address) error {
	encryptedData, err := MakeCipher(publicKey, addr)
	if err != nil {
		return err
	}
	instance := Encrypted{
		UUID:   uuid,
		Type:   types.DataTypeADDR,
		Cipher: encryptedData,
	}
	return instance.Save()
}

//||------------------------------------------------------------------------------------------------||
//|| SaveCRCD / LoadCRCD
//||------------------------------------------------------------------------------------------------||

func SaveCRCD(publicKey string, uuid string, cc types.CreditCard) error {
	encryptedData, err := MakeCipher(publicKey, cc)
	if err != nil {
		return err
	}
	instance := Encrypted{
		UUID:   uuid,
		Type:   types.DataTypeCRCD,
		Cipher: encryptedData,
	}
	return instance.Save()
}

//||------------------------------------------------------------------------------------------------||
//|| SaveFACE / LoadFACE
//||------------------------------------------------------------------------------------------------||

func SaveFACE(publicKey string, uuid string, face types.Facial) error {
	encryptedData, err := MakeCipher(publicKey, face)
	if err != nil {
		return err
	}
	instance := Encrypted{
		UUID:   uuid,
		Type:   types.DataTypeFACE,
		Cipher: encryptedData,
	}
	return instance.Save()
}

//||------------------------------------------------------------------------------------------------||
//|| SaveIDEN / LoadIDEN
//||------------------------------------------------------------------------------------------------||

func SaveIDEN(publicKey string, uuid string, iden types.Identification) error {
	encryptedData, err := MakeCipher(publicKey, iden)
	if err != nil {
		return err
	}
	instance := Encrypted{
		UUID:   uuid,
		Type:   types.DataTypeIDEN,
		Cipher: encryptedData,
	}
	return instance.Save()
}

//||------------------------------------------------------------------------------------------------||
//|| SaveMAIL / LoadMAIL
//||------------------------------------------------------------------------------------------------||

func SaveMAIL(publicKey string, uuid string, mail types.EmailAddress) error {
	encryptedData, err := MakeCipher(publicKey, mail)
	if err != nil {
		return err
	}
	instance := Encrypted{
		UUID:   uuid,
		Type:   types.DataTypeMAIL,
		Cipher: encryptedData,
	}
	return instance.Save()
}

//||------------------------------------------------------------------------------------------------||
//|| SavePHNE / LoadPHNE
//||------------------------------------------------------------------------------------------------||

func SavePHNE(publicKey string, uuid string, phone types.PhoneNumber) error {
	encryptedData, err := MakeCipher(publicKey, phone)
	if err != nil {
		return err
	}
	instance := Encrypted{
		UUID:   uuid,
		Type:   types.DataTypePHNE,
		Cipher: encryptedData,
	}
	return instance.Save()
}
