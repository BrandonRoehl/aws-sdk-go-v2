// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the target On-Demand and target Spot capacities for the instance fleet
// with the specified InstanceFleetID within the cluster specified using ClusterID.
// The call either succeeds or fails atomically. The instance fleet configuration
// is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x
// versions.
func (c *Client) ModifyInstanceFleet(ctx context.Context, params *ModifyInstanceFleetInput, optFns ...func(*Options)) (*ModifyInstanceFleetOutput, error) {
	stack := middleware.NewStack("ModifyInstanceFleet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpModifyInstanceFleetMiddlewares(stack)
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
	addOpModifyInstanceFleetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceFleet(options.Region), middleware.Before)

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
			OperationName: "ModifyInstanceFleet",
			Err:           err,
		}
	}
	out := result.(*ModifyInstanceFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceFleetInput struct {
	// The unique identifier of the cluster.
	ClusterId *string
	// The unique identifier of the instance fleet.
	InstanceFleet *types.InstanceFleetModifyConfig
}

type ModifyInstanceFleetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpModifyInstanceFleetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpModifyInstanceFleet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyInstanceFleet{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyInstanceFleet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ModifyInstanceFleet",
	}
}
