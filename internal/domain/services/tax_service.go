package services

import (
	"github.com/danubiobwm/capital-gains/internal/domain/entities"
	"github.com/danubiobwm/capital-gains/internal/domain/repositories"
)

type TaxService struct {
	repo repositories.TaxRepository
}

func NewTaxService(repo repositories.TaxRepository) *TaxService {
	return &TaxService{repo: repo}
}

func (s *TaxService) CalculateTaxes(operations []entities.Operation) []entities.TaxResult {
	s.repo.ResetState()

	results := make([]entities.TaxResult, len(operations))

	for i, op := range operations {
		tax := s.repo.CalculateOperationTax(op)
		results[i] = entities.TaxResult{Tax: tax}
	}

	return results
}
