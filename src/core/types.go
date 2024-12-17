package core

import "fmt"

type BaseResponseData struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

func (r BaseResponseData) String() string {
	return fmt.Sprintf("Date: %s, Value: %s", r.Date, r.Value)
}
