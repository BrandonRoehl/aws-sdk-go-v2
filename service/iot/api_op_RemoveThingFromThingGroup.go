// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Remove the specified thing from the specified group. You must specify either a
// thingGroupArn or a thingGroupName to identify the thing group and either a
// thingArn or a thingName to identify the thing to remove from the thing group.
func (c *Client) RemoveThingFromThingGroup(ctx context.Context, params *RemoveThingFromThingGroupInput, optFns ...func(*Options)) (*RemoveThingFromThingGroupOutput, error) {
	stack := middleware.NewStack("RemoveThingFromThingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRemoveThingFromThingGroupMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveThingFromThingGroup(options.Region), middleware.Before)

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
			OperationName: "RemoveThingFromThingGroup",
			Err:           err,
		}
	}
	out := result.(*RemoveThingFromThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveThingFromThingGroupInput struct {
	// The ARN of the thing to remove from the group.
	ThingArn *string
	// The name of the thing to remove from the group.
	ThingName *string
	// The group name.
	ThingGroupName *string
	// The group ARN.
	ThingGroupArn *string
}

type RemoveThingFromThingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRemoveThingFromThingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRemoveThingFromThingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveThingFromThingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveThingFromThingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RemoveThingFromThingGroup",
	}
}
