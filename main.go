package main

import (
    "log"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    refrescosInfra "demo/src/refrescos/infraestructure" 
    sabritasInfra "demo/src/sabritas/infraestructure"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    r := gin.Default()

    refrescosInfra.SetupRoutesRefresco(r)

    sabritasInfra.SetupRoutesSabrita(r)

    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}