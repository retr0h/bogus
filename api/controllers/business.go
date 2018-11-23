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

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/retr0h/bogus/api/models"
)

// BusinessController struct.
type BusinessController struct {
}

// NewBusinessController constructs a new `BusinessController`.
func NewBusinessController() (*BusinessController, error) {
	return &BusinessController{}, nil
}

// GetBusiness returns the business identified by `id`.
func (bc BusinessController) GetBusiness(c *gin.Context) {
	id := c.Param("id")
	nb := models.NewBusiness(id, "Business Name")

	c.JSON(http.StatusOK, nb)
}

// GetBusinesses returns all businesses.
func (bc BusinessController) GetBusinesses(c *gin.Context) {
	var businesses models.Businesses

	id1 := "3f01497d-90f4-41b0-5717-baa90cf8158b"
	b1 := models.NewBusiness(id1, "Business Name 1")

	id2 := "3f01497d-90f4-41b0-5717-baa90cf8158b"
	b2 := models.NewBusiness(id2, "Business Name 2")

	businesses.Items = []models.Business{}
	businesses.Items = append(businesses.Items, *b1)
	businesses.Items = append(businesses.Items, *b2)

	c.JSON(http.StatusOK, businesses)
}
