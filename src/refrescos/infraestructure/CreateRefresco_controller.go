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
	// Definir una estructura para recibir el cuerpo de la solicitud
	var requestBody struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}

	// Bind JSON: Obtener name y price desde el cuerpo de la solicitud
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Ejecutar el caso de uso con los valores proporcionados
	cp_c.useCase.Execute(requestBody.Name, requestBody.Price)

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Refresco creado exitosamente"})
}