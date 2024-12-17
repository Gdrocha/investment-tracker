package selic

import "fmt"

type SelicResponseData struct {
	Date  string `json:"data"`
	Value string `json:"valor"`
}

func (r SelicResponseData) String() string {
	return fmt.Sprintf("Date: %s, Value: %s", r.Date, r.Value)
}
