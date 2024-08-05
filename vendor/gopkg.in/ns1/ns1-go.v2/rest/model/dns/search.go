package dns

type SearchResult struct {
	Next         string    `json:"next"`
	Limit        int       `json:"limit"`
	TotalResults int       `json:"total_results"`
	Results      []*Record `json:"results"`
}
