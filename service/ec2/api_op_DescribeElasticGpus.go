// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the Elastic Graphics accelerator associated with your instances. For
// more information about Elastic Graphics, see Amazon Elastic Graphics
// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html).
func (c *Client) DescribeElasticGpus(ctx context.Context, params *DescribeElasticGpusInput, optFns ...func(*Options)) (*DescribeElasticGpusOutput, error) {
	stack := middleware.NewStack("DescribeElasticGpus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeElasticGpusMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticGpus(options.Region), middleware.Before)

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
			OperationName: "DescribeElasticGpus",
			Err:           err,
		}
	}
	out := result.(*DescribeElasticGpusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeElasticGpusInput struct {
	// The Elastic Graphics accelerator IDs.
	ElasticGpuIds []*string
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000.
	MaxResults *int32
	// The filters.
	//
	//     * availability-zone - The Availability Zone in which the
	// Elastic Graphics accelerator resides.
	//
	//     * elastic-gpu-health - The status of
	// the Elastic Graphics accelerator (OK | IMPAIRED).
	//
	//     * elastic-gpu-state - The
	// state of the Elastic Graphics accelerator (ATTACHED).
	//
	//     * elastic-gpu-type -
	// The type of Elastic Graphics accelerator; for example, eg1.medium.
	//
	//     *
	// instance-id - The ID of the instance to which the Elastic Graphics accelerator
	// is associated.
	Filters []*types.Filter
	// The token to request the next page of results.
	NextToken *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeElasticGpusOutput struct {
	// The total number of items to return. If the total number of items available is
	// more than the value specified in max-items then a Next-Token will be provided in
	// the output that you can use to resume pagination.
	MaxResults *int32
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// Information about the Elastic Graphics accelerators.
	ElasticGpuSet []*types.ElasticGpus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeElasticGpusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeElasticGpus{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeElasticGpus{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeElasticGpus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeElasticGpus",
	}
}
