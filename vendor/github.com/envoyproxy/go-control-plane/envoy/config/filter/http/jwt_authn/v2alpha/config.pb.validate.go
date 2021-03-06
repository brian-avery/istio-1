// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/jwt_authn/v2alpha/config.proto

package envoy_config_filter_http_jwt_authn_v2alpha

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

// define the regex for a UUID once up-front
var _config_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on JwtProvider with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *JwtProvider) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetIssuer()) < 1 {
		return JwtProviderValidationError{
			field:  "Issuer",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for Forward

	for idx, item := range m.GetFromHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtProviderValidationError{
					field:  fmt.Sprintf("FromHeaders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ForwardPayloadHeader

	// no validation rules for PayloadInMetadata

	switch m.JwksSourceSpecifier.(type) {

	case *JwtProvider_RemoteJwks:

		if v, ok := interface{}(m.GetRemoteJwks()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtProviderValidationError{
					field:  "RemoteJwks",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *JwtProvider_LocalJwks:

		if v, ok := interface{}(m.GetLocalJwks()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtProviderValidationError{
					field:  "LocalJwks",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return JwtProviderValidationError{
			field:  "JwksSourceSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// JwtProviderValidationError is the validation error returned by
// JwtProvider.Validate if the designated constraints aren't met.
type JwtProviderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtProviderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtProviderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtProviderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtProviderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtProviderValidationError) ErrorName() string { return "JwtProviderValidationError" }

// Error satisfies the builtin error interface
func (e JwtProviderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtProvider.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtProviderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtProviderValidationError{}

// Validate checks the field values on RemoteJwks with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RemoteJwks) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttpUri()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoteJwksValidationError{
				field:  "HttpUri",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCacheDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoteJwksValidationError{
				field:  "CacheDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RemoteJwksValidationError is the validation error returned by
// RemoteJwks.Validate if the designated constraints aren't met.
type RemoteJwksValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoteJwksValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoteJwksValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoteJwksValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoteJwksValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoteJwksValidationError) ErrorName() string { return "RemoteJwksValidationError" }

// Error satisfies the builtin error interface
func (e RemoteJwksValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoteJwks.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoteJwksValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoteJwksValidationError{}

// Validate checks the field values on JwtHeader with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JwtHeader) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return JwtHeaderValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for ValuePrefix

	return nil
}

// JwtHeaderValidationError is the validation error returned by
// JwtHeader.Validate if the designated constraints aren't met.
type JwtHeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtHeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtHeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtHeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtHeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtHeaderValidationError) ErrorName() string { return "JwtHeaderValidationError" }

// Error satisfies the builtin error interface
func (e JwtHeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtHeader.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtHeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtHeaderValidationError{}

// Validate checks the field values on ProviderWithAudiences with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProviderWithAudiences) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProviderName

	return nil
}

// ProviderWithAudiencesValidationError is the validation error returned by
// ProviderWithAudiences.Validate if the designated constraints aren't met.
type ProviderWithAudiencesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProviderWithAudiencesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProviderWithAudiencesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProviderWithAudiencesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProviderWithAudiencesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProviderWithAudiencesValidationError) ErrorName() string {
	return "ProviderWithAudiencesValidationError"
}

// Error satisfies the builtin error interface
func (e ProviderWithAudiencesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProviderWithAudiences.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProviderWithAudiencesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProviderWithAudiencesValidationError{}

// Validate checks the field values on JwtRequirement with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *JwtRequirement) Validate() error {
	if m == nil {
		return nil
	}

	switch m.RequiresType.(type) {

	case *JwtRequirement_ProviderName:
		// no validation rules for ProviderName

	case *JwtRequirement_ProviderAndAudiences:

		if v, ok := interface{}(m.GetProviderAndAudiences()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtRequirementValidationError{
					field:  "ProviderAndAudiences",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *JwtRequirement_RequiresAny:

		if v, ok := interface{}(m.GetRequiresAny()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtRequirementValidationError{
					field:  "RequiresAny",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *JwtRequirement_RequiresAll:

		if v, ok := interface{}(m.GetRequiresAll()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtRequirementValidationError{
					field:  "RequiresAll",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *JwtRequirement_AllowMissingOrFailed:

		if v, ok := interface{}(m.GetAllowMissingOrFailed()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtRequirementValidationError{
					field:  "AllowMissingOrFailed",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// JwtRequirementValidationError is the validation error returned by
// JwtRequirement.Validate if the designated constraints aren't met.
type JwtRequirementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtRequirementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtRequirementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtRequirementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtRequirementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtRequirementValidationError) ErrorName() string { return "JwtRequirementValidationError" }

// Error satisfies the builtin error interface
func (e JwtRequirementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtRequirement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtRequirementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtRequirementValidationError{}

// Validate checks the field values on JwtRequirementOrList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *JwtRequirementOrList) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRequirements()) < 2 {
		return JwtRequirementOrListValidationError{
			field:  "Requirements",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetRequirements() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtRequirementOrListValidationError{
					field:  fmt.Sprintf("Requirements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// JwtRequirementOrListValidationError is the validation error returned by
// JwtRequirementOrList.Validate if the designated constraints aren't met.
type JwtRequirementOrListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtRequirementOrListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtRequirementOrListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtRequirementOrListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtRequirementOrListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtRequirementOrListValidationError) ErrorName() string {
	return "JwtRequirementOrListValidationError"
}

// Error satisfies the builtin error interface
func (e JwtRequirementOrListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtRequirementOrList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtRequirementOrListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtRequirementOrListValidationError{}

// Validate checks the field values on JwtRequirementAndList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *JwtRequirementAndList) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRequirements()) < 2 {
		return JwtRequirementAndListValidationError{
			field:  "Requirements",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetRequirements() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtRequirementAndListValidationError{
					field:  fmt.Sprintf("Requirements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// JwtRequirementAndListValidationError is the validation error returned by
// JwtRequirementAndList.Validate if the designated constraints aren't met.
type JwtRequirementAndListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtRequirementAndListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtRequirementAndListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtRequirementAndListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtRequirementAndListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtRequirementAndListValidationError) ErrorName() string {
	return "JwtRequirementAndListValidationError"
}

// Error satisfies the builtin error interface
func (e JwtRequirementAndListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtRequirementAndList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtRequirementAndListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtRequirementAndListValidationError{}

// Validate checks the field values on RequirementRule with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RequirementRule) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMatch() == nil {
		return RequirementRuleValidationError{
			field:  "Match",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequirementRuleValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRequires()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequirementRuleValidationError{
				field:  "Requires",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RequirementRuleValidationError is the validation error returned by
// RequirementRule.Validate if the designated constraints aren't met.
type RequirementRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequirementRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequirementRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequirementRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequirementRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequirementRuleValidationError) ErrorName() string { return "RequirementRuleValidationError" }

// Error satisfies the builtin error interface
func (e RequirementRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequirementRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequirementRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequirementRuleValidationError{}

// Validate checks the field values on FilterStateRule with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FilterStateRule) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return FilterStateRuleValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for Requires

	return nil
}

// FilterStateRuleValidationError is the validation error returned by
// FilterStateRule.Validate if the designated constraints aren't met.
type FilterStateRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterStateRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterStateRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterStateRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterStateRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterStateRuleValidationError) ErrorName() string { return "FilterStateRuleValidationError" }

// Error satisfies the builtin error interface
func (e FilterStateRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterStateRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterStateRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterStateRuleValidationError{}

// Validate checks the field values on JwtAuthentication with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *JwtAuthentication) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Providers

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JwtAuthenticationValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetFilterStateRules()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return JwtAuthenticationValidationError{
				field:  "FilterStateRules",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// JwtAuthenticationValidationError is the validation error returned by
// JwtAuthentication.Validate if the designated constraints aren't met.
type JwtAuthenticationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtAuthenticationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtAuthenticationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtAuthenticationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtAuthenticationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtAuthenticationValidationError) ErrorName() string {
	return "JwtAuthenticationValidationError"
}

// Error satisfies the builtin error interface
func (e JwtAuthenticationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtAuthentication.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtAuthenticationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtAuthenticationValidationError{}
