package keeper

import "github.com/ralphferrara/aria/cache"

var GATE_COOKIE_NAME = "complyage_session"
var GATE_VERSION = "1.0.1"
var SessionCache *cache.RedisCacheWrapper

func Init(s *cache.RedisCacheWrapper) {
	SessionCache = s
}
