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

	var requestBody struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	}


	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}


	cp_c.useCase.Execute(requestBody.Name, requestBody.Price)


	c.JSON(http.StatusOK, gin.H{"message": "Sabrita creada exitosamente"})
}