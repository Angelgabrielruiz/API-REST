package infraestructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"demo/src/refrescos/application"
)

type GetRefrescoByIDController struct {
	useCase application.GetRefrescoByID
}

func NewGetRefrescoByIDController(useCase application.GetRefrescoByID) *GetRefrescoByIDController {
	return &GetRefrescoByIDController{useCase: useCase}
}

func (gp_c *GetRefrescoByIDController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	refresco, err := gp_c.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el refresco"})
		return
	}

	if refresco == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Refresco no encontrado"})
		return
	}

	c.JSON(http.StatusOK, refresco)
}