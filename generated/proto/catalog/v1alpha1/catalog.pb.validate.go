// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/catalog/v1alpha1/catalog.proto

package v1alpha1

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

// Validate checks the field values on Catalog with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Catalog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Catalog with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CatalogMultiError, or nil if none found.
func (m *Catalog) ValidateAll() error {
	return m.validate(true)
}

func (m *Catalog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetType()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CatalogValidationError{
					field:  "Type",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CatalogValidationError{
					field:  "Type",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CatalogValidationError{
				field:  "Type",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLanguages()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CatalogValidationError{
					field:  "Languages",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CatalogValidationError{
					field:  "Languages",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLanguages()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CatalogValidationError{
				field:  "Languages",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Owner

	// no validation rules for Version

	// no validation rules for Link

	// no validation rules for Description

	// no validation rules for Repository

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CatalogValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CatalogValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CatalogValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CatalogValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CatalogValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CatalogValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CatalogMultiError(errors)
	}

	return nil
}

// CatalogMultiError is an error wrapping multiple validation errors returned
// by Catalog.ValidateAll() if the designated constraints aren't met.
type CatalogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CatalogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CatalogMultiError) AllErrors() []error { return m }

// CatalogValidationError is the validation error returned by Catalog.Validate
// if the designated constraints aren't met.
type CatalogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CatalogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CatalogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CatalogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CatalogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CatalogValidationError) ErrorName() string { return "CatalogValidationError" }

// Error satisfies the builtin error interface
func (e CatalogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCatalog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CatalogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CatalogValidationError{}

// Validate checks the field values on Type with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Type) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Type with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TypeMultiError, or nil if none found.
func (m *Type) ValidateAll() error {
	return m.validate(true)
}

func (m *Type) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Module

	// no validation rules for Library

	// no validation rules for Workflow

	// no validation rules for Project

	// no validation rules for Chart

	// no validation rules for Package

	// no validation rules for Container

	if len(errors) > 0 {
		return TypeMultiError(errors)
	}

	return nil
}

// TypeMultiError is an error wrapping multiple validation errors returned by
// Type.ValidateAll() if the designated constraints aren't met.
type TypeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TypeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TypeMultiError) AllErrors() []error { return m }

// TypeValidationError is the validation error returned by Type.Validate if the
// designated constraints aren't met.
type TypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TypeValidationError) ErrorName() string { return "TypeValidationError" }

// Error satisfies the builtin error interface
func (e TypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TypeValidationError{}

// Validate checks the field values on Languages with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Languages) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Languages with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LanguagesMultiError, or nil
// if none found.
func (m *Languages) ValidateAll() error {
	return m.validate(true)
}

func (m *Languages) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Golang

	// no validation rules for Kotlin

	// no validation rules for Java

	// no validation rules for Terraform

	// no validation rules for Helm

	// no validation rules for Javascript

	// no validation rules for Yaml

	// no validation rules for Docker

	if len(errors) > 0 {
		return LanguagesMultiError(errors)
	}

	return nil
}

// LanguagesMultiError is an error wrapping multiple validation errors returned
// by Languages.ValidateAll() if the designated constraints aren't met.
type LanguagesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LanguagesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LanguagesMultiError) AllErrors() []error { return m }

// LanguagesValidationError is the validation error returned by
// Languages.Validate if the designated constraints aren't met.
type LanguagesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LanguagesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LanguagesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LanguagesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LanguagesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LanguagesValidationError) ErrorName() string { return "LanguagesValidationError" }

