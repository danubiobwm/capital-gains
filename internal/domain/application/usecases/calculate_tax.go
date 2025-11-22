package usecases

import (
	"github.com/danubiobwm/capital-gains/internal/domain/entities"
	"github.com/danubiobwm/capital-gains/internal/domain/repositories"
)

type CalculateTaxUseCase struct {
	taxRepo repositories.TaxRepository
}

func NewCalculateTaxUseCase(taxRepo repositories.TaxRepository) *CalculateTaxUseCase {
	return &CalculateTaxUseCase{taxRepo: taxRepo}
}

func (uc *CalculateTaxUseCase) Execute(operations []entities.Operation) []entities.TaxResult {
	results := make([]entities.TaxResult, len(operations))

	for i, op := range operations {
		tax := uc.taxRepo.CalculateOperationTax(op)
		results[i] = entities.TaxResult{Tax: tax}
	}

	return results
}
