package officialAccount

import (
	"github.com/gin-gonic/gin"
	"github.com/gowechat/example/config"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	log "github.com/sirupsen/logrus"
)

//ExampleOfficialAccount 公众号操作样例
type ExampleOfficialAccount struct {
	wc              *wechat.Wechat
	officialAccount *officialaccount.OfficialAccount
}

//ExampleOfficialAccount new
func NewExampleOfficialAccount(wc *wechat.Wechat) *ExampleOfficialAccount {
	//init config
	globalCfg := config.GetConfig()
	offCfg := &offConfig.Config{
		AppID:          globalCfg.AppID,
		AppSecret:      globalCfg.AppSecret,
		Token:          globalCfg.Token,
		EncodingAESKey: globalCfg.EncodingAESKey,
	}
	log.Debugf("offCfg=%+v", offCfg)
	officialAccount := wc.GetOfficialAccount(offCfg)
	return &ExampleOfficialAccount{
		wc:              wc,
		officialAccount: officialAccount,
	}
}

//Serve 处理消息
func (ex *ExampleOfficialAccount) Serve(c *gin.Context) {
	// 传入request和responseWriter
	server := ex.officialAccount.GetServer(c.Request, c.Writer)
	server.SkipValidate(true)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeImage, MsgData: text}

		//article1 := message.NewArticle("测试图文1", "图文描述", "", "")
		//articles := []*message.Article{article1}
		//news := message.NewNews(articles)
		//return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}

		//voice := message.NewVoice(mediaID)
		//return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}

		//
		//video := message.NewVideo(mediaID, "标题", "描述")
		//return &message.Reply{MsgType: message.MsgTypeVideo, MsgData: video}

		//music := message.NewMusic("标题", "描述", "音乐链接", "HQMusicUrl", "缩略图的媒体id")
		//return &message.Reply{MsgType: message.MsgTypeMusic, MsgData: music}

		//多客服消息转发
		//transferCustomer := message.NewTransferCustomer("")
		//return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: transferCustomer}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		log.Error("Serve Error, err=%+v", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		log.Error("Send Error, err=%+v", err)
		return
	}
}
