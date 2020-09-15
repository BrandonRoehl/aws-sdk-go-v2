// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your instances, including information about the
// operating system platform, the version of SSM Agent installed on the instance,
// instance status, and so on. If you specify one or more instance IDs, it returns
// information for those instances. If you do not specify instance IDs, it returns
// information for all your instances. If you specify an instance ID that is not
// valid or an instance that you do not own, you receive an error. The IamRole
// field for this API action is the Amazon Identity and Access Management (IAM)
// role assigned to on-premises instances. This call does not return the IAM role
// for EC2 instances.
func (c *Client) DescribeInstanceInformation(ctx context.Context, params *DescribeInstanceInformationInput, optFns ...func(*Options)) (*DescribeInstanceInformationOutput, error) {
	stack := middleware.NewStack("DescribeInstanceInformation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeInstanceInformationMiddlewares(stack)
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
	addOpDescribeInstanceInformationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceInformation(options.Region), middleware.Before)

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
			OperationName: "DescribeInstanceInformation",
			Err:           err,
		}
	}
	out := result.(*DescribeInstanceInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceInformationInput struct {
	// This is a legacy method. We recommend that you don't use this method. Instead,
	// use the Filters data type. Filters enables you to return instance information by
	// filtering based on tags applied to managed instances. Attempting to use
	// InstanceInformationFilterList and Filters leads to an exception error.
	InstanceInformationFilterList []*types.InstanceInformationFilter
	// One or more filters. Use a filter to return a more specific list of instances.
	// You can filter based on tags applied to EC2 instances. Use this Filters data
	// type instead of InstanceInformationFilterList, which is deprecated.
	Filters []*types.InstanceInformationStringFilter
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
}

type DescribeInstanceInformationOutput struct {
	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string
	// The instance information list.
	InstanceInformationList []*types.InstanceInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeInstanceInformationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstanceInformation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstanceInformation{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInstanceInformation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeInstanceInformation",
	}
}
