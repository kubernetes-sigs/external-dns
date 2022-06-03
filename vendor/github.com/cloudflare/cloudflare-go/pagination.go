package cloudflare

// Done returns true for the last page and false otherwise.
func (p ResultInfo) Done() bool {
	return p.Page > 1 && p.Page > p.TotalPages
}

// Next advances the page of a paginated API response, but does not fetch the
// next page of results.
func (p ResultInfo) Next() ResultInfo {
	p.Page++
	return p
}

// HasMorePages returns whether there is another page of results after the
// current one.
func (p ResultInfo) HasMorePages() bool {
	return p.Page > 1 && p.Page < p.TotalPages
}
