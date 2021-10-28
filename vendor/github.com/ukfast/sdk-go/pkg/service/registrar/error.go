package registrar

import "fmt"

// DomainNotFoundError indicates a domain was not found
type DomainNotFoundError struct {
	Name string
}

func (e *DomainNotFoundError) Error() string {
	return fmt.Sprintf("Domain not found with name [%s]", e.Name)
}
