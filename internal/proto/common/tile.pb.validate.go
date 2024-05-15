// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/tile.proto

package common

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

// Validate checks the field values on Coordinate with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Coordinate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Coordinate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CoordinateMultiError, or
// nil if none found.
func (m *Coordinate) ValidateAll() error {
	return m.validate(true)
}

func (m *Coordinate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for X

	// no validation rules for Y

	if len(errors) > 0 {
		return CoordinateMultiError(errors)
	}

	return nil
}

// CoordinateMultiError is an error wrapping multiple validation errors
// returned by Coordinate.ValidateAll() if the designated constraints aren't met.
type CoordinateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CoordinateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CoordinateMultiError) AllErrors() []error { return m }

// CoordinateValidationError is the validation error returned by
// Coordinate.Validate if the designated constraints aren't met.
type CoordinateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CoordinateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CoordinateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CoordinateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CoordinateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CoordinateValidationError) ErrorName() string { return "CoordinateValidationError" }

// Error satisfies the builtin error interface
func (e CoordinateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCoordinate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CoordinateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CoordinateValidationError{}

// Validate checks the field values on Tile with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Tile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tile with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TileMultiError, or nil if none found.
func (m *Tile) ValidateAll() error {
	return m.validate(true)
}

func (m *Tile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if all {
		switch v := interface{}(m.GetCoords()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TileValidationError{
					field:  "Coords",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TileValidationError{
					field:  "Coords",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCoords()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TileValidationError{
				field:  "Coords",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TileMultiError(errors)
	}

	return nil
}

// TileMultiError is an error wrapping multiple validation errors returned by
// Tile.ValidateAll() if the designated constraints aren't met.
type TileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TileMultiError) AllErrors() []error { return m }

// TileValidationError is the validation error returned by Tile.Validate if the
// designated constraints aren't met.
type TileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TileValidationError) ErrorName() string { return "TileValidationError" }

// Error satisfies the builtin error interface
func (e TileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TileValidationError{}
