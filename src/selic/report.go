package selic

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type SelicData struct {
	Date  string `json:"data"`
	Value string `json:"valor"`
}

func ReportSelicRate() {
	cachePath := "src/selic/cache/.lastSelicRate.json"
	lastRate, err := ReadLastSelicRate(cachePath)

	if err != nil {
		fmt.Println("Error fetching last Selic rate:", err)
		return
	}

	selicRates, err := FetchSelicRate(lastRate.Date, time.Now().Format("02/01/2006"))

	if err != nil {
		fmt.Println("Error fetching Selic rate:", err)
		return
	}

	if len(selicRates) > 0 {
		newRate := selicRates[len(selicRates)-1]

		if newRate.Value == lastRate.Value {
			fmt.Printf("Selic rate unchanged. Storing current value into cache. Rate: %s%%\n", newRate.Value)
			err := UpdateLastSelicRate("src/selic/cache/.lastSelicRate.json", newRate)

			if err != nil {
				fmt.Println("Error updating the last Selic rate:", err)
				return
			}
		} else {
			fmt.Printf("Selic rate changed! Old: %s%%, New: %s%%\n", lastRate.Value, newRate.Value)

			err := UpdateLastSelicRate("src/selic/cache/.lastSelicRate.json", newRate)

			if err != nil {
				fmt.Println("Error updating the last Selic rate:", err)
				return
			}
		}
	}
}

func ReadLastSelicRate(filePath string) (SelicData, error) {
	var lastSelicRate SelicData

	fileInfo, err := os.Stat(filePath)

	if os.IsNotExist(err) || fileInfo.Size() == 0 {
		return SelicData{Date: time.Now().Format("02/01/2006")}, nil
	}

	data, err := os.ReadFile(filePath)

	if err != nil {
		return lastSelicRate, err
	}

	err = json.Unmarshal(data, &lastSelicRate)
	return lastSelicRate, err
}

func UpdateLastSelicRate(filePath string, rate SelicData) error {
	data, err := json.Marshal(rate)

	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	return err
}
