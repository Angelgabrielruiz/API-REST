package application

import (
    "demo/src/sabritas/domain"
)

type FilterSabritasByNameAndPrice struct {
    sabritaRepo domain.ISabritaRepository
}


func NewFilterSabritasByNameAndPrice(sabritaRepo domain.ISabritaRepository) *FilterSabritasByNameAndPrice {
    return &FilterSabritasByNameAndPrice{
        sabritaRepo: sabritaRepo,
    }
}


func (f *FilterSabritasByNameAndPrice) FilterByNameAndPrice(name string, price float64) ([]map[string]interface{}, error) {
    return f.sabritaRepo.FilterByNameAndPrice(name, price)
}