// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves attributes attached to the facet.
func (c *Client) ListFacetAttributes(ctx context.Context, params *ListFacetAttributesInput, optFns ...func(*Options)) (*ListFacetAttributesOutput, error) {
	stack := middleware.NewStack("ListFacetAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListFacetAttributesMiddlewares(stack)
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
	addOpListFacetAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFacetAttributes(options.Region), middleware.Before)

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
			OperationName: "ListFacetAttributes",
			Err:           err,
		}
	}
	out := result.(*ListFacetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFacetAttributesInput struct {
	// The name of the facet whose attributes will be retrieved.
	Name *string
	// The ARN of the schema where the facet resides.
	SchemaArn *string
	// The pagination token.
	NextToken *string
	// The maximum number of results to retrieve.
	MaxResults *int32
}

type ListFacetAttributesOutput struct {
	// The attributes attached to the facet.
	Attributes []*types.FacetAttribute
	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListFacetAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListFacetAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListFacetAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFacetAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListFacetAttributes",
	}
}
