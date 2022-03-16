package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/plancer16/mall_end/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserRegister(t *testing.T) {
	router := gin.Default()
	router.POST("/api/user/add", UserHandler.AddUserHandler)//服务器增加处理器
	user := model.User{
		UserId:    "1285",
		NickName:  "eqeq",
		Mobile:    "18751803321",
		Password:  "qwe",
		Address:   "xuanwu",
		IsDeleted: false,
		IsLocked:  false,
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
	}
	marshal, _ := json.Marshal(user)
	s := string(marshal)
	buffer := bytes.NewBufferString(s)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/user/add", buffer)
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w,req)
	assert.Equal(t,w.Code,http.StatusOK)
}