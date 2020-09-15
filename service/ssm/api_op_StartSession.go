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

// Initiates a connection to a target (for example, an instance) for a Session
// Manager session. Returns a URL and token that can be used to open a WebSocket
// connection for sending input and receiving outputs. AWS CLI usage: start-session
// is an interactive command that requires the Session Manager plugin to be
// installed on the client machine making the call. For information, see Install
// the Session Manager plugin for the AWS CLI
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html)
// in the AWS Systems Manager User Guide. AWS Tools for PowerShell usage:
// Start-SSMSession is not currently supported by AWS Tools for PowerShell on
// Windows local machines.
func (c *Client) StartSession(ctx context.Context, params *StartSessionInput, optFns ...func(*Options)) (*StartSessionOutput, error) {
	stack := middleware.NewStack("StartSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartSessionMiddlewares(stack)
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
	addOpStartSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSession(options.Region), middleware.Before)

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
			OperationName: "StartSession",
			Err:           err,
		}
	}
	out := result.(*StartSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSessionInput struct {
	// Reserved for future use.
	Parameters map[string][]*string
	// The instance to connect to for the session.
	Target *string
	// The name of the SSM document to define the parameters and plugin settings for
	// the session. For example, SSM-SessionManagerRunShell. You can call the
	// GetDocument () API to verify the document exists before attempting to start a
	// session. If no document name is provided, a shell to the instance is launched by
	// default.
	DocumentName *string
}

type StartSessionOutput struct {
	// An encrypted token value containing session and caller information. Used to
	// authenticate the connection to the instance.
	TokenValue *string
	// A URL back to SSM Agent on the instance that the Session Manager client uses to
	// send commands and receive output from the instance. Format:
	// wss://ssmmessages.region.amazonaws.com/v1/data-channel/session-id?stream=(input|output)
	// region represents the Region identifier for an AWS Region supported by AWS
	// Systems Manager, such as us-east-2 for the US East (Ohio) Region. For a list of
	// supported region values, see the Region column in Systems Manager service
	// endpoints (http://docs.aws.amazon.com/general/latest/gr/ssm.html#ssm_region) in
	// the AWS General Reference. session-id represents the ID of a Session Manager
	// session, such as 1a2b3c4dEXAMPLE.
	StreamUrl *string
	// The ID of the session.
	SessionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartSession{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "StartSession",
	}
}
