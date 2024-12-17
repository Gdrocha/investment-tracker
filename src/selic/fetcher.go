package selic

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SelicFetcher struct{}

func (f *SelicFetcher) Fetch(startDate, endDate string) ([]SelicResponseData, error) {
	println("Fetching selic rate")
	url := fmt.Sprintf("https://api.bcb.gov.br/dados/serie/bcdata.sgs.432/dados?formato=json&dataInicial=%s&dataFinal=%s", startDate, endDate)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var selicRates []SelicResponseData

	if err := json.NewDecoder(resp.Body).Decode(&selicRates); err != nil {
		return nil, err
	}

	return selicRates, nil
}
