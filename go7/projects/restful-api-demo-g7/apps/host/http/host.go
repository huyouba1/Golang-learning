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

func (H *Handler) queryHost(c *gin.Context) {
	// 从http请求的query string中获取参数
	req := host.NewQueryHostFromHTTP(c.Request)

	// 进行接口调用，返回肯定有成功或者失败
	set, err := H.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

func (H *Handler) DescribeHost(c *gin.Context) {
	// 从http请求的query string中获取参数
	req := host.NewDescribeHostRequestWithId(c.Param("id"))

	// 进行接口调用，返回肯定有成功或者失败
	set, err := H.svc.DescribeHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

func (H *Handler) PutHost(c *gin.Context) {
	// 从http请求的query string中获取参数
	req := host.NewPutUpdateHostRequest(c.Param("id"))

	// 接触body里面数据
	if err := c.Bind(req.Host); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	req.Id = c.Param("id")
	// 进行接口调用，返回肯定有成功或者失败
	set, err := H.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

func (H *Handler) PatchHost(c *gin.Context) {
	// 从http请求的query string中获取参数
	req := host.NewPatchUpdateHostRequest(c.Param("id"))

	// 接触body里面数据
	if err := c.Bind(req.Host); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	req.Id = c.Param("id")

	// 进行接口调用，返回肯定有成功或者失败
	set, err := H.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

func (H *Handler) DeleteHost(c *gin.Context) {
	req := host.NewDeleteHostRequest(c.Param("id"))

	if err := c.Bind(req.Id); err != nil {
		response.Failed(c.Writer, err)
		return
	}

	// 接口调用
	ins, err := H.svc.DeleteHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, ins)

}
