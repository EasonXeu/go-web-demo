package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-web-demo/src/service/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserController_CreateUser(t *testing.T) {
	mockUserService := mocks.UserService{}
	mockUserService.EXPECT().CreateUser(mock.Anything).Return(nil)

	userController := UserController{UserService: &mockUserService}
	w := performMockRequest(
		userController.CreateUser,
		map[string]string{},
		map[string]string{},
		"")
	assert.Equal(t, InvalidParam.httpCode, w.Code)
	httpResult := HttpResult{}
	json.Unmarshal(w.Body.Bytes(), &httpResult)
	assert.Equal(t, InvalidParam.errorCode, httpResult.Code)
	assert.NotNil(t, httpResult.ErrorMessage)
}

func performMockRequest(handler gin.HandlerFunc, params, headers map[string]string, requestBodyString string) *httptest.ResponseRecorder {
	// create a gin mock environment
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/test", handler)
	// create a gin mock request
	if len(requestBodyString) == 0 {
		c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	} else {
		requestBody := strings.NewReader(requestBodyString)
		c.Request, _ = http.NewRequest(http.MethodGet, "/test", requestBody)
	}
	parameters := []gin.Param{}
	for k, v := range params {
		parameters = append(parameters, gin.Param{
			Key:   k,
			Value: v,
		})
	}
	c.Params = parameters
	for k, v := range headers {
		c.Request.Header.Add(k, v)
	}
	// perform the mock request
	r.ServeHTTP(w, c.Request)
	return w
}
