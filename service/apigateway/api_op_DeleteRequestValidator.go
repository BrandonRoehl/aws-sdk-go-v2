// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a RequestValidator () of a given RestApi ().
func (c *Client) DeleteRequestValidator(ctx context.Context, params *DeleteRequestValidatorInput, optFns ...func(*Options)) (*DeleteRequestValidatorOutput, error) {
	stack := middleware.NewStack("DeleteRequestValidator", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteRequestValidatorMiddlewares(stack)
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
	addOpDeleteRequestValidatorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRequestValidator(options.Region), middleware.Before)

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
			OperationName: "DeleteRequestValidator",
			Err:           err,
		}
	}
	out := result.(*DeleteRequestValidatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Deletes a specified RequestValidator () of a given RestApi ().
type DeleteRequestValidatorInput struct {
	TemplateSkipList []*string
	Template         *bool
	// [Required] The string identifier of the associated RestApi ().
	RestApiId *string
	Title     *string
	Name      *string
	// [Required] The identifier of the RequestValidator () to be deleted.
	RequestValidatorId *string
}

type DeleteRequestValidatorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteRequestValidatorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRequestValidator{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRequestValidator{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteRequestValidator(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteRequestValidator",
	}
}
