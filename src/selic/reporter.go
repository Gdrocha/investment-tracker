package selic

import (
	"encoding/json"
	"fmt"
	"investment-tracker/src/core/interfaces"
	"investment-tracker/src/core/registry"
	"os"
	"time"
)

type SelicReporter struct{}

func (f *SelicReporter) Report() {
	var selicFetcher SelicFetcher

	cachePath := "src/selic/.cache.json"
	lastRate, err := ReadLastSelicRate(cachePath)

	if err != nil {
		fmt.Println("Error fetching last Selic rate:", err)
		return
	}

	selicRates, err := selicFetcher.Fetch(lastRate.Date, time.Now().Format("02/01/2006"))

	if err != nil {
		fmt.Println("Error fetching Selic rate:", err)
		return
	}

	if len(selicRates) > 0 {
		newRate := selicRates[len(selicRates)-1]
		err := UpdateLastSelicRate(cachePath, lastRate, newRate)

		if err != nil {
			fmt.Println("Error updating the last Selic rate:", err)
			return
		}
	}
}

func ReadLastSelicRate(filePath string) (SelicResponseData, error) {
	var lastSelicRate SelicResponseData

	fileInfo, err := os.Stat(filePath)

	if os.IsNotExist(err) || fileInfo.Size() == 0 {
		return SelicResponseData{Date: time.Now().Format("02/01/2006")}, nil
	}

	data, err := os.ReadFile(filePath)

	if err != nil {
		return lastSelicRate, err
	}

	err = json.Unmarshal(data, &lastSelicRate)
	return lastSelicRate, err
}

func UpdateLastSelicRate(filePath string, lastRate SelicResponseData, newRate SelicResponseData) error {
	fmt.Printf("Updating Selic rate! Old: %s%%, New: %s%%\n", lastRate.Value, newRate.Value)

	data, err := json.Marshal(newRate)

	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	return err
}

func init() {
	registry.GetRegistry().Register((*interfaces.Reporter)(nil), &SelicReporter{})
}
