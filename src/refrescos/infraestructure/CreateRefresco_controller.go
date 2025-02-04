package infraestructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo/src/refrescos/application"
)

type CreateRefrescoController struct {
	useCase application.CreateRefresco
}

func NewCreateRefrescoController(useCase application.CreateRefresco) *CreateRefrescoController {
	return &CreateRefrescoController{useCase: useCase}
}

func (cp_c *CreateRefrescoController) Execute(c *gin.Context) {
	
	var requestBody struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	
	cp_c.useCase.Execute(requestBody.Name, requestBody.Price)


	c.JSON(http.StatusOK, gin.H{"message": "Refresco creado exitosamente"})
}