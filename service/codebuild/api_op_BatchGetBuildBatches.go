// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about one or more batch builds.
func (c *Client) BatchGetBuildBatches(ctx context.Context, params *BatchGetBuildBatchesInput, optFns ...func(*Options)) (*BatchGetBuildBatchesOutput, error) {
	stack := middleware.NewStack("BatchGetBuildBatches", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetBuildBatchesMiddlewares(stack)
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
	addOpBatchGetBuildBatchesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetBuildBatches(options.Region), middleware.Before)

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
			OperationName: "BatchGetBuildBatches",
			Err:           err,
		}
	}
	out := result.(*BatchGetBuildBatchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetBuildBatchesInput struct {
	// An array that contains the batch build identifiers to retrieve.
	Ids []*string
}

type BatchGetBuildBatchesOutput struct {
	// An array that contains the identifiers of any batch builds that are not found.
	BuildBatchesNotFound []*string
	// An array of BuildBatch objects that represent the retrieved batch builds.
	BuildBatches []*types.BuildBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetBuildBatchesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetBuildBatches{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetBuildBatches{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetBuildBatches(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "BatchGetBuildBatches",
	}
}
