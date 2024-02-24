// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kratos_layout_notice.proto

package noticepb

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

// Validate checks the field values on Notice with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Notice) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Notice with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NoticeMultiError, or nil if none found.
func (m *Notice) ValidateAll() error {
	return m.validate(true)
}

func (m *Notice) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Desc

	// no validation rules for Unit

	// no validation rules for Content

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return NoticeMultiError(errors)
	}

	return nil
}

// NoticeMultiError is an error wrapping multiple validation errors returned by
// Notice.ValidateAll() if the designated constraints aren't met.
type NoticeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoticeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoticeMultiError) AllErrors() []error { return m }

// NoticeValidationError is the validation error returned by Notice.Validate if
// the designated constraints aren't met.
type NoticeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoticeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoticeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoticeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoticeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoticeValidationError) ErrorName() string { return "NoticeValidationError" }

// Error satisfies the builtin error interface
func (e NoticeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotice.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoticeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoticeValidationError{}

// Validate checks the field values on PageNoticeUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageNoticeUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageNoticeUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageNoticeUserRequestMultiError, or nil if none found.
func (m *PageNoticeUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageNoticeUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageNoticeUserRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageNoticeUserRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 50]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetNoticeId() <= 0 {
		err := PageNoticeUserRequestValidationError{
			field:  "NoticeId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PageNoticeUserRequestMultiError(errors)
	}

	return nil
}

// PageNoticeUserRequestMultiError is an error wrapping multiple validation
// errors returned by PageNoticeUserRequest.ValidateAll() if the designated
// constraints aren't met.
type PageNoticeUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageNoticeUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageNoticeUserRequestMultiError) AllErrors() []error { return m }

// PageNoticeUserRequestValidationError is the validation error returned by
// PageNoticeUserRequest.Validate if the designated constraints aren't met.
type PageNoticeUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageNoticeUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageNoticeUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageNoticeUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageNoticeUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageNoticeUserRequestValidationError) ErrorName() string {
	return "PageNoticeUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PageNoticeUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageNoticeUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageNoticeUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageNoticeUserRequestValidationError{}

// Validate checks the field values on PageNoticeUserReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageNoticeUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageNoticeUserReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageNoticeUserReplyMultiError, or nil if none found.
func (m *PageNoticeUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageNoticeUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PageNoticeUserReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageNoticeUserReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageNoticeUserReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageNoticeUserReplyMultiError(errors)
	}

	return nil
}

// PageNoticeUserReplyMultiError is an error wrapping multiple validation
// errors returned by PageNoticeUserReply.ValidateAll() if the designated
// constraints aren't met.
type PageNoticeUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageNoticeUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageNoticeUserReplyMultiError) AllErrors() []error { return m }

// PageNoticeUserReplyValidationError is the validation error returned by
// PageNoticeUserReply.Validate if the designated constraints aren't met.
type PageNoticeUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageNoticeUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageNoticeUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageNoticeUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageNoticeUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageNoticeUserReplyValidationError) ErrorName() string {
	return "PageNoticeUserReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PageNoticeUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageNoticeUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageNoticeUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageNoticeUserReplyValidationError{}

// Validate checks the field values on PageNoticeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PageNoticeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageNoticeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageNoticeRequestMultiError, or nil if none found.
func (m *PageNoticeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageNoticeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageNoticeRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageNoticeRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 50]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Title != nil {
		// no validation rules for Title
	}

	if len(errors) > 0 {
		return PageNoticeRequestMultiError(errors)
	}

	return nil
}

// PageNoticeRequestMultiError is an error wrapping multiple validation errors
// returned by PageNoticeRequest.ValidateAll() if the designated constraints
// aren't met.
type PageNoticeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageNoticeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageNoticeRequestMultiError) AllErrors() []error { return m }

// PageNoticeRequestValidationError is the validation error returned by
// PageNoticeRequest.Validate if the designated constraints aren't met.
type PageNoticeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageNoticeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageNoticeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageNoticeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageNoticeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageNoticeRequestValidationError) ErrorName() string {
	return "PageNoticeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PageNoticeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageNoticeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageNoticeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageNoticeRequestValidationError{}

// Validate checks the field values on PageNoticeReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PageNoticeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageNoticeReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageNoticeReplyMultiError, or nil if none found.
func (m *PageNoticeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageNoticeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PageNoticeReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageNoticeReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageNoticeReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageNoticeReplyMultiError(errors)
	}

	return nil
}

// PageNoticeReplyMultiError is an error wrapping multiple validation errors
// returned by PageNoticeReply.ValidateAll() if the designated constraints
// aren't met.
type PageNoticeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageNoticeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageNoticeReplyMultiError) AllErrors() []error { return m }

// PageNoticeReplyValidationError is the validation error returned by
// PageNoticeReply.Validate if the designated constraints aren't met.
type PageNoticeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageNoticeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageNoticeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageNoticeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageNoticeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageNoticeReplyValidationError) ErrorName() string { return "PageNoticeReplyValidationError" }

