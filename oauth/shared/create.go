package shared

import (
	"fmt"
	"time"

	"github.com/complyage/base/encrypted"
	"github.com/complyage/base/enforce"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"
	"github.com/ralphferrara/aria/base/random"
)

//||------------------------------------------------------------------------------------------------||
//|| Create
//||------------------------------------------------------------------------------------------------||

func Create(enforcement enforce.Enforcement) (OAuthSharedAccess, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Token
	//||------------------------------------------------------------------------------------------------||

	token := random.RandomString(32)

	//||------------------------------------------------------------------------------------------------||
	//|| DB/Storage
	//||------------------------------------------------------------------------------------------------||

	data := verify.Data{}
	privateKey := enforcement.User.Private

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Shared Data
	//||------------------------------------------------------------------------------------------------||

	for _, s := range enforcement.Scopes {
		switch s.Code {
		//||------------------------------------------------------------------------------------------------||
		//|| Address
		//||------------------------------------------------------------------------------------------------||
		case types.DataTypeADDR:
			dataAddr, _ := encrypted.LoadADDR(s.Verification, privateKey)
			data.ADDR = dataAddr
		//||------------------------------------------------------------------------------------------------||
		//|| Credit Card
		//||------------------------------------------------------------------------------------------------||
		case types.DataTypeCRCD:
			dataCRCD, _ := encrypted.LoadCRCD(s.Verification, privateKey)
			data.CRCD = dataCRCD
		//||------------------------------------------------------------------------------------------------||
		//|| Facial
		//||------------------------------------------------------------------------------------------------||
		case types.DataTypeFACE:
			dataFACE, _ := encrypted.LoadFACE(s.Verification, privateKey)
			data.FACE = dataFACE
		//||------------------------------------------------------------------------------------------------||
		//|| Identification
		//||------------------------------------------------------------------------------------------------||
		case types.DataTypeIDEN:
			dataID, _ := encrypted.LoadIDEN(s.Verification, privateKey)
			data.IDEN = dataID
		//||------------------------------------------------------------------------------------------------||
		//|| Email
		//||------------------------------------------------------------------------------------------------||
		case types.DataTypeMAIL:
			dataEMAL, _ := encrypted.LoadMAIL(s.Verification, privateKey)
			data.MAIL = dataEMAL
		//||------------------------------------------------------------------------------------------------||
		//|| Phone
		//||------------------------------------------------------------------------------------------------||
		case types.DataTypePHNE:
			dataPHON, _ := encrypted.LoadPHNE(s.Verification, privateKey)
			data.PHNE = dataPHON
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Shared Data
	//||------------------------------------------------------------------------------------------------||

	shared := OAuthShared{
		Token:     token,
		AccountId: enforcement.User.ID,
		ClientId:  enforcement.Site.ClientId,
		Scope:     enforcement.Scopes.ToList(),
		State:     enforcement.State,
		Age:       enforcement.Age,
		Data:      data,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create The Shared Wrapper
	//||------------------------------------------------------------------------------------------------||

	wrapper := OAuthSharedAccess{
		Token:     token,
		Shared:    shared,
		Status:    "PEND",
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Store to the Database
	//||------------------------------------------------------------------------------------------------||

	err := wrapper.Store()
	if err != nil {
		return OAuthSharedAccess{}, fmt.Errorf("failed to store shared access: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| All Done
	//||------------------------------------------------------------------------------------------------||

	return wrapper, nil
}
