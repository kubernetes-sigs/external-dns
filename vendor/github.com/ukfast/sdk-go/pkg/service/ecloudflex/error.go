package ecloudflex

import "fmt"

// ProjectNotFoundError indicates a project was not found
type ProjectNotFoundError struct {
	ID int
}

func (e *ProjectNotFoundError) Error() string {
	return fmt.Sprintf("Project not found with ID [%d]", e.ID)
}
