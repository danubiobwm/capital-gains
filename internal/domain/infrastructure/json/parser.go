package json

import (
	"encoding/json"

	"github.com/danubiobwm/capital-gains/internal/domain/entities"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseOperations(input string) ([]entities.Operation, error) {
	var operations []entities.Operation
	err := json.Unmarshal([]byte(input), &operations)
	if err != nil {
		return nil, err
	}
	return operations, nil
}

func (p *Parser) FormatTaxResults(results []entities.TaxResult) (string, error) {
	jsonData, err := json.Marshal(results)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
