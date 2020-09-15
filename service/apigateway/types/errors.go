// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// The submitted request is not valid, for example, the input is incomplete or
// incorrect. See the accompanying error message for details.
type BadRequestException struct {
	Message *string
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *BadRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *BadRequestException) HasMessage() bool {
	return e.Message != nil
}

// The request configuration has conflicts. For details, see the accompanying error
// message.
type ConflictException struct {
	Message *string
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ConflictException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ConflictException) HasMessage() bool {
	return e.Message != nil
}

// The request exceeded the rate limit. Retry after the specified time period.
type LimitExceededException struct {
	Message *string

	RetryAfterSeconds *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}
func (e *LimitExceededException) GetRetryAfterSeconds() string {
	return ptr.ToString(e.RetryAfterSeconds)
}
func (e *LimitExceededException) HasRetryAfterSeconds() bool {
	return e.RetryAfterSeconds != nil
}

// The requested resource is not found. Make sure that the request URI is correct.
type NotFoundException struct {
	Message *string
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The requested service is not available. For details see the accompanying error
// message. Retry after the specified time period.
type ServiceUnavailableException struct {
	Message *string

	RetryAfterSeconds *string
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string             { return "ServiceUnavailableException" }
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ServiceUnavailableException) GetRetryAfterSeconds() string {
	return ptr.ToString(e.RetryAfterSeconds)
}
func (e *ServiceUnavailableException) HasRetryAfterSeconds() bool {
	return e.RetryAfterSeconds != nil
}
func (e *ServiceUnavailableException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServiceUnavailableException) HasMessage() bool {
	return e.Message != nil
}

// The request has reached its throttling limit. Retry after the specified time
// period.
type TooManyRequestsException struct {
	Message *string

	RetryAfterSeconds *string
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string             { return "TooManyRequestsException" }
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *TooManyRequestsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *TooManyRequestsException) HasMessage() bool {
	return e.Message != nil
}
func (e *TooManyRequestsException) GetRetryAfterSeconds() string {
	return ptr.ToString(e.RetryAfterSeconds)
}
func (e *TooManyRequestsException) HasRetryAfterSeconds() bool {
	return e.RetryAfterSeconds != nil
}

// The request is denied because the caller has insufficient permissions.
type UnauthorizedException struct {
	Message *string
}

func (e *UnauthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedException) ErrorCode() string             { return "UnauthorizedException" }
func (e *UnauthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *UnauthorizedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UnauthorizedException) HasMessage() bool {
	return e.Message != nil
}
