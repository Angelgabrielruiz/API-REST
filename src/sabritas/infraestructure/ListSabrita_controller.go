package infraestructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo/src/sabritas/application"
)

type ListSabritaController struct {
	useCase application.ListSabrita
}

func NewListSabritaController(useCase application.ListSabrita) *ListSabritaController {
	return &ListSabritaController{useCase: useCase}
}

func (lp_c *ListSabritaController) Execute(c *gin.Context) {
	refrescos, err := lp_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los refrescos"})
		return
	}

	c.JSON(http.StatusOK, refrescos)
}