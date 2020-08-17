package main

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Relay is an anonymous email forwarder
type Relay struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Destination string
}

// BeforeCreate sets the relay's ID to a new UUID
func (r *Relay) BeforeCreate(s *gorm.Scope) error {
	uuid, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	return s.SetColumn("ID", uuid)
}
