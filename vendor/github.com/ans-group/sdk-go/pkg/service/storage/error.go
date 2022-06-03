package storage

import "fmt"

// SolutionNotFoundError indicates a virtual machine was not found
type SolutionNotFoundError struct {
	ID int
}

func (e *SolutionNotFoundError) Error() string {
	return fmt.Sprintf("solution not found with ID [%d]", e.ID)
}

// VolumeNotFoundError indicates a virtual machine was not found
type VolumeNotFoundError struct {
	ID int
}

func (e *VolumeNotFoundError) Error() string {
	return fmt.Sprintf("volume not found with ID [%d]", e.ID)
}

// HostNotFoundError indicates a virtual machine was not found
type HostNotFoundError struct {
	ID int
}

func (e *HostNotFoundError) Error() string {
	return fmt.Sprintf("host not found with ID [%d]", e.ID)
}
