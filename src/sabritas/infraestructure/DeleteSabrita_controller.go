package infraestructure

import (
	
	"demo/src/sabritas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteSabritaController struct {
	useCase application.Deletesabrita
}

func NewDeleteSabritaController(useCase application.Deletesabrita) *DeleteSabritaController {
	return &DeleteSabritaController{useCase: useCase}
}

func (ds_c *DeleteSabritaController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	ds_c.useCase.Execute(id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
