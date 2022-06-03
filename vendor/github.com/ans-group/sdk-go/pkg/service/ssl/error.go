package ssl

import "fmt"

// CertificateNotFoundError indicates a virtual machine was not found
type CertificateNotFoundError struct {
	ID int
}

func (e *CertificateNotFoundError) Error() string {
	return fmt.Sprintf("certificate not found with ID [%d]", e.ID)
}
