package interfaces

import (
	"time"

	"github.com/google/uuid"
)

type BaseInterface interface {
	GetID() uuid.UUID
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(id uuid.UUID)
	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
}