// Error satisfies the builtin error interface
func (e LanguagesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLanguages.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LanguagesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LanguagesValidationError{}

// Validate checks the field values on CreateUpdateCatalogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUpdateCatalogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUpdateCatalogRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUpdateCatalogRequestMultiError, or nil if none found.
func (m *CreateUpdateCatalogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUpdateCatalogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Service

	// no validation rules for Who

	if len(errors) > 0 {
		return CreateUpdateCatalogRequestMultiError(errors)
	}

	return nil
}

// CreateUpdateCatalogRequestMultiError is an error wrapping multiple
// validation errors returned by CreateUpdateCatalogRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateUpdateCatalogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUpdateCatalogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUpdateCatalogRequestMultiError) AllErrors() []error { return m }

// CreateUpdateCatalogRequestValidationError is the validation error returned
// by CreateUpdateCatalogRequest.Validate if the designated constraints aren't met.
type CreateUpdateCatalogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUpdateCatalogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUpdateCatalogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUpdateCatalogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUpdateCatalogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUpdateCatalogRequestValidationError) ErrorName() string {
	return "CreateUpdateCatalogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUpdateCatalogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUpdateCatalogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUpdateCatalogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUpdateCatalogRequestValidationError{}

// Validate checks the field values on CreateUpdateCatalogResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUpdateCatalogResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUpdateCatalogResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUpdateCatalogResponseMultiError, or nil if none found.
func (m *CreateUpdateCatalogResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUpdateCatalogResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCatalog()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateUpdateCatalogResponseValidationError{
					field:  "Catalog",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateUpdateCatalogResponseValidationError{
					field:  "Catalog",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCatalog()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateUpdateCatalogResponseValidationError{
				field:  "Catalog",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateUpdateCatalogResponseMultiError(errors)
	}

	return nil
}

// CreateUpdateCatalogResponseMultiError is an error wrapping multiple
// validation errors returned by CreateUpdateCatalogResponse.ValidateAll() if
// the designated constraints aren't met.
type CreateUpdateCatalogResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUpdateCatalogResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUpdateCatalogResponseMultiError) AllErrors() []error { return m }

// CreateUpdateCatalogResponseValidationError is the validation error returned
// by CreateUpdateCatalogResponse.Validate if the designated constraints
// aren't met.
type CreateUpdateCatalogResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUpdateCatalogResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUpdateCatalogResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUpdateCatalogResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUpdateCatalogResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUpdateCatalogResponseValidationError) ErrorName() string {
	return "CreateUpdateCatalogResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUpdateCatalogResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUpdateCatalogResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUpdateCatalogResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUpdateCatalogResponseValidationError{}

// Validate checks the field values on GetCatalogRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCatalogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCatalogRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCatalogRequestMultiError, or nil if none found.
func (m *GetCatalogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCatalogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return GetCatalogRequestMultiError(errors)
	}

	return nil
}

// GetCatalogRequestMultiError is an error wrapping multiple validation errors
// returned by GetCatalogRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCatalogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCatalogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCatalogRequestMultiError) AllErrors() []error { return m }

// GetCatalogRequestValidationError is the validation error returned by
// GetCatalogRequest.Validate if the designated constraints aren't met.
type GetCatalogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCatalogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCatalogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCatalogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCatalogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCatalogRequestValidationError) ErrorName() string {
	return "GetCatalogRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCatalogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCatalogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCatalogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCatalogRequestValidationError{}

// Validate checks the field values on GetCatalogResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCatalogResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCatalogResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCatalogResponseMultiError, or nil if none found.
func (m *GetCatalogResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCatalogResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCatalog()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetCatalogResponseValidationError{
					field:  "Catalog",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetCatalogResponseValidationError{
					field:  "Catalog",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCatalog()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCatalogResponseValidationError{
				field:  "Catalog",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetCatalogResponseMultiError(errors)
	}

	return nil
}

// GetCatalogResponseMultiError is an error wrapping multiple validation errors
// returned by GetCatalogResponse.ValidateAll() if the designated constraints
// aren't met.
type GetCatalogResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCatalogResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCatalogResponseMultiError) AllErrors() []error { return m }

// GetCatalogResponseValidationError is the validation error returned by
// GetCatalogResponse.Validate if the designated constraints aren't met.
type GetCatalogResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCatalogResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCatalogResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCatalogResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCatalogResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCatalogResponseValidationError) ErrorName() string {
	return "GetCatalogResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCatalogResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCatalogResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCatalogResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCatalogResponseValidationError{}

// Validate checks the field values on UnCatalogRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UnCatalogRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnCatalogRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnCatalogRequestMultiError, or nil if none found.
func (m *UnCatalogRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UnCatalogRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return UnCatalogRequestMultiError(errors)
	}

	return nil
}

// UnCatalogRequestMultiError is an error wrapping multiple validation errors
// returned by UnCatalogRequest.ValidateAll() if the designated constraints
// aren't met.
type UnCatalogRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnCatalogRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnCatalogRequestMultiError) AllErrors() []error { return m }

