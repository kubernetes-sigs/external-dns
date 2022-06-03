package cloudflare

import "fmt"

// AccountNotFoundError indicates an account was not found
type AccountNotFoundError struct {
	ID string
}

func (e *AccountNotFoundError) Error() string {
	return fmt.Sprintf("Account not found with ID [%s]", e.ID)
}

// ZoneNotFoundError indicates an zone was not found
type ZoneNotFoundError struct {
	ID string
}

func (e *ZoneNotFoundError) Error() string {
	return fmt.Sprintf("Zone not found with ID [%s]", e.ID)
}
