package infraestructure

import (
	"demo/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductByIDController struct {
	useCase application.GetProductByID
}

func NewGetProductByIDController(useCase application.GetProductByID) *GetProductByIDController {
	return &GetProductByIDController{useCase: useCase}
}

func (gp_c *GetProductByIDController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	gp_c.useCase.Execute(id)
	c.JSON(http.StatusOK, gin.H{"message": "Product retrieved successfully"})
}