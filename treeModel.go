package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Tree struct {
	gorm.Model
	UserID    uuid.UUID
	RequestID int
	Name      string
	Size      int
	Count     int
	CutTree   bool
	Reason    string
	Images    string
	Premit    bool
}

type TreeSummery struct {
	ID        uint
	RequestID int
	Name      string
	Size      int
	Count     int
	CutTree   bool
	Reason    string
	Images    string
}
