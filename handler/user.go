package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/plancer16/mall_end/enum"
	"github.com/plancer16/mall_end/model"
	"github.com/plancer16/mall_end/resp"
	"github.com/plancer16/mall_end/service"
	"net/http"
)

type UserHandler struct {
	UserService service.UserService
}

func (h *UserHandler) AddUserHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OpreateFail),
		Msg:   enum.OpreateFail.String(),
		Total: 0,
		Data:  nil,
	}
	//var user model.User（之前的变量不会改变）
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err!= nil {
		entity.Msg="绑定错误"
		c.JSON(http.StatusOK,gin.H{"entity":entity})
		return
	}
	u, err := h.UserService.Add(user)
	if err != nil || u.UserId == "" {
		entity.Msg = err.Error()
		c.JSON(http.StatusOK,gin.H{"entity":entity})
		return
	}
	entity.Code = http.StatusOK
	entity.Msg  = enum.OperateOk.String()
	c.JSON(http.StatusOK,gin.H{"entity":entity})
}