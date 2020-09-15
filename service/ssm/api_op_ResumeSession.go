// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reconnects a session to an instance after it has been disconnected. Connections
// can be resumed for disconnected sessions, but not terminated sessions. This
// command is primarily for use by client machines to automatically reconnect
// during intermittent network issues. It is not intended for any other use.
func (c *Client) ResumeSession(ctx context.Context, params *ResumeSessionInput, optFns ...func(*Options)) (*ResumeSessionOutput, error) {
	stack := middleware.NewStack("ResumeSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResumeSessionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpResumeSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResumeSession(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ResumeSession",
			Err:           err,
		}
	}
	out := result.(*ResumeSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeSessionInput struct {
	// The ID of the disconnected session to resume.
	SessionId *string
}

type ResumeSessionOutput struct {
	// The ID of the session.
	SessionId *string
	// A URL back to SSM Agent on the instance that the Session Manager client uses to
	// send commands and receive output from the instance. Format:
	// wss://ssmmessages.region.amazonaws.com/v1/data-channel/session-id?stream=(input|output).
	// region represents the Region identifier for an AWS Region supported by AWS
	// Systems Manager, such as us-east-2 for the US East (Ohio) Region. For a list of
	// supported region values, see the Region column in Systems Manager service
	// endpoints (http://docs.aws.amazon.com/general/latest/gr/ssm.html#ssm_region) in
	// the AWS General Reference. session-id represents the ID of a Session Manager
	// session, such as 1a2b3c4dEXAMPLE.
	StreamUrl *string
	// An encrypted token value containing session and caller information. Used to
	// authenticate the connection to the instance.
	TokenValue *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResumeSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResumeSession{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResumeSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opResumeSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ResumeSession",
	}
}
