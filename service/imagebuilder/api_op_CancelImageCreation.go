// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// CancelImageCreation cancels the creation of Image. This operation can only be
// used on images in a non-terminal state.
func (c *Client) CancelImageCreation(ctx context.Context, params *CancelImageCreationInput, optFns ...func(*Options)) (*CancelImageCreationOutput, error) {
	stack := middleware.NewStack("CancelImageCreation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCancelImageCreationMiddlewares(stack)
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
	addIdempotencyToken_opCancelImageCreationMiddleware(stack, options)
	addOpCancelImageCreationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelImageCreation(options.Region), middleware.Before)

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
			OperationName: "CancelImageCreation",
			Err:           err,
		}
	}
	out := result.(*CancelImageCreationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelImageCreationInput struct {
	// The Amazon Resource Name (ARN) of the image whose creation you want to cancel.
	ImageBuildVersionArn *string
	// The idempotency token used to make this request idempotent.
	ClientToken *string
}

type CancelImageCreationOutput struct {
	// The request ID that uniquely identifies this request.
	RequestId *string
	// The idempotency token used to make this request idempotent.
	ClientToken *string
	// The Amazon Resource Name (ARN) of the image whose creation has been cancelled.
	ImageBuildVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCancelImageCreationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCancelImageCreation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelImageCreation{}, middleware.After)
}

type idempotencyToken_initializeOpCancelImageCreation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCancelImageCreation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCancelImageCreation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CancelImageCreationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CancelImageCreationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCancelImageCreationMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCancelImageCreation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCancelImageCreation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "CancelImageCreation",
	}
}
