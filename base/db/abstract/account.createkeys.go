package abstract

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/complyage/base/db/models"
	"github.com/complyage/base/send"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/bip39"
	"github.com/ralphferrara/aria/base/encrypt"
	"github.com/ralphferrara/aria/base/validate"
)

//||------------------------------------------------------------------------------------------------||
//|| Secrets
//||------------------------------------------------------------------------------------------------||

type AccountSecrets struct {
	EncryptionLevel int
	PrivateKey      string
	PublicKey       string
	CheckKey        string
}

//||------------------------------------------------------------------------------------------------||
//|| Create
//||------------------------------------------------------------------------------------------------||

func CreateAccountSecrets(r *http.Request, accountId int64, identifier string) (AccountSecrets, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	rawEncrypt := r.FormValue("encryptionLevel")
	privateKeyInput := r.FormValue("privateKey")
	publicKeyInput := r.FormValue("publicKey")
	wordListJSON := r.FormValue("wordList")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Encryption Level
	//||------------------------------------------------------------------------------------------------||

	encryptionLevel, err := strconv.Atoi(rawEncrypt)
	if err != nil || encryptionLevel < 1 || encryptionLevel > 3 {
		return AccountSecrets{}, app.Err("Auth").Error("INVALID_ACCOUNT_TYPE")
	}

	app.Log.Info("Encryption Level = ", encryptionLevel)

	//||------------------------------------------------------------------------------------------------||
	//|| Unmarshal the BIP39 Word List
	//||------------------------------------------------------------------------------------------------||

	var wordList bip39.BIP39WordList
	err = json.Unmarshal([]byte(wordListJSON), &wordList)
	if err != nil && encryptionLevel >= 2 && encryptionLevel <= 5 {
		return AccountSecrets{}, app.Err("Auth").Error("BIP39_GEN_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate and Generate Private/Public Key
	//||------------------------------------------------------------------------------------------------||

	var privateKey, publicKey string

	//||------------------------------------------------------------------------------------------------||
	//|| Level 1 - We handle the keys
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel == 1 {
		genPrivateKey, genPublicKey, err := encrypt.GenerateKeyPair()
		if err != nil {
			return AccountSecrets{}, app.Err("Auth").Error("PRIVPUB_FAILED")
		}
		privateKey = genPrivateKey
		publicKey = genPublicKey
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Level 2 - BIPList
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel >= 2 && encryptionLevel <= 5 {
		err := bip39.ValidateBIP39List(wordList, KeyLevelToWordCount(encryptionLevel))
		if err != nil {
			return AccountSecrets{}, app.Err("Auth").Error("BAD_BIP39")
		}
		genPrivate, genPublic, err := bip39.GenerateBIP39Keys(wordList)
		if err != nil {
			return AccountSecrets{}, app.Err("Auth").Error("BIP39_GEN_FAILED")
		}
		privateKey = genPrivate
		publicKey = genPublic
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Level 3 requires both keys
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel == 3 {
		privateKey = privateKeyInput
		publicKey = publicKeyInput
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Key Pair
	//||------------------------------------------------------------------------------------------------||

	app.Log.Info("Validating Key Pair")
	err = validate.ValidateKeyPair(privateKey, publicKey)
	if err != nil {
		app.Log.Info("Error validating key pair:", err)
		return AccountSecrets{}, app.Err("Auth").Error("PRIVPUB_MISMATCH")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Private Key Hash
	//||------------------------------------------------------------------------------------------------||

	privateKeyHash, err := encrypt.GenerateCheckKey(privateKey)
	if err != nil {
		app.Log.Info("Error generating private key hash:", err)
		return AccountSecrets{}, app.Err("Auth").Error("PRIVPUB_CHECKEY_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Contact the User with the keys if needed
	//||------------------------------------------------------------------------------------------------||

	app.Log.Info("Sending keys to user")
	if encryptionLevel == 1 {
		_ = send.EmailPrivateKeyToUser(identifier, privateKey)
	}

	if encryptionLevel == 2 {
		_ = send.EmailBIPListToUser(identifier, wordList, privateKey)
	}

	if encryptionLevel == 3 {
		_ = send.EmailPrivateKeyToUser(identifier, privateKey)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Add Key to DB
	//||------------------------------------------------------------------------------------------------||

	app.Log.Info("Creating Encryption Type")
	if encryptionLevel == 1 {
		CreateKey(&models.ModelSecrets{
			FidAccount: uint(accountId),
			Level:      encryptionLevel,
			Private:    privateKey,
			Public:     publicKey,
			CheckKey:   privateKeyHash,
			CreatedAt:  time.Now().UTC(),
		})
	} else {
		CreateKey(&models.ModelSecrets{
			FidAccount: uint(accountId),
			Level:      encryptionLevel,
			Private:    "",
			Public:     publicKey,
			CheckKey:   privateKeyHash,
			CreatedAt:  time.Now().UTC(),
		})
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return
	//||------------------------------------------------------------------------------------------------||

	return AccountSecrets{
		EncryptionLevel: encryptionLevel,
		PrivateKey:      privateKey,
		PublicKey:       publicKey,
		CheckKey:        privateKeyHash,
	}, nil

}
