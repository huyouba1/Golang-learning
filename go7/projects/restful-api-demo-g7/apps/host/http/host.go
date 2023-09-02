package http

import (
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

// 用于暴露host service的接口
// H handler中，调用了host.Service接口
func (H *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	// 用户传递过来的参数进行解析
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		//c.Data(http.StatusBadRequest, "application/json")
		//c.Writer.Write()
		return
	}

	// 进行接口调用
	//gin.HandlerFunc()
	ins, err := H.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, ins)
}
