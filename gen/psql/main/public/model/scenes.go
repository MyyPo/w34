//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Scenes struct {
	ID         int32 `sql:"primary_key"`
	IngameID   int32
	LocationID int32
	Options    string
}
