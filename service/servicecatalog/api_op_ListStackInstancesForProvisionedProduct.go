// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns summary information about stack instances that are associated with the
// specified CFN_STACKSET type provisioned product. You can filter for stack
// instances that are associated with a specific AWS account name or region.
func (c *Client) ListStackInstancesForProvisionedProduct(ctx context.Context, params *ListStackInstancesForProvisionedProductInput, optFns ...func(*Options)) (*ListStackInstancesForProvisionedProductOutput, error) {
	stack := middleware.NewStack("ListStackInstancesForProvisionedProduct", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListStackInstancesForProvisionedProductMiddlewares(stack)
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
	addOpListStackInstancesForProvisionedProductValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListStackInstancesForProvisionedProduct(options.Region), middleware.Before)

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
			OperationName: "ListStackInstancesForProvisionedProduct",
			Err:           err,
		}
	}
	out := result.(*ListStackInstancesForProvisionedProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStackInstancesForProvisionedProductInput struct {
	// The identifier of the provisioned product.
	ProvisionedProductId *string
	// The maximum number of items to return with this call.
	PageSize *int32
	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
}

type ListStackInstancesForProvisionedProductOutput struct {
	// List of stack instances.
	StackInstances []*types.StackInstance
	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListStackInstancesForProvisionedProductMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListStackInstancesForProvisionedProduct{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStackInstancesForProvisionedProduct{}, middleware.After)
}

func newServiceMetadataMiddleware_opListStackInstancesForProvisionedProduct(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListStackInstancesForProvisionedProduct",
	}
}
