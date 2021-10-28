package safedns

import (
	"fmt"
)

// ZoneNotFoundError indicates a zone was not found within SafeDNS
type ZoneNotFoundError struct {
	ZoneName string
}

// ZoneRecordNotFoundError indicates a record was not found within SafeDNS
type ZoneRecordNotFoundError struct {
	ZoneName string
	RecordID int
}

// ZoneNoteNotFoundError indicates a zone note was not found within SafeDNS
type ZoneNoteNotFoundError struct {
	ZoneName string
	NoteID   int
}

// TemplateNotFoundError indicates a template was not found within SafeDNS
type TemplateNotFoundError struct {
	TemplateID int
}

// TemplateRecordNotFoundError indicates a record was not found within SafeDNS
type TemplateRecordNotFoundError struct {
	TemplateID int
	RecordID   int
}

func (e *ZoneNotFoundError) Error() string {
	return fmt.Sprintf("zone not found with name [%s]", e.ZoneName)
}

func (e *ZoneRecordNotFoundError) Error() string {
	return fmt.Sprintf("record not found with ID [%d] in zone [%s]", e.RecordID, e.ZoneName)
}

func (e *ZoneNoteNotFoundError) Error() string {
	return fmt.Sprintf("zone note not found with ID [%d] in zone [%s]", e.NoteID, e.ZoneName)
}

func (e *TemplateNotFoundError) Error() string {
	return fmt.Sprintf("template not found with ID [%d]", e.TemplateID)
}

func (e *TemplateRecordNotFoundError) Error() string {
	return fmt.Sprintf("record not found with ID [%d] in template [%d]", e.RecordID, e.TemplateID)
}
