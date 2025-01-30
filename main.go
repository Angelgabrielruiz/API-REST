package main

import (
    "log"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    refrescosInfra "demo/src/refrescos/infraestructure" 
    sabritasInfra "demo/src/sabritas/infraestructure"
)

func main() {
    // Cargar el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    // Inicializar Gin
    r := gin.Default()

    // Configurar las rutas para refrescos
    refrescosInfra.SetupRoutesRefresco(r)

    // Configurar las rutas para sabritas
    sabritasInfra.SetupRoutesSabrita(r)

    // Iniciar el servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}