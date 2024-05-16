// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin/admin.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on AddTile with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddTile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTile with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AddTileMultiError, or nil if none found.
func (m *AddTile) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddTileValidationError{
					field:  "Tile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddTileValidationError{
					field:  "Tile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddTileValidationError{
				field:  "Tile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddTileMultiError(errors)
	}

	return nil
}

// AddTileMultiError is an error wrapping multiple validation errors returned
// by AddTile.ValidateAll() if the designated constraints aren't met.
type AddTileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTileMultiError) AllErrors() []error { return m }

// AddTileValidationError is the validation error returned by AddTile.Validate
// if the designated constraints aren't met.
type AddTileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTileValidationError) ErrorName() string { return "AddTileValidationError" }

// Error satisfies the builtin error interface
func (e AddTileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTileValidationError{}

// Validate checks the field values on RemoveTile with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RemoveTile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveTile with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RemoveTileMultiError, or
// nil if none found.
func (m *RemoveTile) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveTile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCoords()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveTileValidationError{
					field:  "Coords",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveTileValidationError{
					field:  "Coords",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCoords()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveTileValidationError{
				field:  "Coords",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RemoveTileMultiError(errors)
	}

	return nil
}

// RemoveTileMultiError is an error wrapping multiple validation errors
// returned by RemoveTile.ValidateAll() if the designated constraints aren't met.
type RemoveTileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveTileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveTileMultiError) AllErrors() []error { return m }

// RemoveTileValidationError is the validation error returned by
// RemoveTile.Validate if the designated constraints aren't met.
type RemoveTileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveTileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveTileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveTileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveTileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveTileValidationError) ErrorName() string { return "RemoveTileValidationError" }

// Error satisfies the builtin error interface
func (e RemoveTileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveTile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveTileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveTileValidationError{}

// Validate checks the field values on AdminRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AdminRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdminRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AdminRequestMultiError, or
// nil if none found.
func (m *AdminRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AdminRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	switch v := m.Request.(type) {
	case *AdminRequest_AddTile:
		if v == nil {
			err := AdminRequestValidationError{
				field:  "Request",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetAddTile()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AdminRequestValidationError{
						field:  "AddTile",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdminRequestValidationError{
						field:  "AddTile",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAddTile()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdminRequestValidationError{
					field:  "AddTile",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AdminRequest_RemoveTile:
		if v == nil {
			err := AdminRequestValidationError{
				field:  "Request",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetRemoveTile()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AdminRequestValidationError{
						field:  "RemoveTile",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdminRequestValidationError{
						field:  "RemoveTile",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRemoveTile()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdminRequestValidationError{
					field:  "RemoveTile",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return AdminRequestMultiError(errors)
	}

	return nil
}

// AdminRequestMultiError is an error wrapping multiple validation errors
// returned by AdminRequest.ValidateAll() if the designated constraints aren't met.
type AdminRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdminRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdminRequestMultiError) AllErrors() []error { return m }

// AdminRequestValidationError is the validation error returned by
// AdminRequest.Validate if the designated constraints aren't met.
type AdminRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminRequestValidationError) ErrorName() string { return "AdminRequestValidationError" }

// Error satisfies the builtin error interface
func (e AdminRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminRequestValidationError{}

// Validate checks the field values on AdminResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AdminResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdminResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AdminResponseMultiError, or
// nil if none found.
func (m *AdminResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AdminResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ServerId

	if all {
		switch v := interface{}(m.GetLog()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AdminResponseValidationError{
					field:  "Log",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AdminResponseValidationError{
					field:  "Log",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLog()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdminResponseValidationError{
				field:  "Log",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Result

	if len(errors) > 0 {
		return AdminResponseMultiError(errors)
	}

	return nil
}

// AdminResponseMultiError is an error wrapping multiple validation errors
// returned by AdminResponse.ValidateAll() if the designated constraints
// aren't met.
type AdminResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdminResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdminResponseMultiError) AllErrors() []error { return m }

// AdminResponseValidationError is the validation error returned by
// AdminResponse.Validate if the designated constraints aren't met.
type AdminResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminResponseValidationError) ErrorName() string { return "AdminResponseValidationError" }

// Error satisfies the builtin error interface
func (e AdminResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminResponseValidationError{}
