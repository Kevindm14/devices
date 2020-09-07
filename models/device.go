package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/pop/v5/slices"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Device is used by pop to map your devices database table to your go code.
type Device struct {
	ID              uuid.UUID     `json:"id" db:"id"`
	Manufacture     string        `json:"manufacture" db:"manufacture"`
	Make            string        `json:"make" db:"make"`
	Model           string        `json:"model" db:"model"`
	Storage         string        `json:"storage" db:"storage"`
	Cost            float64       `json:"cost" db:"cost"`
	OperatingSystem slices.String `json:"operating_system" db:"operating_system"`
	Image           string        `json:"image" db:"image"`
	IsNew           bool          `json:"is_new" db:"is_new"`
	CreatedAt       time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (d Device) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Devices is not required by pop and may be deleted
type Devices []Device

// String is not required by pop and may be deleted
func (d Devices) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Device) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.Manufacture, Name: "Manufacture"},
		&validators.StringIsPresent{Field: d.Make, Name: "Make"},
		&validators.StringIsPresent{Field: d.Model, Name: "Model"},
		&validators.StringIsPresent{Field: d.Storage, Name: "Storage"},
		&validators.StringIsPresent{Field: d.Image, Name: "Image"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Device) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Device) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
