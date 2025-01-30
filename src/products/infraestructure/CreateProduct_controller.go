package infraestructure

import (
	"demo/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCase application.CreateProduct
}

func NewCreateProductController(useCase application.CreateProduct) *CreateProductController {
	return &CreateProductController{useCase: useCase}
}

func (cp_c *CreateProductController) Execute(c *gin.Context) {
	cp_c.useCase.Execute()
	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}