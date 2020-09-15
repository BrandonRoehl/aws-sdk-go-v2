// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches phone numbers that can be ordered.
func (c *Client) SearchAvailablePhoneNumbers(ctx context.Context, params *SearchAvailablePhoneNumbersInput, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) {
	stack := middleware.NewStack("SearchAvailablePhoneNumbers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSearchAvailablePhoneNumbersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchAvailablePhoneNumbers(options.Region), middleware.Before)

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
			OperationName: "SearchAvailablePhoneNumbers",
			Err:           err,
		}
	}
	out := result.(*SearchAvailablePhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchAvailablePhoneNumbersInput struct {
	// The token to use to retrieve the next page of results.
	NextToken *string
	// The toll-free prefix that you use to filter results.
	TollFreePrefix *string
	// The country used to filter results.
	Country *string
	// The city used to filter results.
	City *string
	// The state used to filter results.
	State *string
	// The maximum number of results to return in a single call.
	MaxResults *int32
	// The area code used to filter results.
	AreaCode *string
}

type SearchAvailablePhoneNumbersOutput struct {
	// List of phone numbers, in E.164 format.
	E164PhoneNumbers []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSearchAvailablePhoneNumbersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSearchAvailablePhoneNumbers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchAvailablePhoneNumbers{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchAvailablePhoneNumbers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "SearchAvailablePhoneNumbers",
	}
}
