// Code generated by sysl DO NOT EDIT.
package pokeapi

import (
	"time"

	"github.com/anz-bank/sysl-go/validator"

	"github.com/rickb777/date"
)

// Reference imports to suppress unused errors
var _ = time.Parse

// Reference imports to suppress unused errors
var _ = date.Parse

// Error_ ...
type Error_ struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Pokemon ...
type Pokemon struct {
	Height *int64  `json:"height,omitempty"`
	ID     *int64  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// GetPokemonRequest ...
type GetPokemonRequest struct {
	ID int64
}

// *Error_ validator
func (s *Error_) Validate() error {
	return validator.Validate(s)
}

// *Pokemon validator
func (s *Pokemon) Validate() error {
	return validator.Validate(s)
}
