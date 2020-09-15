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
	"time"
)

// Describes the events for the specified Spot Fleet request during the specified
// time. Spot Fleet events are delayed by up to 30 seconds before they can be
// described. This ensures that you can query by the last evaluated time and not
// miss a recorded event. Spot Fleet events are available for 48 hours.
func (c *Client) DescribeSpotFleetRequestHistory(ctx context.Context, params *DescribeSpotFleetRequestHistoryInput, optFns ...func(*Options)) (*DescribeSpotFleetRequestHistoryOutput, error) {
	stack := middleware.NewStack("DescribeSpotFleetRequestHistory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeSpotFleetRequestHistoryMiddlewares(stack)
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
	addOpDescribeSpotFleetRequestHistoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotFleetRequestHistory(options.Region), middleware.Before)

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
			OperationName: "DescribeSpotFleetRequestHistory",
			Err:           err,
		}
	}
	out := result.(*DescribeSpotFleetRequestHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotFleetRequestHistory.
type DescribeSpotFleetRequestHistoryInput struct {
	// The type of events to describe. By default, all events are described.
	EventType types.EventType
	// The ID of the Spot Fleet request.
	SpotFleetRequestId *string
	// The token for the next set of results.
	NextToken *string
	// The starting date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int32
}

// Contains the output of DescribeSpotFleetRequestHistory.
type DescribeSpotFleetRequestHistoryOutput struct {
	// Information about the events in the history of the Spot Fleet request.
	HistoryRecords []*types.HistoryRecord
	// The starting date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time
	// The ID of the Spot Fleet request.
	SpotFleetRequestId *string
	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// The last date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ). All records up to this time were retrieved. If nextToken
	// indicates that there are more results, this value is not present.
	LastEvaluatedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeSpotFleetRequestHistoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotFleetRequestHistory{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotFleetRequestHistory{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSpotFleetRequestHistory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotFleetRequestHistory",
	}
}
