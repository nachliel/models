package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// consts
const (
	ClassTypeOwner            = "owner"
	ClassTypeTenant           = "tenant"
	ClassTypeSharedHouseOwner = "apartment"
	ClassTypeRenter           = "renter"

	RequestStatusOpen       = "open"
	RequestStatusSubmit     = "submit"
	RequestStatusPayed      = "payed"
	RequestStatusProccessed = "proccess"
	RequestStatusComplete   = "complete"
	RequestStatusError      = "error"

	DocumentTypeAuth       = "auth"
	DocumentTypeSigning    = "signing"
	DocumentTypePlanA      = "plan-a"
	DocumentTypePlanB      = "plan-b"
	DocumentTypePlanC      = "plan-c"
	DocumentTypePlanD      = "plan-d"
	DocumentTypeMedical    = "medical"
	DocumentTypeAgronomist = "agronom"
	DocumentTypeTreeSurvey = "survey"
	DocumentTypeOther      = "other"
	DocumentTypeFinal      = "final"
)

type Request struct {
	gorm.Model
	UserID                 uuid.UUID
	City                   string
	Address                string
	Gush                   int
	Helka                  int
	ClassType              string
	ReasonSafetyIssue      bool
	ReasonSickTreeIssue    bool
	ReasonHealthIssue      bool
	ReasonDevelopmentIssue bool
	ReasonAgriIssue        bool
	ReasonOtherIssue       bool
	TollPayment            string
	AttachAuthFile         string
	AttachSigningFile      string
	AttachPlanA            string
	AttachPlanB            string
	AttachPlanC            string
	AttachPlanD            string
	AttachMedical          string
	AttachAgronomist       string
	AttachTreeSurvey       string
	AttachOtherDoc         string
	DocumentOutPut         string
	SubmitDate             time.Time
	PaymentDate            time.Time
	PaymentToken           string
	SentTime               time.Time
	AuthorityID            int
	Status                 string
	Error                  string
}

type RequestSummery struct {
	gorm.Model
	City           string
	Address        string
	Gush           int
	Helka          int
	ClassType      string
	DocumentOutPut string
	SubmitDate     time.Time
	PaymentDate    time.Time
	PaymentToken   string
	SentTime       time.Time
	AuthorityID    int
	Status         string
	Error          string
}
