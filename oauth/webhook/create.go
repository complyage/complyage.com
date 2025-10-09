package webhook

import (
	"time"

	"github.com/complyage/base/enforce"
	"github.com/complyage/base/keeper"
	"github.com/google/uuid"
)

func Create(e enforce.Enforcement, k keeper.KeeperRecord, scope string) WebhookPayload {
	hook := WebhookPayload{
		ID:         uuid.New().String(),
		ClientID:   e.Site.ClientId,
		CheckToken: e.Site.Verifier,
		SessionID:  k.KeeperId,
		Username:   e.User.Username,
		IPAddress:  e.Zone.IPAddress,
		Region:     e.Zone.Region,
		Country:    e.Zone.Country,
		DOB:        e.Age.DOB.String(),
		Age:        e.Age.DOB.Age(),
		Method:     scope,
		Timestamp:  time.Now(),
	}
	return hook
}
