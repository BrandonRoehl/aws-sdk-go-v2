// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation lists details about a partner event source that is shared with
// your account.
func (c *Client) DescribeEventSource(ctx context.Context, params *DescribeEventSourceInput, optFns ...func(*Options)) (*DescribeEventSourceOutput, error) {
	stack := middleware.NewStack("DescribeEventSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEventSourceMiddlewares(stack)
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
	addOpDescribeEventSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventSource(options.Region), middleware.Before)

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
			OperationName: "DescribeEventSource",
			Err:           err,
		}
	}
	out := result.(*DescribeEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventSourceInput struct {
	// The name of the partner event source to display the details of.
	Name *string
}

type DescribeEventSourceOutput struct {
	// The name of the SaaS partner that created the event source.
	CreatedBy *string
	// The name of the partner event source.
	Name *string
	// The date and time that the event source will expire if you do not create a
	// matching event bus.
	ExpirationTime *time.Time
	// The ARN of the partner event source.
	Arn *string
	// The state of the event source. If it is ACTIVE, you have already created a
	// matching event bus for this event source, and that event bus is active. If it is
	// PENDING, either you haven't yet created a matching event bus, or that event bus
	// is deactivated. If it is DELETED, you have created a matching event bus, but the
	// event source has since been deleted.
	State types.EventSourceState
	// The date and time that the event source was created.
	CreationTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEventSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEventSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeEventSource",
	}
}
