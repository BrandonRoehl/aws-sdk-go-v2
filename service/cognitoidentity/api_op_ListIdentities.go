// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the identities in an identity pool. You must use AWS Developer credentials
// to call this API.
func (c *Client) ListIdentities(ctx context.Context, params *ListIdentitiesInput, optFns ...func(*Options)) (*ListIdentitiesOutput, error) {
	stack := middleware.NewStack("ListIdentities", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListIdentitiesMiddlewares(stack)
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
	addOpListIdentitiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentities(options.Region), middleware.Before)

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
			OperationName: "ListIdentities",
			Err:           err,
		}
	}
	out := result.(*ListIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the ListIdentities action.
type ListIdentitiesInput struct {
	// A pagination token.
	NextToken *string
	// The maximum number of identities to return.
	MaxResults *int32
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string
	// An optional boolean parameter that allows you to hide disabled identities. If
	// omitted, the ListIdentities API will include disabled identities in the
	// response.
	HideDisabled *bool
}

// The response to a ListIdentities request.
type ListIdentitiesOutput struct {
	// An object containing a set of identities and associated mappings.
	Identities []*types.IdentityDescription
	// A pagination token.
	NextToken *string
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListIdentitiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListIdentities{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListIdentities{}, middleware.After)
}

func newServiceMetadataMiddleware_opListIdentities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "ListIdentities",
	}
}
