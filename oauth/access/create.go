package access

import (
	"github.com/complyage/base/enforce"

	"github.com/ralphferrara/aria/base/random"
)

//||------------------------------------------------------------------------------------------------||
//|| Create
//||------------------------------------------------------------------------------------------------||

func Create(enforcement enforce.Enforcement) (OAuthAccess, error) {
	token := random.RandomString(32)
	oa := OAuthAccess{
		Token:        token,
		Enforcement:  enforcement,
		Approved:     false,
		PreviewKey:   random.RandomString(16),
		AuthorizeKey: random.RandomString(16),
		DenyKey:      random.RandomString(16),
		BypassKey:    random.RandomString(16),
	}
	err := oa.Store()
	if err != nil {
		return OAuthAccess{}, err
	}
	return oa, nil
}
