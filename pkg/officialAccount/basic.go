package officialAccount

import (
	"github.com/gin-gonic/gin"
	"github.com/gowechat/example/pkg/util"
	log "github.com/sirupsen/logrus"
)

//GetAccessToken 获取ak
func (ex *ExampleOfficialAccount) GetAccessToken(c *gin.Context) {
	ak, err := ex.officialAccount.GetAccessToken()
	if err != nil {
		log.Errorf("get ak error, err=%+v", err)
		util.RenderError(c, err)
		return
	}
	util.RenderSuccess(c, ak)
}

//GetCallbackIP 获取微信callback IP地址
func (ex *ExampleOfficialAccount) GetCallbackIP(c *gin.Context) {
	ipList, err := ex.officialAccount.GetBasic().GetCallbackIP()
	if err != nil {
		log.Errorf("GetCallbackIP error, err=%+v", err)
		util.RenderError(c, err)
		return
	}
	util.RenderSuccess(c, ipList)
}

//GetAPIDomainIP 获取微信callback IP地址
func (ex *ExampleOfficialAccount) GetAPIDomainIP(c *gin.Context) {
	ipList, err := ex.officialAccount.GetBasic().GetAPIDomainIP()
	if err != nil {
		log.Errorf("GetAPIDomainIP error, err=%+v", err)
		util.RenderError(c, err)
		return
	}
	util.RenderSuccess(c, ipList)
}

//GetAPIDomainIP  清理接口调用次数
func (ex *ExampleOfficialAccount) ClearQuota(c *gin.Context) {
	err := ex.officialAccount.GetBasic().ClearQuota()
	if err != nil {
		log.Errorf("ClearQuota error, err=%+v", err)
		util.RenderError(c, err)
		return
	}
	util.RenderSuccess(c, "success")
}
