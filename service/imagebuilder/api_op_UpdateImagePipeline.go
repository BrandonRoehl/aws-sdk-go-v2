// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a new image pipeline. Image pipelines enable you to automate the
// creation and distribution of images.
func (c *Client) UpdateImagePipeline(ctx context.Context, params *UpdateImagePipelineInput, optFns ...func(*Options)) (*UpdateImagePipelineOutput, error) {
	stack := middleware.NewStack("UpdateImagePipeline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateImagePipelineMiddlewares(stack)
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
	addIdempotencyToken_opUpdateImagePipelineMiddleware(stack, options)
	addOpUpdateImagePipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateImagePipeline(options.Region), middleware.Before)

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
			OperationName: "UpdateImagePipeline",
			Err:           err,
		}
	}
	out := result.(*UpdateImagePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateImagePipelineInput struct {
	// The description of the image pipeline.
	Description *string
	// The idempotency token used to make this request idempotent.
	ClientToken *string
	// The status of the image pipeline.
	Status types.PipelineStatus
	// The image test configuration of the image pipeline.
	ImageTestsConfiguration *types.ImageTestsConfiguration
	// Collects additional information about the image being created, including the
	// operating system (OS) version and package list. This information is used to
	// enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled *bool
	// The Amazon Resource Name (ARN) of the distribution configuration that will be
	// used to configure and distribute images updated by this image pipeline.
	DistributionConfigurationArn *string
	// The Amazon Resource Name (ARN) of the infrastructure configuration that will be
	// used to build images updated by this image pipeline.
	InfrastructureConfigurationArn *string
	// The Amazon Resource Name (ARN) of the image pipeline that you want to update.
	ImagePipelineArn *string
	// The schedule of the image pipeline.
	Schedule *types.Schedule
	// The Amazon Resource Name (ARN) of the image recipe that will be used to
	// configure images updated by this image pipeline.
	ImageRecipeArn *string
}

type UpdateImagePipelineOutput struct {
	// The idempotency token used to make this request idempotent.
	ClientToken *string
	// The request ID that uniquely identifies this request.
	RequestId *string
	// The Amazon Resource Name (ARN) of the image pipeline that was updated by this
	// request.
	ImagePipelineArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateImagePipelineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateImagePipeline{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateImagePipeline{}, middleware.After)
}

type idempotencyToken_initializeOpUpdateImagePipeline struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateImagePipeline) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateImagePipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateImagePipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateImagePipelineInput ")
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
func addIdempotencyToken_opUpdateImagePipelineMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateImagePipeline{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateImagePipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "UpdateImagePipeline",
	}
}
