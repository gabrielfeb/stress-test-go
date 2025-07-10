package usecase

import (
	"stress-test/internal/entity"
	"sync"
	"time"
)

// Define a interface (porta) para fazer uma requisição.
type Requester interface {
	MakeRequest(url string) entity.RequestResult
}

// Define a interface (porta) para apresentar o relatório final.
type Presenter interface {
	Present(report entity.Report)
}

// Encapsula a lógica de negócio para executar o teste de carga.
type RunLoadTestUseCase struct {
	requester Requester
	presenter Presenter
}

// Construtor que injeta as dependências.
func NewRunLoadTestUseCase(r Requester, p Presenter) *RunLoadTestUseCase {
	return &RunLoadTestUseCase{
		requester: r,
		presenter: p,
	}
}

// Orquestra a execução do teste de carga.
func (uc *RunLoadTestUseCase) Execute(url string, totalRequests, concurrency int) {
	startTime := time.Now()
	jobs := make(chan string, totalRequests)
	results := make(chan entity.RequestResult, totalRequests)
	var wg sync.WaitGroup

	for w := 0; w < concurrency; w++ {
		wg.Add(1)
		go uc.worker(&wg, jobs, results)
	}

	for r := 0; r < totalRequests; r++ {
		jobs <- url
	}
	close(jobs)

	wg.Wait()
	close(results)

	report := entity.Report{
		StatusDistribution: make(map[int]int),
	}
	for res := range results {
		report.TotalRequests++
		if res.StatusCode == 200 {
			report.SuccessfulRequests++
		}
		if res.Error != nil {
			continue
		}
		report.StatusDistribution[res.StatusCode]++
	}
	report.TotalTime = time.Since(startTime)

	uc.presenter.Present(report)
}

// Executa em cada goroutine.
func (uc *RunLoadTestUseCase) worker(wg *sync.WaitGroup, jobs <-chan string, results chan<- entity.RequestResult) {
	defer wg.Done()
	for url := range jobs {
		result := uc.requester.MakeRequest(url)
		results <- result
	}
}
