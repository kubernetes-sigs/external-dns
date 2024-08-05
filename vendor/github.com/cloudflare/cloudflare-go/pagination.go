package cloudflare

<<<<<<< HEAD
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
import (
	"math"
)

// Look first for total_pages, but if total_count and per_page are set then use that to get page count.
func (p ResultInfo) getTotalPages() int {
	totalPages := p.TotalPages
	if totalPages == 0 && p.Total > 0 && p.PerPage > 0 {
		totalPages = int(math.Ceil(float64(p.Total) / float64(p.PerPage)))
	}
	return totalPages
}

// Done returns true for the last page and false otherwise.
func (p ResultInfo) Done() bool {
	// A little hacky but if the response body is lacking a defined `ResultInfo`
	// object the page will be 1 however the counts will be empty so if we have
	// that response, we just assume this is the only page.
	totalPages := p.getTotalPages()
	if p.Page == 1 && totalPages == 0 {
		return true
	}

	return p.Page > 1 && p.Page > totalPages
}

// Next advances the page of a paginated API response, but does not fetch the
// next page of results.
func (p ResultInfo) Next() ResultInfo {
	// A little hacky but if the response body is lacking a defined `ResultInfo`
	// object the page will be 1 however the counts will be empty so if we have
	// that response, we just assume this is the only page.
	totalPages := p.getTotalPages()
	if p.Page == 1 && totalPages == 0 {
		return p
	}

	// This shouldn't happen normally however, when it does just return the
	// current page.
	if p.Page > totalPages {
		return p
	}

	p.Page++
	return p
}

// HasMorePages returns whether there is another page of results after the
// current one.
func (p ResultInfo) HasMorePages() bool {
	totalPages := p.getTotalPages()
	if totalPages == 0 {
		return false
	}

	return p.Page >= 1 && p.Page < totalPages
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
