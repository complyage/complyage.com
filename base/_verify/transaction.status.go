package verify

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| TransactionStatus (iota-based)
//||------------------------------------------------------------------------------------------------||

type TransactionStatus int

const (
	TransactionPending TransactionStatus = iota
	TransactionProcessing
	TransactionApproved
	TransactionRefunded
	TransactionChargeback
)

//||------------------------------------------------------------------------------------------------||
//|| String
//||------------------------------------------------------------------------------------------------||

func (s TransactionStatus) String() string {
	switch s {
	case TransactionPending:
		return "PENDING"
	case TransactionProcessing:
		return "PROCESSING"
	case TransactionApproved:
		return "APPROVED"
	case TransactionRefunded:
		return "REFUNDED"
	case TransactionChargeback:
		return "CHARGEBACK"
	default:
		return "UNKNOWN"
	}
}

//||------------------------------------------------------------------------------------------------||
//|| JSON Marshal/Unmarshal
//||------------------------------------------------------------------------------------------------||

func (s TransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TransactionStatus) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	switch val {
	case "PENDING":
		*s = TransactionPending
	case "PROCESSING":
		*s = TransactionProcessing
	case "APPROVED":
		*s = TransactionApproved
	case "REFUNDED":
		*s = TransactionRefunded
	case "CHARGEBACK":
		*s = TransactionChargeback
	default:
		return fmt.Errorf("invalid TransactionStatus: %q", val)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Namespace for dot notation
//||------------------------------------------------------------------------------------------------||

type nsTransactionStatus struct {
	Pending    TransactionStatus
	Processing TransactionStatus
	Approved   TransactionStatus
	Refunded   TransactionStatus
	Chargeback TransactionStatus
}

var TRANSACTION_STATUS = nsTransactionStatus{
	Pending:    TransactionPending,
	Processing: TransactionProcessing,
	Approved:   TransactionApproved,
	Refunded:   TransactionRefunded,
	Chargeback: TransactionChargeback,
}
