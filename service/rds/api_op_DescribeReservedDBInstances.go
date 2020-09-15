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

// Returns information about reserved DB instances for this account, or about a
// specified reserved DB instance.
func (c *Client) DescribeReservedDBInstances(ctx context.Context, params *DescribeReservedDBInstancesInput, optFns ...func(*Options)) (*DescribeReservedDBInstancesOutput, error) {
	stack := middleware.NewStack("DescribeReservedDBInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeReservedDBInstancesMiddlewares(stack)
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
	addOpDescribeReservedDBInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedDBInstances(options.Region), middleware.Before)

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
			OperationName: "DescribeReservedDBInstances",
			Err:           err,
		}
	}
	out := result.(*DescribeReservedDBInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReservedDBInstancesInput struct {
	// The product description filter value. Specify this parameter to show only those
	// reservations matching the specified product description.
	ProductDescription *string
	// The offering identifier filter value. Specify this parameter to show only
	// purchased reservations matching the specified offering identifier.
	ReservedDBInstancesOfferingId *string
	// The duration filter value, specified in years or seconds. Specify this parameter
	// to show only reservations for this duration. Valid Values: 1 | 3 | 31536000 |
	// 94608000
	Duration *string
	// The DB instance class filter value. Specify this parameter to show only those
	// reservations matching the specified DB instances class.
	DBInstanceClass *string
	// The lease identifier filter value. Specify this parameter to show only the
	// reservation that matches the specified lease ID. AWS Support might request the
	// lease ID for an issue related to a reserved DB instance.
	LeaseId *string
	// The reserved DB instance identifier filter value. Specify this parameter to show
	// only the reservation that matches the specified reservation ID.
	ReservedDBInstanceId *string
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string
	// This parameter isn't currently supported.
	Filters []*types.Filter
	// The offering type filter value. Specify this parameter to show only the
	// available offerings matching the specified offering type. Valid Values: "Partial
	// Upfront" | "All Upfront" | "No Upfront"
	OfferingType *string
	// A value that indicates whether to show only those reservations that support
	// Multi-AZ.
	MultiAZ *bool
	// The maximum number of records to include in the response. If more than the
	// MaxRecords value is available, a pagination token called a marker is included in
	// the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

// Contains the result of a successful invocation of the
// DescribeReservedDBInstances action.
type DescribeReservedDBInstancesOutput struct {
	// A list of reserved DB instances.
	ReservedDBInstances []*types.ReservedDBInstance
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeReservedDBInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReservedDBInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReservedDBInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeReservedDBInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeReservedDBInstances",
	}
}
