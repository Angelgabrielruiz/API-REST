package infraestructure

import (
	
	"demo/src/refrescos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteRefrescoController struct {
	useCase application.DeleteRefresco
}

func NewDeleteRefrescoController(useCase application.DeleteRefresco) *DeleteRefrescoController {
	return &DeleteRefrescoController{useCase: useCase}
}

func (dr_c *DeleteRefrescoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	dr_c.useCase.Execute(id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
