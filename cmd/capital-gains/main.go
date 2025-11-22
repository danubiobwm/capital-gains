package main

import "github.com/danubiobwm/capital-gains/internal/domain/infrastructure/cli"

func main() {
	handler := cli.NewCLIHandler()
	handler.Handle()
}
