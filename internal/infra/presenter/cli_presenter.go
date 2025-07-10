package presenter

import (
	"fmt"
	"stress-test/internal/entity"
)

// Implementação para apresentar o relatório na linha de comando.
type CliPresenter struct{}

func NewCliPresenter() *CliPresenter {
	return &CliPresenter{}
}

// Implementa a interface Presenter.
func (p *CliPresenter) Present(report entity.Report) {
	fmt.Println("--- Relatório Final (Clean Architecture) ---")
	fmt.Printf("Tempo total gasto: %v\n", report.TotalTime)
	fmt.Printf("Quantidade total de requests: %d\n", report.TotalRequests)
	fmt.Printf("Requests com status 200 (OK): %d\n", report.SuccessfulRequests)
	fmt.Println("Distribuição de status HTTP:")
	for code, count := range report.StatusDistribution {
		fmt.Printf("  - Status %d: %d\n", code, count)
	}
	fmt.Println("------------------------------------------")
}
