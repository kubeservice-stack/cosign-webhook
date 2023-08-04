// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: xds/type/matcher/v3/regex.proto

package v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on RegexMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RegexMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetRegex()) < 1 {
		return RegexMatcherValidationError{
			field:  "Regex",
			reason: "value length must be at least 1 runes",
		}
	}

	switch m.EngineType.(type) {

	case *RegexMatcher_GoogleRe2:

		if m.GetGoogleRe2() == nil {
			return RegexMatcherValidationError{
				field:  "GoogleRe2",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetGoogleRe2()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegexMatcherValidationError{
					field:  "GoogleRe2",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return RegexMatcherValidationError{
			field:  "EngineType",
			reason: "value is required",
		}

	}

	return nil
}

// RegexMatcherValidationError is the validation error returned by
// RegexMatcher.Validate if the designated constraints aren't met.
type RegexMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegexMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegexMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegexMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegexMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegexMatcherValidationError) ErrorName() string { return "RegexMatcherValidationError" }

// Error satisfies the builtin error interface
func (e RegexMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegexMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegexMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegexMatcherValidationError{}

// Validate checks the field values on RegexMatcher_GoogleRE2 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegexMatcher_GoogleRE2) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RegexMatcher_GoogleRE2ValidationError is the validation error returned by
// RegexMatcher_GoogleRE2.Validate if the designated constraints aren't met.
type RegexMatcher_GoogleRE2ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegexMatcher_GoogleRE2ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegexMatcher_GoogleRE2ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegexMatcher_GoogleRE2ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegexMatcher_GoogleRE2ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegexMatcher_GoogleRE2ValidationError) ErrorName() string {
	return "RegexMatcher_GoogleRE2ValidationError"
}

// Error satisfies the builtin error interface
func (e RegexMatcher_GoogleRE2ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegexMatcher_GoogleRE2.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegexMatcher_GoogleRE2ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegexMatcher_GoogleRE2ValidationError{}
