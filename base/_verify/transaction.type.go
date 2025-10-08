package verify

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| TransactionType (iota-based)
//||------------------------------------------------------------------------------------------------||

type TransactionType int

const (
	TransactionTypeCash TransactionType = iota
	TransactionTypeCrypto
	TransactionTypePaypal
	TransactionTypeCredit
)

//||------------------------------------------------------------------------------------------------||
//|| String
//||------------------------------------------------------------------------------------------------||

func (t TransactionType) String() string {
	switch t {
	case TransactionTypeCash:
		return "CASH"
	case TransactionTypeCrypto:
		return "CRYPTO"
	case TransactionTypePaypal:
		return "PAYPAL"
	case TransactionTypeCredit:
		return "CREDIT"
	default:
		return "UNKNOWN"
	}
}

//||------------------------------------------------------------------------------------------------||
//|| JSON Marshal/Unmarshal
//||------------------------------------------------------------------------------------------------||

func (t TransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TransactionType) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	switch val {
	case "CASH":
		*t = TransactionTypeCash
	case "CRYPTO":
		*t = TransactionTypeCrypto
	case "PAYPAL":
		*t = TransactionTypePaypal
	case "CREDIT":
		*t = TransactionTypeCredit
	default:
		return fmt.Errorf("invalid TransactionType: %q", val)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Namespace for dot notation
//||------------------------------------------------------------------------------------------------||

type nsTransactionType struct {
	Cash   TransactionType
	Crypto TransactionType
	Paypal TransactionType
	Credit TransactionType
}

var TRANSACTION_TYPE = nsTransactionType{
	Cash:   TransactionTypeCash,
	Crypto: TransactionTypeCrypto,
	Paypal: TransactionTypePaypal,
	Credit: TransactionTypeCredit,
}
