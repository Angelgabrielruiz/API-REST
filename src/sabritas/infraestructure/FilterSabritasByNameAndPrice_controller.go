package infraestructure

import (
    "demo/src/sabritas/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type FilterSabritasByNameAndPriceController struct {
    filterService *application.FilterSabritasByNameAndPrice
}

func NewFilterSabritasByNameAndPriceController(filterService *application.FilterSabritasByNameAndPrice) *FilterSabritasByNameAndPriceController {
    return &FilterSabritasByNameAndPriceController{
        filterService: filterService,
    }
}

func (f *FilterSabritasByNameAndPriceController) Execute(ctx *gin.Context) {
    name := ctx.Param("name")
    minPriceStr := ctx.Param("price")

    price, err := strconv.ParseFloat(minPriceStr, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
        return
    }

    sabritas, err := f.filterService.FilterByNameAndPrice(name, price)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, sabritas)
}