package infraestructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo/src/refrescos/application"
)

type ListRefrescoController struct {
	useCase application.ListRefresco
}

func NewListRefrescoController(useCase application.ListRefresco) *ListRefrescoController {
	return &ListRefrescoController{useCase: useCase}
}

func (lp_c *ListRefrescoController) Execute(c *gin.Context) {
	refrescos, err := lp_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los refrescos"})
		return
	}

	c.JSON(http.StatusOK, refrescos)
}