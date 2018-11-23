// Copyright (c) 2018 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/retr0h/bogus/api/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite
}

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName(".bogus")
	viper.AddConfigPath("..")
	viper.ReadInConfig()
}

func (suite *APITestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
}

func (suite *APITestSuite) TearDownTest() {
}

func (suite *APITestSuite) performRequest(method, target string, router *gin.Engine) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	return w
}

func (suite *APITestSuite) TestGetBusiness() {
	id := "3f01497d-90f4-41b0-5717-baa90cf8158b"
	path := fmt.Sprintf("/v1/business/%s", id)
	resp := suite.performRequest("GET", path, GetMainEngine())

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Equal(suite.T(), "application/json; charset=utf-8", resp.Header().Get("Content-Type"))

	got := &models.Business{}
	json.Unmarshal(resp.Body.Bytes(), &got)
	want := models.NewBusiness(id, "Business Name")
	assert.Equal(suite.T(), want, got)
}

func (suite *APITestSuite) TestGetBusinesses() {
	path := fmt.Sprintf("/v1/businesses")
	resp := suite.performRequest("GET", path, GetMainEngine())

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Equal(suite.T(), "application/json; charset=utf-8", resp.Header().Get("Content-Type"))

	got := models.Businesses{}
	json.Unmarshal(resp.Body.Bytes(), &got)

	var want models.Businesses

	id1 := "3f01497d-90f4-41b0-5717-baa90cf8158b"
	b1 := models.NewBusiness(id1, "Business Name 1")

	id2 := "3f01497d-90f4-41b0-5717-baa90cf8158b"
	b2 := models.NewBusiness(id2, "Business Name 2")

	want.Items = []models.Business{}
	want.Items = append(want.Items, *b1)
	want.Items = append(want.Items, *b2)

	assert.Equal(suite.T(), want, got)
}

// In order for `go test` to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestAPITestSuiteEntry(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
