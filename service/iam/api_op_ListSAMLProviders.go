// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the SAML provider resource objects defined in IAM in the account. This
// operation requires Signature Version 4
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
func (c *Client) ListSAMLProviders(ctx context.Context, params *ListSAMLProvidersInput, optFns ...func(*Options)) (*ListSAMLProvidersOutput, error) {
	stack := middleware.NewStack("ListSAMLProviders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListSAMLProvidersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSAMLProviders(options.Region), middleware.Before)

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
			OperationName: "ListSAMLProviders",
			Err:           err,
		}
	}
	out := result.(*ListSAMLProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSAMLProvidersInput struct {
}

// Contains the response to a successful ListSAMLProviders () request.
type ListSAMLProvidersOutput struct {
	// The list of SAML provider resource objects defined in IAM for this AWS account.
	SAMLProviderList []*types.SAMLProviderListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListSAMLProvidersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListSAMLProviders{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListSAMLProviders{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSAMLProviders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListSAMLProviders",
	}
}
