package cli

import (
	"bufio"

	"fmt"
	"os"

	"github.com/danubiobwm/capital-gains/internal/domain/infrastructure/json"
	"github.com/danubiobwm/capital-gains/internal/domain/infrastructure/repositories"
	"github.com/danubiobwm/capital-gains/internal/domain/services"
)

type CLIHandler struct {
	jsonParser *json.Parser
	taxService *services.TaxService
}

func NewCLIHandler() *CLIHandler {
	taxRepo := repositories.NewMemoryTaxRepository()
	taxService := services.NewTaxService(taxRepo)

	return &CLIHandler{
		jsonParser: json.NewParser(),
		taxService: taxService,
	}
}

func (h *CLIHandler) Handle() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		operations, err := h.jsonParser.ParseOperations(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
			continue
		}

		results := h.taxService.CalculateTaxes(operations)

		output, err := h.jsonParser.FormatTaxResults(results)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
			continue
		}

		fmt.Println(output)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
