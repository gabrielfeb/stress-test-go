package requester

import (
	"net/http"
	"stress-test/internal/entity"
)

// Implementação concreta para fazer requisições HTTP.
type HttpRequester struct{}

func NewHttpRequester() *HttpRequester {
	return &HttpRequester{}
}

// Implementa a interface Requester.
func (r *HttpRequester) MakeRequest(url string) entity.RequestResult {
	resp, err := http.Get(url)
	if err != nil {
		return entity.RequestResult{Error: err}
	}
	defer resp.Body.Close()

	return entity.RequestResult{
		StatusCode: resp.StatusCode,
	}
}
