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

// Cancels the specified Spot Fleet requests. After you cancel a Spot Fleet
// request, the Spot Fleet launches no new Spot Instances. You must specify whether
// the Spot Fleet should also terminate its Spot Instances. If you terminate the
// instances, the Spot Fleet request enters the cancelled_terminating state.
// Otherwise, the Spot Fleet request enters the cancelled_running state and the
// instances continue to run until they are interrupted or you terminate them
// manually.
func (c *Client) CancelSpotFleetRequests(ctx context.Context, params *CancelSpotFleetRequestsInput, optFns ...func(*Options)) (*CancelSpotFleetRequestsOutput, error) {
	stack := middleware.NewStack("CancelSpotFleetRequests", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCancelSpotFleetRequestsMiddlewares(stack)
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
	addOpCancelSpotFleetRequestsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelSpotFleetRequests(options.Region), middleware.Before)

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
			OperationName: "CancelSpotFleetRequests",
			Err:           err,
		}
	}
	out := result.(*CancelSpotFleetRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CancelSpotFleetRequests.
type CancelSpotFleetRequestsInput struct {
	// The IDs of the Spot Fleet requests.
	SpotFleetRequestIds []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// Indicates whether to terminate instances for a Spot Fleet request if it is
	// canceled successfully.
	TerminateInstances *bool
}

// Contains the output of CancelSpotFleetRequests.
type CancelSpotFleetRequestsOutput struct {
	// Information about the Spot Fleet requests that are successfully canceled.
	SuccessfulFleetRequests []*types.CancelSpotFleetRequestsSuccessItem
	// Information about the Spot Fleet requests that are not successfully canceled.
	UnsuccessfulFleetRequests []*types.CancelSpotFleetRequestsErrorItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCancelSpotFleetRequestsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCancelSpotFleetRequests{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCancelSpotFleetRequests{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelSpotFleetRequests(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CancelSpotFleetRequests",
	}
}
