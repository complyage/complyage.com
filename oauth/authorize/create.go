package authorize

import (
	"time"

	"github.com/ralphferrara/aria/base/random"
)

//||------------------------------------------------------------------------------------------------||
//|| Create the Authorize Session
//||------------------------------------------------------------------------------------------------||

func Create(clientId, redirectURI, scope, state, externalUserId string) (AuthorizeSession, error) {
	var session AuthorizeSession
	session.SessionId = time.Now().Format("20060102150405") + "_" + clientId + "_" + random.RandomString(16)
	session.ClientId = clientId
	session.RedirectURI = redirectURI
	session.Scope = scope
	session.State = state
	session.ExternalUserId = externalUserId
	session.Salt = random.RandomString(32)
	session.CreatedAt = time.Now()
	session.Approved = false
	session.Authorized = false
	session.Retrieved = false
	session.Payload = CreatePayload(0)
	return session, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Create the Payload
//||------------------------------------------------------------------------------------------------||

func CreatePayload(accountId int64) AuthorizePayload {
	var payload AuthorizePayload
	payload.AccountId = accountId
	payload.Audits = AuthorizeAudits{}
	return payload
}
