// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a pipeline. A pipeline consumes messages from a channel and allows you
// to process the messages before storing them in a data store. You must specify
// both a channel and a datastore activity and, optionally, as many as 23
// additional activities in the pipelineActivities array.
func (c *Client) CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) {
	stack := middleware.NewStack("CreatePipeline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreatePipelineMiddlewares(stack)
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
	addOpCreatePipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePipeline(options.Region), middleware.Before)

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
			OperationName: "CreatePipeline",
			Err:           err,
		}
	}
	out := result.(*CreatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePipelineInput struct {
	// Metadata which can be used to manage the pipeline.
	Tags []*types.Tag
	// A list of "PipelineActivity" objects. Activities perform transformations on your
	// messages, such as removing, renaming or adding message attributes; filtering
	// messages based on attribute values; invoking your Lambda functions on messages
	// for advanced processing; or performing mathematical transformations to normalize
	// device data. The list can be 2-25 PipelineActivity objects and must contain both
	// a channel and a datastore activity. Each entry in the list must contain only one
	// activity, for example: pipelineActivities = [ { "channel": { ... } }, {
	// "lambda": { ... } }, ... ]
	PipelineActivities []*types.PipelineActivity
	// The name of the pipeline.
	PipelineName *string
}

type CreatePipelineOutput struct {
	// The name of the pipeline.
	PipelineName *string
	// The ARN of the pipeline.
	PipelineArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreatePipelineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreatePipeline{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePipeline{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "CreatePipeline",
	}
}
