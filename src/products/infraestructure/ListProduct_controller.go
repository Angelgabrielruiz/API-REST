package infraestructure

import (
	"demo/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListProductController struct {
	useCase application.ListProduct
}

func NewListProductController(useCase application.ListProduct) *ListProductController {
	return &ListProductController{useCase: useCase}
}

func (lp_c *ListProductController) Execute(c *gin.Context) {
	lp_c.useCase.Execute()
	c.JSON(http.StatusOK, gin.H{"message": "Products listed successfully"})
}