package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Allocator An allocator object represent an account on the blockchain.
//
// swagger:model Allocator
type Allocator struct {

	// The initial balance of the account, unit is in Wei.
	// Units * 10^0  Wei
	//       * 10^12 Szabo
	//       * 10^15 Finney
	//       * 10^18 Ether
	//
	Balance string `json:"balance,omitempty"`

	// The code for the account.
	Code string `json:"code,omitempty"`

	// The private key generate along with the use of new_allocator,
	// DO NOT SUBMIT YOUR PRIVATE KEY if you fills allocators.
	//
	PrivateKey string `json:"private_key,omitempty"`

	// The initial storage.
	Storage string `json:"storage,omitempty"`
}

// Validate validates this allocator
func (m *Allocator) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
