//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Projects struct {
	ID        int32 `sql:"primary_key"`
	Name      string
	OwnerID   int32
	IsPublic  *bool
	CreatedAt *time.Time
	LastEdit  *time.Time
}
