package infraestructure

import (
	"demo/src/refrescos/application"

	"github.com/gin-gonic/gin"
)

func SetupRoutesRefresco(r *gin.Engine) {
	
	ps := NewMySQLRefresco()


	createRefrescoController := NewCreateRefrescoController(*application.NewCreateRefresco(ps))
	listRefrescoController := NewListRefrescoController(*application.NewListRefresco(ps))
	getRefrescoByIDController := NewGetRefrescoByIDController(*application.NewGetRefrescoByID(ps))
	updateRefrescoController := NewUpdateRefrescoController(*application.NewUpdateRefresco(ps))
	deleteRefrescoController := NewDeleteRefrescoController(*application.NewDeleteRefresco(ps))


	r.POST("/refrescos", createRefrescoController.Execute)
	r.GET("refrescos", listRefrescoController.Execute)
	r.GET("/refrescos/:id", getRefrescoByIDController.Execute)
	r.PUT("/refrescos/:id", updateRefrescoController.Execute)
	r.DELETE("/refrescos/:id", deleteRefrescoController.Execute)
}