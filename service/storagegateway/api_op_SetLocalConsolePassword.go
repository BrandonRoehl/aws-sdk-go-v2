// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the password for your VM local console. When you log in to the local
// console for the first time, you log in to the VM with the default credentials.
// We recommend that you set a new password. You don't need to know the default
// password to set a new password.
func (c *Client) SetLocalConsolePassword(ctx context.Context, params *SetLocalConsolePasswordInput, optFns ...func(*Options)) (*SetLocalConsolePasswordOutput, error) {
	stack := middleware.NewStack("SetLocalConsolePassword", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSetLocalConsolePasswordMiddlewares(stack)
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
	addOpSetLocalConsolePasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetLocalConsolePassword(options.Region), middleware.Before)

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
			OperationName: "SetLocalConsolePassword",
			Err:           err,
		}
	}
	out := result.(*SetLocalConsolePasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// SetLocalConsolePasswordInput
type SetLocalConsolePasswordInput struct {
	// The password you want to set for your VM local console.
	LocalConsolePassword *string
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
}

type SetLocalConsolePasswordOutput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSetLocalConsolePasswordMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSetLocalConsolePassword{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetLocalConsolePassword{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetLocalConsolePassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "SetLocalConsolePassword",
	}
}
