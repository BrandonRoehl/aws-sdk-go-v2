// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associate the specified TagOption with the specified portfolio or product.
func (c *Client) AssociateTagOptionWithResource(ctx context.Context, params *AssociateTagOptionWithResourceInput, optFns ...func(*Options)) (*AssociateTagOptionWithResourceOutput, error) {
	stack := middleware.NewStack("AssociateTagOptionWithResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateTagOptionWithResourceMiddlewares(stack)
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
	addOpAssociateTagOptionWithResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTagOptionWithResource(options.Region), middleware.Before)

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
			OperationName: "AssociateTagOptionWithResource",
			Err:           err,
		}
	}
	out := result.(*AssociateTagOptionWithResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTagOptionWithResourceInput struct {
	// The TagOption identifier.
	TagOptionId *string
	// The resource identifier.
	ResourceId *string
}

type AssociateTagOptionWithResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateTagOptionWithResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateTagOptionWithResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateTagOptionWithResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateTagOptionWithResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "AssociateTagOptionWithResource",
	}
}
