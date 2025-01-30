package infraestructure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"demo/src/sabritas/application"
)

type CreateSabritaController struct {
	useCase application.CreateSabrita
}

func NewCreateSabritaController(useCase application.CreateSabrita) *CreateSabritaController {
	return &CreateSabritaController{useCase: useCase}
}

func (cp_c *CreateSabritaController) Execute(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"message": "Sabrita creada exitosamente"})
}