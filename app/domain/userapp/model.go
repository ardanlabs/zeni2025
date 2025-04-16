package userapp

import (
	"encoding/json"
	"fmt"

	"github.com/ardanlabs/service/app/sdk/errs"
)

type User struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Roles        []string `json:"roles"`
	PasswordHash []byte   `json:"-"`
	Department   string   `json:"department"`
	Enabled      bool     `json:"enabled"`
	DateCreated  string   `json:"dateCreated"`
	DateUpdated  string   `json:"dateUpdated"`
}

// =============================================================================

type NewUser struct {
	Name            string   `json:"name" validate:"required"`
	Email           string   `json:"email" validate:"required,email"`
	Roles           []string `json:"roles" validate:"required"`
	Department      string   `json:"department"`
	Password        string   `json:"password" validate:"required"`
	PasswordConfirm string   `json:"passwordConfirm" validate:"eqfield=Password"`
}

func (app *NewUser) Decode(data []byte) error {
	return json.Unmarshal(data, app)
}

func (app NewUser) Validate() error {
	if err := errs.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}

// =============================================================================

type UpdateUserRole struct {
	Roles []string `json:"roles" validate:"required"`
}

func (app *UpdateUserRole) Decode(data []byte) error {
	return json.Unmarshal(data, app)
}

func (app UpdateUserRole) Validate() error {
	if err := errs.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}

// =============================================================================

type UpdateUser struct {
	Name            *string `json:"name"`
	Email           *string `json:"email" validate:"omitempty,email"`
	Department      *string `json:"department"`
	Password        *string `json:"password"`
	PasswordConfirm *string `json:"passwordConfirm" validate:"omitempty,eqfield=Password"`
	Enabled         *bool   `json:"enabled"`
}

func (app *UpdateUser) Decode(data []byte) error {
	return json.Unmarshal(data, app)
}

func (app UpdateUser) Validate() error {
	if err := errs.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}
