// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Purchases a reserved DB instance offering.
func (c *Client) PurchaseReservedDBInstancesOffering(ctx context.Context, params *PurchaseReservedDBInstancesOfferingInput, optFns ...func(*Options)) (*PurchaseReservedDBInstancesOfferingOutput, error) {
	stack := middleware.NewStack("PurchaseReservedDBInstancesOffering", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPurchaseReservedDBInstancesOfferingMiddlewares(stack)
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
	addOpPurchaseReservedDBInstancesOfferingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedDBInstancesOffering(options.Region), middleware.Before)

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
			OperationName: "PurchaseReservedDBInstancesOffering",
			Err:           err,
		}
	}
	out := result.(*PurchaseReservedDBInstancesOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type PurchaseReservedDBInstancesOfferingInput struct {
	// Customer-specified identifier to track this reservation. Example:
	// myreservationID
	ReservedDBInstanceId *string
	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []*types.Tag
	// The number of instances to reserve. Default: 1
	DBInstanceCount *int32
	// The ID of the Reserved DB instance offering to purchase. Example:
	// 438012d3-4052-4cc7-b2e3-8d3372e0e706
	ReservedDBInstancesOfferingId *string
}

type PurchaseReservedDBInstancesOfferingOutput struct {
	// This data type is used as a response element in the DescribeReservedDBInstances
	// and PurchaseReservedDBInstancesOffering actions.
	ReservedDBInstance *types.ReservedDBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPurchaseReservedDBInstancesOfferingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPurchaseReservedDBInstancesOffering{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPurchaseReservedDBInstancesOffering{}, middleware.After)
}

func newServiceMetadataMiddleware_opPurchaseReservedDBInstancesOffering(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "PurchaseReservedDBInstancesOffering",
	}
}
