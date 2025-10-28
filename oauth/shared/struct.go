package shared

import (
	"time"

	"github.com/complyage/base/enforce"
	"github.com/complyage/base/verify"
)

type OAuthSharedAccess struct {
	Token     string      `json:"token"`
	ExpiresAt time.Time   `json:"expires_at"`
	Status    string      `json:"status"`
	Shared    OAuthShared `json:"shared"`
}

type OAuthShared struct {
	Token     string      `json:"token"`
	AccountId int64       `json:"account_id"`
	ClientId  string      `json:"client_id"`
	Scope     []string    `json:"scope"`
	State     string      `json:"state"`
	Age       enforce.Age `json:"age"`
	Data      verify.Data `json:"data"`
}

type OAuthVerification struct {
	Type     string `json:"type,omitempty"`
	Verified bool   `json:"verified,omitempty"`
}
