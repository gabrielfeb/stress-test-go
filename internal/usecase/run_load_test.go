package usecase

import (
	"errors"
	"stress-test/internal/entity"
	"testing"
)

// Simula erros ou sucessos nos testes
type MockRequester struct {
	ShouldError bool
}

// Simula o comportamento do Requester real
func (m *MockRequester) MakeRequest(url string) entity.RequestResult {
	if m.ShouldError {
		return entity.RequestResult{
			Error: errors.New("simulated connection error"),
		}
	}
	return entity.RequestResult{
		StatusCode: 200,
	}
}

// Captura o resultado do teste
type MockPresenter struct {
	FinalReport entity.Report
}

// Captura o relatório em vez de imprimi-lo
func (m *MockPresenter) Present(report entity.Report) {
	m.FinalReport = report
}

// Teste para o cenário de SUCESSO.
func TestRunLoadTestUseCase_Execute_Success(t *testing.T) {
	mockRequester := &MockRequester{ShouldError: false}
	mockPresenter := &MockPresenter{}
	usecase := NewRunLoadTestUseCase(mockRequester, mockPresenter)
	totalRequests := 20

	usecase.Execute("http://fake-url.com", totalRequests, 5)

	if mockPresenter.FinalReport.TotalRequests != totalRequests {
		t.Errorf("Esperado %d requests no total, mas obtido %d", totalRequests, mockPresenter.FinalReport.TotalRequests)
	}
	if mockPresenter.FinalReport.SuccessfulRequests != totalRequests {
		t.Errorf("Esperado %d requests com sucesso, mas obtido %d", totalRequests, mockPresenter.FinalReport.SuccessfulRequests)
	}
}

// Teste para o cenário de ERRO.
func TestRunLoadTestUseCase_Execute_WithError(t *testing.T) {
	mockRequester := &MockRequester{ShouldError: true}
	mockPresenter := &MockPresenter{}
	usecase := NewRunLoadTestUseCase(mockRequester, mockPresenter)
	totalRequests := 10

	usecase.Execute("http://fake-url.com", totalRequests, 2)

	if mockPresenter.FinalReport.TotalRequests != totalRequests {
		t.Errorf("Esperado %d requests no total, mas obtido %d", totalRequests, mockPresenter.FinalReport.TotalRequests)
	}
	if mockPresenter.FinalReport.SuccessfulRequests != 0 {
		t.Errorf("Esperado 0 requests com sucesso, mas obtido %d", mockPresenter.FinalReport.SuccessfulRequests)
	}
}
