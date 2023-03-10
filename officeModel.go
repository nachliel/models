package models

import (
	"gorm.io/gorm"
)

// Office is the structure for database v1.1
type Office struct {
	gorm.Model
	Fid          int
	GlobalID     string
	AreaName     string
	Link         string
	Phone        string
	Fax          string
	Address      string
	Mail         string
	Name         string
	StatusActive bool
}

// OfficeResponse is the response after you have SpacialReferences
type OfficeResponse struct {
	Features []Feature    `json:"features"`
	Error    *OfficeError `json:"error"`
}

type Feature struct {
	Attributes Attribute `json:"attributes"`
}

type Attribute struct {
	Address  string `json:"Adress"`
	Fax      string `json:"Fax"`
	Link     string `json:"Link"`
	Mail     string `json:"Mail"`
	Palkid   string `json:"Palkid"`
	Tel      string `json:"Tel"`
	EzorName string `json:"ezor_name"`
	FID      int    `json:"FID"`
	GlobalID string `json:"GlobalID"`
}

type OfficeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SpactialResponse is the first inquiry response for address coordiantes
type SpactialResponse struct {
	Candidates []candidate `json:"candidates"`
}

type candidate struct {
	Address  string    `json:"address"`
	Location Locations `json:"location"`
	Score    float64   `json:"score"`
	Extent   Extents   `json:"extent"`
}

type Locations struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Extents struct {
	Xmin float64 `json:"xmin"`
	Ymin float64 `json:"ymin"`
	Xmax float64 `json:"xmax"`
	Ymax float64 `json:"ymax"`
}

// GushResponse is the struct type of Gush inquiry.
type GushResponse struct {
	Features []FeatureGush `json:"features"`
	Error    *OfficeError  `json:"error"`
}

type FeatureGush struct {
	Attributes AttributeGush `json:"attributes"`
	Geometry   Geometry      `json:"geometry"`
}

type AttributeGush struct {
	FID           int    `json:"FID"`
	Gush          int    `json:"GUSH_NUM"`
	LocalityIndex int    `json:"LOCALITY_I"`
	LocalityName  string `json:"LOCALITY_N"`
}

type Geometry struct {
	Rings [][][]float64 `json:"rings"`
}
