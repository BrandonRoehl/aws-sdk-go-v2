// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the phone number orders for the administrator's Amazon Chime account.
func (c *Client) ListPhoneNumberOrders(ctx context.Context, params *ListPhoneNumberOrdersInput, optFns ...func(*Options)) (*ListPhoneNumberOrdersOutput, error) {
	stack := middleware.NewStack("ListPhoneNumberOrders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPhoneNumberOrdersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPhoneNumberOrders(options.Region), middleware.Before)

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
			OperationName: "ListPhoneNumberOrders",
			Err:           err,
		}
	}
	out := result.(*ListPhoneNumberOrdersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPhoneNumberOrdersInput struct {
	// The maximum number of results to return in a single call.
	MaxResults *int32
	// The token to use to retrieve the next page of results.
	NextToken *string
}

type ListPhoneNumberOrdersOutput struct {
	// The token to use to retrieve the next page of results.
	NextToken *string
	// The phone number order details.
	PhoneNumberOrders []*types.PhoneNumberOrder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPhoneNumberOrdersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPhoneNumberOrders{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPhoneNumberOrders{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPhoneNumberOrders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListPhoneNumberOrders",
	}
}
