package example

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/gowechat/example/config"
	exampleOffAcount "github.com/gowechat/example/pkg/officialAccount"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	log "github.com/sirupsen/logrus"
)

func init() {
	flag.Parse()
}

//Run 程序入口
func Run() error {
	log.Info("start wechat sdk example project")

	cfg := config.GetConfig()
	r := gin.Default()

	//获取wechat实例
	wc := InitWechat()

	//公众号例子相关操作
	exampleOffAccount := exampleOffAcount.NewExampleOfficialAccount(wc)
	//处理推送消息以及事件
	r.Any("/api/v1/serve", exampleOffAccount.Serve)
	//获取ak
	r.GET("/api/v1/oa/basic/get_access_token", exampleOffAccount.GetAccessToken)
	//获取微信callback IP
	r.GET("/api/v1/oa/basic/get_callback_ip", exampleOffAccount.GetCallbackIP)
	//获取微信API接口 IP
	r.GET("/api/v1/oa/basic/get_api_domain_ip", exampleOffAccount.GetAPIDomainIP)
	//清理接口调用次数
	r.GET("/api/v1/oa/basic/clear_quota", exampleOffAccount.ClearQuota)

	//获取

	//显示首页
	r.GET("/", Index)

	return r.Run(cfg.Listen)
}

//Index 显示首页
func Index(c *gin.Context) {
	c.JSON(200, "index")
}

//InitWechat 获取wechat实例
//在这里已经设置了全局cache，则在具体获取公众号/小程序等操作实例之后无需再设置，设置即覆盖
func InitWechat() *wechat.Wechat {
	cfg := config.GetConfig()
	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        cfg.Redis.Host,
		Password:    cfg.Redis.Password,
		Database:    cfg.Redis.Database,
		MaxActive:   cfg.Redis.MaxActive,
		MaxIdle:     cfg.Redis.MaxIdle,
		IdleTimeout: cfg.Redis.IdleTimeout,
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)
	return wc
}
