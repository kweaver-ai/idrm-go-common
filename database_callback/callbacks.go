package database_callback

import "gorm.io/gorm"

type EntityChangeCallback Register
type DataLineageCallback Register

func NewEmptyEntityChangeCallback() EntityChangeCallback {
	return NewDefaultRegister()
}

func NewEmptyDataLineageCallback() DataLineageCallback {
	return NewDefaultRegister()
}

type Callbacks struct {
	db           *gorm.DB
	entityChange EntityChangeCallback
	dataLineage  DataLineageCallback
}

func NewCallbacks(db *gorm.DB, e EntityChangeCallback, d DataLineageCallback) *Callbacks {
	return &Callbacks{
		db:           db,
		entityChange: e,
		dataLineage:  d,
	}
}

func (c *Callbacks) Registers() {
	if !c.dataLineage.IsEmpty() {
		c.dataLineage.RegisterCallback(c.db)
	}
	if !c.entityChange.IsEmpty() {
		c.entityChange.RegisterCallback(c.db)
	}
}
