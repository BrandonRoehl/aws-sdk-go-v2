// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the detailed parameter list for a particular cache parameter group.
func (c *Client) DescribeCacheParameters(ctx context.Context, params *DescribeCacheParametersInput, optFns ...func(*Options)) (*DescribeCacheParametersOutput, error) {
	stack := middleware.NewStack("DescribeCacheParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeCacheParametersMiddlewares(stack)
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
	addOpDescribeCacheParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheParameters(options.Region), middleware.Before)

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
			OperationName: "DescribeCacheParameters",
			Err:           err,
		}
	}
	out := result.(*DescribeCacheParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheParameters operation.
type DescribeCacheParametersInput struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved.  <p>Default: 100</p>
	// <p>Constraints: minimum 20; maximum 100.</p>
	MaxRecords *int32
	// The parameter types to return. Valid values: user | system | engine-default
	Source *string
	// The name of a specific cache parameter group to return details for.
	CacheParameterGroupName *string
	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string
}

// Represents the output of a DescribeCacheParameters operation.
type DescribeCacheParametersOutput struct {
	// A list of Parameter () instances.
	Parameters []*types.Parameter
	// Provides an identifier to allow retrieval of paginated results.
	Marker *string
	// A list of parameters specific to a particular cache node type. Each element in
	// the list contains detailed information about one parameter.
	CacheNodeTypeSpecificParameters []*types.CacheNodeTypeSpecificParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeCacheParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCacheParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheParameters",
	}
}
