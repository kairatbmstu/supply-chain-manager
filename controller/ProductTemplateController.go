package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductTemplateController struct {
}

func (p ProductTemplateController) GetCreateForm(c *gin.Context) {
	c.HTML(http.StatusOK, "product_template_create.html", gin.H{})
}

func (p ProductTemplateController) GetUpdateForm(c *gin.Context) {
	c.HTML(http.StatusOK, "product_template_update.html", gin.H{})
}

func (p ProductTemplateController) Create(c *gin.Context) {
	c.Redirect(http.StatusFound, "/product-templates/get-details")
}

func (p ProductTemplateController) Update(c *gin.Context) {
	c.Redirect(http.StatusFound, "/product-templates/get-details")
}

func (p ProductTemplateController) GetDetails(c *gin.Context) {
	c.HTML(http.StatusOK, "product_template_details.html", gin.H{})
}

func (p ProductTemplateController) GetProductTemplatesList(c *gin.Context) {
	c.HTML(http.StatusOK, "product_template_list.html", gin.H{})
}