// Error satisfies the builtin error interface
func (e PageNoticeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageNoticeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageNoticeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageNoticeReplyValidationError{}

// Validate checks the field values on AddNoticeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddNoticeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddNoticeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddNoticeRequestMultiError, or nil if none found.
func (m *AddNoticeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddNoticeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := AddNoticeRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDesc()) < 1 {
		err := AddNoticeRequestValidationError{
			field:  "Desc",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetUnit()) < 1 {
		err := AddNoticeRequestValidationError{
			field:  "Unit",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetContent()) < 1 {
		err := AddNoticeRequestValidationError{
			field:  "Content",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddNoticeRequestMultiError(errors)
	}

	return nil
}

// AddNoticeRequestMultiError is an error wrapping multiple validation errors
// returned by AddNoticeRequest.ValidateAll() if the designated constraints
// aren't met.
type AddNoticeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddNoticeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddNoticeRequestMultiError) AllErrors() []error { return m }

// AddNoticeRequestValidationError is the validation error returned by
// AddNoticeRequest.Validate if the designated constraints aren't met.
type AddNoticeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddNoticeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddNoticeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddNoticeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddNoticeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddNoticeRequestValidationError) ErrorName() string { return "AddNoticeRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddNoticeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddNoticeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddNoticeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddNoticeRequestValidationError{}

// Validate checks the field values on AddNoticeReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddNoticeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddNoticeReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddNoticeReplyMultiError,
// or nil if none found.
func (m *AddNoticeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AddNoticeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddNoticeReplyMultiError(errors)
	}

	return nil
}

// AddNoticeReplyMultiError is an error wrapping multiple validation errors
// returned by AddNoticeReply.ValidateAll() if the designated constraints
// aren't met.
type AddNoticeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddNoticeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddNoticeReplyMultiError) AllErrors() []error { return m }

// AddNoticeReplyValidationError is the validation error returned by
// AddNoticeReply.Validate if the designated constraints aren't met.
type AddNoticeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddNoticeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddNoticeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddNoticeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddNoticeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddNoticeReplyValidationError) ErrorName() string { return "AddNoticeReplyValidationError" }

// Error satisfies the builtin error interface
func (e AddNoticeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddNoticeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddNoticeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddNoticeReplyValidationError{}

// Validate checks the field values on GetNoticeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetNoticeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoticeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNoticeRequestMultiError, or nil if none found.
func (m *GetNoticeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoticeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetNoticeRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetNoticeRequestMultiError(errors)
	}

	return nil
}

// GetNoticeRequestMultiError is an error wrapping multiple validation errors
// returned by GetNoticeRequest.ValidateAll() if the designated constraints
// aren't met.
type GetNoticeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoticeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoticeRequestMultiError) AllErrors() []error { return m }

// GetNoticeRequestValidationError is the validation error returned by
// GetNoticeRequest.Validate if the designated constraints aren't met.
type GetNoticeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoticeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoticeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoticeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoticeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoticeRequestValidationError) ErrorName() string { return "GetNoticeRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetNoticeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoticeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoticeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoticeRequestValidationError{}

// Validate checks the field values on UpdateNoticeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateNoticeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoticeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNoticeRequestMultiError, or nil if none found.
func (m *UpdateNoticeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoticeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateNoticeRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := UpdateNoticeRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDesc()) < 1 {
		err := UpdateNoticeRequestValidationError{
			field:  "Desc",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetUnit()) < 1 {
		err := UpdateNoticeRequestValidationError{
			field:  "Unit",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetContent()) < 1 {
		err := UpdateNoticeRequestValidationError{
			field:  "Content",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateNoticeRequestMultiError(errors)
	}

	return nil
}

// UpdateNoticeRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateNoticeRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateNoticeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoticeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoticeRequestMultiError) AllErrors() []error { return m }

// UpdateNoticeRequestValidationError is the validation error returned by
// UpdateNoticeRequest.Validate if the designated constraints aren't met.
type UpdateNoticeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoticeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoticeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoticeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoticeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoticeRequestValidationError) ErrorName() string {
	return "UpdateNoticeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNoticeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoticeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoticeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoticeRequestValidationError{}

// Validate checks the field values on DeleteNoticeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteNoticeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNoticeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteNoticeRequestMultiError, or nil if none found.
func (m *DeleteNoticeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNoticeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteNoticeRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteNoticeRequestMultiError(errors)
	}

	return nil
}

// DeleteNoticeRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteNoticeRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteNoticeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNoticeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNoticeRequestMultiError) AllErrors() []error { return m }

// DeleteNoticeRequestValidationError is the validation error returned by
// DeleteNoticeRequest.Validate if the designated constraints aren't met.
type DeleteNoticeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNoticeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNoticeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNoticeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNoticeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNoticeRequestValidationError) ErrorName() string {
	return "DeleteNoticeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNoticeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNoticeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNoticeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNoticeRequestValidationError{}