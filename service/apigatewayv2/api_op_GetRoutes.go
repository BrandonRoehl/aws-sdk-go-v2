// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the Routes for an API.
func (c *Client) GetRoutes(ctx context.Context, params *GetRoutesInput, optFns ...func(*Options)) (*GetRoutesOutput, error) {
	stack := middleware.NewStack("GetRoutes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetRoutesMiddlewares(stack)
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
	addOpGetRoutesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRoutes(options.Region), middleware.Before)

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
			OperationName: "GetRoutes",
			Err:           err,
		}
	}
	out := result.(*GetRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRoutesInput struct {
	// The API identifier.
	ApiId *string
	// The maximum number of elements to be returned for this resource.
	MaxResults *string
	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string
}

type GetRoutesOutput struct {
	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string
	// The elements from this collection.
	Items []*types.Route

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetRoutesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetRoutes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRoutes{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRoutes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetRoutes",
	}
}
