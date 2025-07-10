package main

import (
	"flag"
	"stress-test/internal/infra/presenter"
	"stress-test/internal/infra/requester"
	"stress-test/internal/usecase"
)

func main() {
	// Parse de flags
	url := flag.String("url", "", "URL do serviço a ser testado.")
	requests := flag.Int("requests", 0, "Número total de requests.")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas.")
	flag.Parse()

	if *url == "" || *requests == 0 {
		flag.PrintDefaults()
		return
	}

	httpRequester := requester.NewHttpRequester()
	cliPresenter := presenter.NewCliPresenter()
	runLoadTest := usecase.NewRunLoadTestUseCase(httpRequester, cliPresenter)

	runLoadTest.Execute(*url, *requests, *concurrency)
}
