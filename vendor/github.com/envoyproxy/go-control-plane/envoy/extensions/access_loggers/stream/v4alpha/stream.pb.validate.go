// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/access_loggers/stream/v4alpha/stream.proto

package envoy_extensions_access_loggers_stream_v4alpha

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on StdoutAccessLog with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StdoutAccessLog) Validate() error {
	if m == nil {
		return nil
	}

	switch m.AccessLogFormat.(type) {

	case *StdoutAccessLog_LogFormat:

		if m.GetLogFormat() == nil {
			return StdoutAccessLogValidationError{
				field:  "LogFormat",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetLogFormat()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StdoutAccessLogValidationError{
					field:  "LogFormat",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StdoutAccessLogValidationError is the validation error returned by
// StdoutAccessLog.Validate if the designated constraints aren't met.
type StdoutAccessLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StdoutAccessLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StdoutAccessLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StdoutAccessLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StdoutAccessLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StdoutAccessLogValidationError) ErrorName() string { return "StdoutAccessLogValidationError" }

// Error satisfies the builtin error interface
func (e StdoutAccessLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStdoutAccessLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StdoutAccessLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StdoutAccessLogValidationError{}

// Validate checks the field values on StderrAccessLog with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StderrAccessLog) Validate() error {
	if m == nil {
		return nil
	}

	switch m.AccessLogFormat.(type) {

	case *StderrAccessLog_LogFormat:

		if m.GetLogFormat() == nil {
			return StderrAccessLogValidationError{
				field:  "LogFormat",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetLogFormat()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StderrAccessLogValidationError{
					field:  "LogFormat",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StderrAccessLogValidationError is the validation error returned by
// StderrAccessLog.Validate if the designated constraints aren't met.
type StderrAccessLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StderrAccessLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StderrAccessLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StderrAccessLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StderrAccessLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StderrAccessLogValidationError) ErrorName() string { return "StderrAccessLogValidationError" }

// Error satisfies the builtin error interface
func (e StderrAccessLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStderrAccessLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StderrAccessLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StderrAccessLogValidationError{}
