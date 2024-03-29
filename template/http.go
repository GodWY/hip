package template

import (
	"net/http"

	"github.com/GodWY/hip/service"
	"github.com/gin-gonic/gin"
)

var ts AckHttpService

type TemplateReq struct {
}

type TemplateRsp struct {
}

type AckHttpService interface {
	HandleAck(ctx *gin.Context, in *TemplateReq) (out *TemplateRsp, err error)
}

type AckService interface {
	HandleAck(ctx *gin.Context)
}

// RegisterHandler 注册服务
func registerHttpHandler(srv service.Service, srvs AckService) {
	// 注册服务组
	group := srv.Router("ack")
	group.GET("/", srvs.HandleAck)
}

type Template struct{}

// RegisterPbHttpHandler 注册pb服务
func RegisterPbHttpHandler(srv service.Service, srvs AckHttpService) {
	tmp := new(Template)
	ts = srvs
	registerHttpHandler(srv, tmp)
}

func (t *Template) HandleAck(ctx *gin.Context) {
	req := &TemplateReq{}
	if ok := ctx.Bind(req); ok != nil {
		ctx.JSON(http.StatusOK, "bind error")
		return
	}
	rsp, err := ts.HandleAck(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}
