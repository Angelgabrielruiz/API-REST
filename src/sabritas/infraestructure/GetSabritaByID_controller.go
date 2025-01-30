package infraestructure

import (
	"demo/src/sabritas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetSabritaByIDController struct {
	useCase application.GetSabritaByID
}

func NewGetSabritaByIDController(useCase application.GetSabritaByID) *GetSabritaByIDController {
	return &GetSabritaByIDController{useCase: useCase}
}

func (gp_c *GetSabritaByIDController) Execute(c *gin.Context) {
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
