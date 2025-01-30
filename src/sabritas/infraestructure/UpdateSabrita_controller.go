package infraestructure

import (
	"demo/src/sabritas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatesabritaController struct {
	useCase application.UpdateSabrita
}

func NewUpdateSabritaController(useCase application.UpdateSabrita) *UpdatesabritaController {
	return &UpdatesabritaController{useCase: useCase}
}

func (us_c *UpdatesabritaController) Execute(c *gin.Context) {
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

	us_c.useCase.Execute(id, input.Name, input.Price)
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}