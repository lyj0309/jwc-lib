package wxLib

import (
	"github.com/lyj0309/jwc-lib/lib"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func newRedisCache() *cache.Redis {
	redisOpts := &cache.RedisOpts{
		Host:        lib.Config.RedisAddr,
		Database:    0,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60, //second
	}
	return cache.NewRedis(redisOpts)
}

func NewWxMini() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	mini := wc.GetMiniProgram(&miniConfig.Config{
		AppID:     lib.Config.MiniAppId,
		AppSecret: lib.Config.MiniAppSecret,
		Cache:     newRedisCache(),
	})
	return mini
}

func NewOfficial() *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()

	official := wc.GetOfficialAccount(&offConfig.Config{
		AppID:          lib.Config.OffAppId,
		AppSecret:      lib.Config.OffAppSecret,
		Token:          lib.Config.OffToken,
		EncodingAESKey: lib.Config.OffEncodingAESKey,
		Cache:          newRedisCache(),
	})
	return official
}
