package sharedexchange

import "fmt"

// DomainNotFoundError indicates a domain was not found
type DomainNotFoundError struct {
	ID int
}

func (e *DomainNotFoundError) Error() string {
	return fmt.Sprintf("Domain not found with ID [%d]", e.ID)
}
