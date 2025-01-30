package infraestructure

import (
	"demo/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase application.DeleteProduct
}

func NewDeleteProductController(useCase application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{useCase: useCase}
}

func (dp_c *DeleteProductController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	dp_c.useCase.Execute(id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}