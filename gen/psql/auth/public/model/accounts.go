//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Accounts struct {
	UserID        uuid.UUID
	Username      string
	Password      string
	Email         string
	AccountStatus Status
	Online        bool
	CreatedAt     time.Time
	LastLogin     time.Time
}
