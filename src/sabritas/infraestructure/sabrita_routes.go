package infraestructure

import (
	"demo/src/sabritas/application"

	"github.com/gin-gonic/gin"
)

func SetupRoutesSabrita(r *gin.Engine) {
	// Configurar las dependencias
	ps := NewMySQLSabrita()

	// Crear los controladores
	createSabritaController := NewCreateSabritaController(*application.NewCreateSabrita(ps))
	listSabritaController := NewListSabritaController(*application.NewListSabrita(ps))
	getSabritaByIDController := NewGetSabritaByIDController(*application.NewGetSabritaByID(ps))
	updateSabritaController := NewUpdateSabritaController(*application.NewUpdateSabrita(ps))
	deleteSabritaController := NewDeleteSabritaController(*application.NewDeleteSabrita(ps))

	// Definir las rutas
	r.POST("/sabritas", createSabritaController.Execute)
	r.GET("sabritas", listSabritaController.Execute)
	r.GET("/sabritas/:id", getSabritaByIDController.Execute)
	r.PUT("/sabritas/:id", updateSabritaController.Execute)
	r.DELETE("/sabritas/:id", deleteSabritaController.Execute)
}