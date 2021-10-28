package loadbalancer

import "fmt"

// TargetNotFoundError indicates a target was not found
type TargetNotFoundError struct {
	ID int
}

func (e *TargetNotFoundError) Error() string {
	return fmt.Sprintf("Target not found with ID [%d]", e.ID)
}

// ClusterNotFoundError indicates a cluster was not found
type ClusterNotFoundError struct {
	ID int
}

func (e *ClusterNotFoundError) Error() string {
	return fmt.Sprintf("Cluster not found with ID [%d]", e.ID)
}

// TargetGroupNotFoundError indicates a target group was not found
type TargetGroupNotFoundError struct {
	ID int
}

func (e *TargetGroupNotFoundError) Error() string {
	return fmt.Sprintf("Target group not found with ID [%d]", e.ID)
}

// VIPNotFoundError indicates a VIP was not found
type VIPNotFoundError struct {
	ID int
}

func (e *VIPNotFoundError) Error() string {
	return fmt.Sprintf("VIP not found with ID [%d]", e.ID)
}

// ListenerNotFoundError indicates a listener was not found
type ListenerNotFoundError struct {
	ID int
}

func (e *ListenerNotFoundError) Error() string {
	return fmt.Sprintf("Listener not found with ID [%d]", e.ID)
}

// AccessIPNotFoundError indicates an access IP was not found
type AccessIPNotFoundError struct {
	ID int
}

func (e *AccessIPNotFoundError) Error() string {
	return fmt.Sprintf("Access IP not found with ID [%d]", e.ID)
}

// BindNotFoundError indicates a bind was not found
type BindNotFoundError struct {
	ID int
}

func (e *BindNotFoundError) Error() string {
	return fmt.Sprintf("Bind not found with ID [%d]", e.ID)
}

// CertificateNotFoundError indicates a certificate was not found
type CertificateNotFoundError struct {
	ID int
}

func (e *CertificateNotFoundError) Error() string {
	return fmt.Sprintf("Certificate not found with ID [%d]", e.ID)
}

// ACLNotFoundError indicates a certificate was not found
type ACLNotFoundError struct {
	ID int
}

func (e *ACLNotFoundError) Error() string {
	return fmt.Sprintf("ACL not found with ID [%d]", e.ID)
}
