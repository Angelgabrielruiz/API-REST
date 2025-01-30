package infraestructure

import (
	"demo/src/refrescos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateRefrescoController struct {
	useCase application.UpdateRefresco
}

func NewUpdateRefrescoController(useCase application.UpdateRefresco) *UpdateRefrescoController {
	return &UpdateRefrescoController{useCase: useCase}
}

func (up_c *UpdateRefrescoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var input struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	up_c.useCase.Execute(id, input.Name, input.Price)
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}