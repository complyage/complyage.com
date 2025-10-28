package encrypted

import (
	"github.com/ralphferrara/aria/storage"
)

var Storage *storage.Storage

func Init(s *storage.Storage) {
	if s == nil {
		panic("Init called with nil storage")
	}
	Storage = s
}
