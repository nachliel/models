package models

import (
	//"github.com/google/uuid"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID uuid.UUID `gorm:"primary_key;not null;default:gen_random_uuid()"`
}

type Client struct {
	Model
	Phone        string `gorm:"not null"`
	FullName     string
	GovId        string
	Email        string
	City         string
	Address      string
	Salt         string `gorm:"not null"`
	Password     string `gorm:"not null"`
	TempPassword string
	TempPassTime time.Time
	LastLogin    time.Time
	FailedLogin  int8
	Status       string
	SessionID    string
	SessionStart time.Time
	SessionEnd   time.Time
	Role         string
}

type SubUser struct {
	// User uuid
	ID           string
	FullName     string
	GovId        string
	Email        string
	City         string
	Address      string
	Phone        string
	LastLogin    time.Time
	Status       string
	SessionID    string
	SessionStart time.Time
	SessionEnd   time.Time
	Role         string
}

type UserSession struct {
	gorm.Model
	UserSessionID string
	UserID        uuid.UUID
	LastAction    string
	Ips           string
	MacAddress    string
	UserAgent     string
	SessionStart  time.Time
	SessionEnd    time.Time
}
