package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Sign struct {
	gorm.Model
	UserID     uuid.UUID
	RequestID  int
	Apartments int
	Signatures SignaturesArray `gorm:"type:jsonb"`
	Status     bool
	Note       string
}

type SignSummery struct {
	ID         uint
	Apartments int
	Signatures SignaturesArray `gorm:"type:jsonb"`
	Status     bool
	Note       string
}

type Signatures struct {
	FullName  string    `json:"fullname"`
	ID        string    `json:"id"`
	Phone     string    `json:"phone"`
	Apartment int       `json:"apartment"`
	Allowance bool      `json:"allowance"`
	Note      string    `json:"note"`
	SignTime  time.Time `json:"signtime"`
}

type SignaturesArray []Signatures

func (a SignaturesArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *SignaturesArray) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), a)
}
