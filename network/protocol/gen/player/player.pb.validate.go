// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/player.proto

package pb_player

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

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UId

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on AddFriedRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddFriedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddFriedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddFriedRequestMultiError, or nil if none found.
func (m *AddFriedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddFriedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UId

	if len(errors) > 0 {
		return AddFriedRequestMultiError(errors)
	}

	return nil
}

// AddFriedRequestMultiError is an error wrapping multiple validation errors
// returned by AddFriedRequest.ValidateAll() if the designated constraints
// aren't met.
type AddFriedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddFriedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddFriedRequestMultiError) AllErrors() []error { return m }

// AddFriedRequestValidationError is the validation error returned by
// AddFriedRequest.Validate if the designated constraints aren't met.
type AddFriedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddFriedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddFriedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddFriedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddFriedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddFriedRequestValidationError) ErrorName() string { return "AddFriedRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddFriedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddFriedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddFriedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddFriedRequestValidationError{}

// Validate checks the field values on AddFriedReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddFriedReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddFriedReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddFriedReplyMultiError, or
// nil if none found.
func (m *AddFriedReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AddFriedReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Desc

	if len(errors) > 0 {
		return AddFriedReplyMultiError(errors)
	}

	return nil
}

// AddFriedReplyMultiError is an error wrapping multiple validation errors
// returned by AddFriedReply.ValidateAll() if the designated constraints
// aren't met.
type AddFriedReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddFriedReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddFriedReplyMultiError) AllErrors() []error { return m }

// AddFriedReplyValidationError is the validation error returned by
// AddFriedReply.Validate if the designated constraints aren't met.
type AddFriedReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddFriedReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddFriedReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddFriedReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddFriedReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddFriedReplyValidationError) ErrorName() string { return "AddFriedReplyValidationError" }

// Error satisfies the builtin error interface
func (e AddFriedReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddFriedReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddFriedReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddFriedReplyValidationError{}

// Validate checks the field values on DelFriedRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DelFriedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DelFriedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DelFriedRequestMultiError, or nil if none found.
func (m *DelFriedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DelFriedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UId

	if len(errors) > 0 {
		return DelFriedRequestMultiError(errors)
	}

	return nil
}

// DelFriedRequestMultiError is an error wrapping multiple validation errors
// returned by DelFriedRequest.ValidateAll() if the designated constraints
// aren't met.
type DelFriedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DelFriedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DelFriedRequestMultiError) AllErrors() []error { return m }

// DelFriedRequestValidationError is the validation error returned by
// DelFriedRequest.Validate if the designated constraints aren't met.
type DelFriedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DelFriedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DelFriedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DelFriedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DelFriedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DelFriedRequestValidationError) ErrorName() string { return "DelFriedRequestValidationError" }

// Error satisfies the builtin error interface
func (e DelFriedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDelFriedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DelFriedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DelFriedRequestValidationError{}

// Validate checks the field values on DelFriendReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DelFriendReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DelFriendReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DelFriendReplyMultiError,
// or nil if none found.
func (m *DelFriendReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DelFriendReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Desc

	if len(errors) > 0 {
		return DelFriendReplyMultiError(errors)
	}

	return nil
}

// DelFriendReplyMultiError is an error wrapping multiple validation errors
// returned by DelFriendReply.ValidateAll() if the designated constraints
// aren't met.
type DelFriendReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DelFriendReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DelFriendReplyMultiError) AllErrors() []error { return m }

// DelFriendReplyValidationError is the validation error returned by
// DelFriendReply.Validate if the designated constraints aren't met.
type DelFriendReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DelFriendReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DelFriendReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DelFriendReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DelFriendReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DelFriendReplyValidationError) ErrorName() string { return "DelFriendReplyValidationError" }

// Error satisfies the builtin error interface
func (e DelFriendReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDelFriendReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DelFriendReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DelFriendReplyValidationError{}

// Validate checks the field values on SendChatMsgRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendChatMsgRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendChatMsgRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendChatMsgRequestMultiError, or nil if none found.
func (m *SendChatMsgRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendChatMsgRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UId

	if all {
		switch v := interface{}(m.GetMsg()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SendChatMsgRequestValidationError{
					field:  "Msg",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SendChatMsgRequestValidationError{
					field:  "Msg",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMsg()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendChatMsgRequestValidationError{
				field:  "Msg",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SendChatMsgRequestMultiError(errors)
	}

	return nil
}

// SendChatMsgRequestMultiError is an error wrapping multiple validation errors
// returned by SendChatMsgRequest.ValidateAll() if the designated constraints
// aren't met.
type SendChatMsgRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendChatMsgRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendChatMsgRequestMultiError) AllErrors() []error { return m }

// SendChatMsgRequestValidationError is the validation error returned by
// SendChatMsgRequest.Validate if the designated constraints aren't met.
type SendChatMsgRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendChatMsgRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendChatMsgRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendChatMsgRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendChatMsgRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendChatMsgRequestValidationError) ErrorName() string {
	return "SendChatMsgRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SendChatMsgRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendChatMsgRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendChatMsgRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendChatMsgRequestValidationError{}

// Validate checks the field values on SendChatMsgReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SendChatMsgReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendChatMsgReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendChatMsgReplyMultiError, or nil if none found.
func (m *SendChatMsgReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SendChatMsgReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SendChatMsgReplyMultiError(errors)
	}

	return nil
}

// SendChatMsgReplyMultiError is an error wrapping multiple validation errors
// returned by SendChatMsgReply.ValidateAll() if the designated constraints
// aren't met.
type SendChatMsgReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendChatMsgReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendChatMsgReplyMultiError) AllErrors() []error { return m }

// SendChatMsgReplyValidationError is the validation error returned by
// SendChatMsgReply.Validate if the designated constraints aren't met.
type SendChatMsgReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendChatMsgReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendChatMsgReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendChatMsgReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendChatMsgReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendChatMsgReplyValidationError) ErrorName() string { return "SendChatMsgReplyValidationError" }

// Error satisfies the builtin error interface
func (e SendChatMsgReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendChatMsgReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendChatMsgReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendChatMsgReplyValidationError{}

// Validate checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChatMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChatMessageMultiError, or
// nil if none found.
func (m *ChatMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Content

	if len(errors) > 0 {
		return ChatMessageMultiError(errors)
	}

	return nil
}

// ChatMessageMultiError is an error wrapping multiple validation errors
// returned by ChatMessage.ValidateAll() if the designated constraints aren't met.
type ChatMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatMessageMultiError) AllErrors() []error { return m }

// ChatMessageValidationError is the validation error returned by
// ChatMessage.Validate if the designated constraints aren't met.
type ChatMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatMessageValidationError) ErrorName() string { return "ChatMessageValidationError" }

// Error satisfies the builtin error interface
func (e ChatMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatMessageValidationError{}
