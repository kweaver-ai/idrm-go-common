package database_callback

import "gorm.io/gorm"

type Register interface {
	RegisterCallback(db *gorm.DB)
	IsEmpty() bool
}

type EmptyRegister struct{}

func (r EmptyRegister) IsEmpty() bool {
	return true
}

func (r EmptyRegister) RegisterCallback(db *gorm.DB) {
	return
}

func NewDefaultRegister() Register {
	return &EmptyRegister{}
}
