package access

import (
	"time"

	"github.com/complyage/base/enforce"

	"github.com/ralphferrara/aria/base/random"
)

//||------------------------------------------------------------------------------------------------||
//|| Create
//||------------------------------------------------------------------------------------------------||

func Create(enforcement enforce.Enforcement) (OAuthAccess, error) {
	token := random.RandomString(32)
	oa := OAuthAccess{
		Token:          token,
		ExpiresAt:      time.Now().Add(10 * time.Minute),
		Enforcement:    enforcement,
		Approved:       false,
		PreviewKey:     random.RandomString(16),
		DenyKey:        random.RandomString(16),
		BypassKey:      random.RandomString(16),
		AuthorizeKey:   token + "_" + random.RandomString(16),
		AuthorizeToken: token + "_" + random.RandomString(16),
	}
	err := oa.Store()
	if err != nil {
		return OAuthAccess{}, err
	}
	return oa, nil
}
