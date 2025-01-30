package infraestructure

import (
	"demo/src/products/application"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Configurar las dependencias
	ps := NewMySQL()

	// Crear los controladores
	createProductController := NewCreateProductController(*application.NewCreateProduct(ps))
	listProductController := NewListProductController(*application.NewListProduct(ps))
	getProductByIDController := NewGetProductByIDController(*application.NewGetProductByID(ps))
	updateProductController := NewUpdateProductController(*application.NewUpdateProduct(ps))
	deleteProductController := NewDeleteProductController(*application.NewDeleteProduct(ps))

	// Definir las rutas
	r.POST("/products", createProductController.Execute)
	r.GET("/products", listProductController.Execute)
	r.GET("/products/:id", getProductByIDController.Execute)
	r.PUT("/products/:id", updateProductController.Execute)
	r.DELETE("/products/:id", deleteProductController.Execute)
}