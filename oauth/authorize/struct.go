package authorize

import (
	"time"

	"github.com/complyage/base/types"
)

//||------------------------------------------------------------------------------------------------||
//|| Authorization Session
//||------------------------------------------------------------------------------------------------||

type OAuthSession struct {
	SessionId      string    `json:"session_id"`
	ClientId       string    `json:"client_id"`
	RedirectURI    string    `json:"redirect_uri"`
	Scope          string    `json:"scope"`
	State          string    `json:"state"`
	ExternalUserId string    `json:"external_user_id"`
	Salt           string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Approved       bool      `json:"approved"`
	Authorized     bool      `json:"authorized"`
	Retrieved      bool      `json:"retrieved"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserInfo Payload
//||------------------------------------------------------------------------------------------------||

type OAuthPayload struct {
	AccountId  int64       `json:"account_id"`
	Returned   []string    `json:"returned"`
	Face       AuditRecord `json:"face"`
	IDCard     AuditRecord `json:"id_card"`
	Address    AuditRecord `json:"address"`
	Email      AuditRecord `json:"email"`
	Phone      AuditRecord `json:"phone"`
	CreditCard AuditRecord `json:"credit_card"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgePayload
//||------------------------------------------------------------------------------------------------||

type AgePayload struct {
	AccountId int64       `json:"account_id,omitempty"`
	Age       int         `json:"age,omitempty"`
	DOB       types.DOB   `json:"dob,omitempty"`
	Audit     AuditRecord `json:"audit_record,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Audit Record
//||------------------------------------------------------------------------------------------------||

type AuditRecord struct {
	Data       string `json:"data,omitempty"`
	AuditId    string `json:"audit_id"`
	AuditType  string `json:"type"`
	Submitted  string `json:"submitted"`
	VerifiedBy string `json:"verified_by"`
	Timestamp  string `json:"timestamp"`
}
