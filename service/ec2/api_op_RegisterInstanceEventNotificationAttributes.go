// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a set of tag keys to include in scheduled event notifications for your
// resources. To remove tags, use .
func (c *Client) RegisterInstanceEventNotificationAttributes(ctx context.Context, params *RegisterInstanceEventNotificationAttributesInput, optFns ...func(*Options)) (*RegisterInstanceEventNotificationAttributesOutput, error) {
	stack := middleware.NewStack("RegisterInstanceEventNotificationAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpRegisterInstanceEventNotificationAttributesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterInstanceEventNotificationAttributes(options.Region), middleware.Before)

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
			OperationName: "RegisterInstanceEventNotificationAttributes",
			Err:           err,
		}
	}
	out := result.(*RegisterInstanceEventNotificationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterInstanceEventNotificationAttributesInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// Information about the tag keys to register.
	InstanceTagAttribute *types.RegisterInstanceTagAttributeRequest
}

type RegisterInstanceEventNotificationAttributesOutput struct {
	// The resulting set of tag keys.
	InstanceTagAttribute *types.InstanceTagNotificationAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpRegisterInstanceEventNotificationAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpRegisterInstanceEventNotificationAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpRegisterInstanceEventNotificationAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterInstanceEventNotificationAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RegisterInstanceEventNotificationAttributes",
	}
}
