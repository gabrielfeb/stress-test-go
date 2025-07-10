package entity

import "time"

// Representa o resultado de uma única requisição.
type RequestResult struct {
	StatusCode int
	Error      error
}

// Armazena os resultados agregados do teste de carga.
type Report struct {
	TotalTime          time.Duration
	TotalRequests      int
	SuccessfulRequests int
	StatusDistribution map[int]int
}
