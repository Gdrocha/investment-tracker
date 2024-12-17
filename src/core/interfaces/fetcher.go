package interfaces

type Fetcher interface {
	Fetch(startDate, endDate string) (string, error)
}