// UnCatalogRequestValidationError is the validation error returned by
// UnCatalogRequest.Validate if the designated constraints aren't met.
type UnCatalogRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnCatalogRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnCatalogRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnCatalogRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnCatalogRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnCatalogRequestValidationError) ErrorName() string { return "UnCatalogRequestValidationError" }

// Error satisfies the builtin error interface
func (e UnCatalogRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnCatalogRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnCatalogRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnCatalogRequestValidationError{}

// Validate checks the field values on UnCatalogResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UnCatalogResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnCatalogResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnCatalogResponseMultiError, or nil if none found.
func (m *UnCatalogResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UnCatalogResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	// no validation rules for Name

	if len(errors) > 0 {
		return UnCatalogResponseMultiError(errors)
	}

	return nil
}

// UnCatalogResponseMultiError is an error wrapping multiple validation errors
// returned by UnCatalogResponse.ValidateAll() if the designated constraints
// aren't met.
type UnCatalogResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnCatalogResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnCatalogResponseMultiError) AllErrors() []error { return m }

// UnCatalogResponseValidationError is the validation error returned by
// UnCatalogResponse.Validate if the designated constraints aren't met.
type UnCatalogResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnCatalogResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnCatalogResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnCatalogResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnCatalogResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnCatalogResponseValidationError) ErrorName() string {
	return "UnCatalogResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UnCatalogResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnCatalogResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnCatalogResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnCatalogResponseValidationError{}

// Validate checks the field values on ListCatalogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCatalogsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCatalogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCatalogsRequestMultiError, or nil if none found.
func (m *ListCatalogsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCatalogsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPerPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListCatalogsRequestValidationError{
					field:  "PerPage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListCatalogsRequestValidationError{
					field:  "PerPage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPerPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCatalogsRequestValidationError{
				field:  "PerPage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListCatalogsRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListCatalogsRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCatalogsRequestValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListCatalogsRequestMultiError(errors)
	}

	return nil
}

// ListCatalogsRequestMultiError is an error wrapping multiple validation
// errors returned by ListCatalogsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListCatalogsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCatalogsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCatalogsRequestMultiError) AllErrors() []error { return m }

// ListCatalogsRequestValidationError is the validation error returned by
// ListCatalogsRequest.Validate if the designated constraints aren't met.
type ListCatalogsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCatalogsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCatalogsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCatalogsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCatalogsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCatalogsRequestValidationError) ErrorName() string {
	return "ListCatalogsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCatalogsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCatalogsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCatalogsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCatalogsRequestValidationError{}

// Validate checks the field values on ListCatalogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCatalogsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCatalogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCatalogsResponseMultiError, or nil if none found.
func (m *ListCatalogsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCatalogsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCatalogs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCatalogsResponseValidationError{
						field:  fmt.Sprintf("Catalogs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCatalogsResponseValidationError{
						field:  fmt.Sprintf("Catalogs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCatalogsResponseValidationError{
					field:  fmt.Sprintf("Catalogs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalCount

	if len(errors) > 0 {
		return ListCatalogsResponseMultiError(errors)
	}

	return nil
}

// ListCatalogsResponseMultiError is an error wrapping multiple validation
// errors returned by ListCatalogsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListCatalogsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCatalogsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCatalogsResponseMultiError) AllErrors() []error { return m }

// ListCatalogsResponseValidationError is the validation error returned by
// ListCatalogsResponse.Validate if the designated constraints aren't met.
type ListCatalogsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCatalogsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCatalogsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCatalogsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCatalogsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCatalogsResponseValidationError) ErrorName() string {
	return "ListCatalogsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCatalogsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCatalogsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCatalogsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCatalogsResponseValidationError{}
