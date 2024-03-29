// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: app/gen_code/api/data.proto

package api

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

// Validate checks the field values on Node with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Node) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Node with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NodeMultiError, or nil if none found.
func (m *Node) ValidateAll() error {
	return m.validate(true)
}

func (m *Node) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Type

	// no validation rules for Props

	// no validation rules for DisplayName

	// no validation rules for Hidden

	// no validation rules for PagePath

	if len(errors) > 0 {
		return NodeMultiError(errors)
	}

	return nil
}

// NodeMultiError is an error wrapping multiple validation errors returned by
// Node.ValidateAll() if the designated constraints aren't met.
type NodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeMultiError) AllErrors() []error { return m }

// NodeValidationError is the validation error returned by Node.Validate if the
// designated constraints aren't met.
type NodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeValidationError) ErrorName() string { return "NodeValidationError" }

// Error satisfies the builtin error interface
func (e NodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeValidationError{}

// Validate checks the field values on Page with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Page) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Page with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PageMultiError, or nil if none found.
func (m *Page) ValidateAll() error {
	return m.validate(true)
}

func (m *Page) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	// no validation rules for Name

	if len(errors) > 0 {
		return PageMultiError(errors)
	}

	return nil
}

// PageMultiError is an error wrapping multiple validation errors returned by
// Page.ValidateAll() if the designated constraints aren't met.
type PageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageMultiError) AllErrors() []error { return m }

// PageValidationError is the validation error returned by Page.Validate if the
// designated constraints aren't met.
type PageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageValidationError) ErrorName() string { return "PageValidationError" }

// Error satisfies the builtin error interface
func (e PageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageValidationError{}
