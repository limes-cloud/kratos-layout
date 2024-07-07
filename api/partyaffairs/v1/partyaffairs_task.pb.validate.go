// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/partyaffairs/partyaffairs_task.proto

package v1

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

// Validate checks the field values on Task with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Task) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Task with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TaskMultiError, or nil if none found.
func (m *Task) ValidateAll() error {
	return m.validate(true)
}

func (m *Task) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Desc

	// no validation rules for IsUpdate

	// no validation rules for Start

	// no validation rules for End

	// no validation rules for Config

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return TaskMultiError(errors)
	}

	return nil
}

// TaskMultiError is an error wrapping multiple validation errors returned by
// Task.ValidateAll() if the designated constraints aren't met.
type TaskMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TaskMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TaskMultiError) AllErrors() []error { return m }

// TaskValidationError is the validation error returned by Task.Validate if the
// designated constraints aren't met.
type TaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskValidationError) ErrorName() string { return "TaskValidationError" }

// Error satisfies the builtin error interface
func (e TaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskValidationError{}

// Validate checks the field values on GetTaskRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTaskRequestMultiError,
// or nil if none found.
func (m *GetTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetTaskRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetTaskRequestMultiError(errors)
	}

	return nil
}

// GetTaskRequestMultiError is an error wrapping multiple validation errors
// returned by GetTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskRequestMultiError) AllErrors() []error { return m }

// GetTaskRequestValidationError is the validation error returned by
// GetTaskRequest.Validate if the designated constraints aren't met.
type GetTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskRequestValidationError) ErrorName() string { return "GetTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskRequestValidationError{}

// Validate checks the field values on PageTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PageTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageTaskRequestMultiError, or nil if none found.
func (m *PageTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageTaskRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageTaskRequestValidationError{
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

	if m.NotFinish != nil {
		// no validation rules for NotFinish
	}

	if len(errors) > 0 {
		return PageTaskRequestMultiError(errors)
	}

	return nil
}

// PageTaskRequestMultiError is an error wrapping multiple validation errors
// returned by PageTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type PageTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageTaskRequestMultiError) AllErrors() []error { return m }

// PageTaskRequestValidationError is the validation error returned by
// PageTaskRequest.Validate if the designated constraints aren't met.
type PageTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageTaskRequestValidationError) ErrorName() string { return "PageTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e PageTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageTaskRequestValidationError{}

// Validate checks the field values on PageTaskReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageTaskReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageTaskReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageTaskReplyMultiError, or
// nil if none found.
func (m *PageTaskReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageTaskReply) validate(all bool) error {
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
					errors = append(errors, PageTaskReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageTaskReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageTaskReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageTaskReplyMultiError(errors)
	}

	return nil
}

// PageTaskReplyMultiError is an error wrapping multiple validation errors
// returned by PageTaskReply.ValidateAll() if the designated constraints
// aren't met.
type PageTaskReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageTaskReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageTaskReplyMultiError) AllErrors() []error { return m }

// PageTaskReplyValidationError is the validation error returned by
// PageTaskReply.Validate if the designated constraints aren't met.
type PageTaskReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageTaskReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageTaskReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageTaskReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageTaskReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageTaskReplyValidationError) ErrorName() string { return "PageTaskReplyValidationError" }

// Error satisfies the builtin error interface
func (e PageTaskReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageTaskReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageTaskReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageTaskReplyValidationError{}

// Validate checks the field values on AddTaskRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddTaskRequestMultiError,
// or nil if none found.
func (m *AddTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := AddTaskRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDesc()) < 1 {
		err := AddTaskRequestValidationError{
			field:  "Desc",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStart() <= 1 {
		err := AddTaskRequestValidationError{
			field:  "Start",
			reason: "value must be greater than 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEnd() <= 1 {
		err := AddTaskRequestValidationError{
			field:  "End",
			reason: "value must be greater than 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetConfig()) < 1 {
		err := AddTaskRequestValidationError{
			field:  "Config",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.IsUpdate != nil {
		// no validation rules for IsUpdate
	}

	if len(errors) > 0 {
		return AddTaskRequestMultiError(errors)
	}

	return nil
}

// AddTaskRequestMultiError is an error wrapping multiple validation errors
// returned by AddTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type AddTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTaskRequestMultiError) AllErrors() []error { return m }

// AddTaskRequestValidationError is the validation error returned by
// AddTaskRequest.Validate if the designated constraints aren't met.
type AddTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTaskRequestValidationError) ErrorName() string { return "AddTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTaskRequestValidationError{}

// Validate checks the field values on UpdateTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTaskRequestMultiError, or nil if none found.
func (m *UpdateTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := UpdateTaskRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDesc()) < 1 {
		err := UpdateTaskRequestValidationError{
			field:  "Desc",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStart() <= 1 {
		err := UpdateTaskRequestValidationError{
			field:  "Start",
			reason: "value must be greater than 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEnd() <= 1 {
		err := UpdateTaskRequestValidationError{
			field:  "End",
			reason: "value must be greater than 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetConfig()) < 1 {
		err := UpdateTaskRequestValidationError{
			field:  "Config",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetId() <= 0 {
		err := UpdateTaskRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.IsUpdate != nil {
		// no validation rules for IsUpdate
	}

	if len(errors) > 0 {
		return UpdateTaskRequestMultiError(errors)
	}

	return nil
}

// UpdateTaskRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTaskRequestMultiError) AllErrors() []error { return m }

// UpdateTaskRequestValidationError is the validation error returned by
// UpdateTaskRequest.Validate if the designated constraints aren't met.
type UpdateTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTaskRequestValidationError) ErrorName() string {
	return "UpdateTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTaskRequestValidationError{}

// Validate checks the field values on DeleteTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTaskRequestMultiError, or nil if none found.
func (m *DeleteTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteTaskRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteTaskRequestMultiError(errors)
	}

	return nil
}

// DeleteTaskRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTaskRequestMultiError) AllErrors() []error { return m }

// DeleteTaskRequestValidationError is the validation error returned by
// DeleteTaskRequest.Validate if the designated constraints aren't met.
type DeleteTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskRequestValidationError) ErrorName() string {
	return "DeleteTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskRequestValidationError{}
