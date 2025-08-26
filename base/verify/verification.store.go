package verify

import "time"

//||------------------------------------------------------------------------------------------------||
//|| Error Levels
//||------------------------------------------------------------------------------------------------||

type VerificationComplete struct {
	//||------------------------------------------------------------------------------------------------||
	//|| Error Levels
	//||------------------------------------------------------------------------------------------------||
	UUID string `json:"uuid"`
	//||------------------------------------------------------------------------------------------------||
	//|| DB Fields
	//||------------------------------------------------------------------------------------------------||
	Status  StatusType `json:"status"`
	Type    DataType   `json:"type"`
	Display string     `json:"display"`
	//||------------------------------------------------------------------------------------------------||
	//|| Process
	//||------------------------------------------------------------------------------------------------||
	Level int    `json:"level"`
	Step  int    `json:"step"`
	Steps []Step `json:"steps"`
	//||------------------------------------------------------------------------------------------------||
	//|| Approval
	//||------------------------------------------------------------------------------------------------||
	Moderate    Moderate          `json:"approval,omitempty"`
	Transaction TransactionPublic `json:"transaction,omitempty"`
	//||------------------------------------------------------------------------------------------------||
	//|| Process
	//||------------------------------------------------------------------------------------------------||
	Timestamp time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	Completed time.Time `json:"completed,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Error Levels
//||------------------------------------------------------------------------------------------------||

func (v *Verification) AsComplete() VerificationComplete {
	return VerificationComplete{
		UUID:        v.UUID,
		Status:      v.Status,
		Type:        v.Type,
		Display:     v.Display,
		Level:       v.Level,
		Step:        v.Step,
		Steps:       v.Steps,
		Moderate:    v.Moderate,
		Transaction: v.Transaction,
		Timestamp:   v.Timestamp,
		UpdatedAt:   v.UpdatedAt,
		Completed:   v.Completed,
	}
}
