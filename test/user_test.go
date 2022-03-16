package test

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
	user := model.User{
		UserId:    "1285",
		NickName:  "qweowq",
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
	router.ServeHTTP(w,req)
	assert.Equal(t,w.Code,http.StatusOK)
}