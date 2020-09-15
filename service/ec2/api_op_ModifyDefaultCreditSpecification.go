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

// Modifies the default credit option for CPU usage of burstable performance
// instances. The default credit option is set at the account level per AWS Region,
// and is specified per instance family. All new burstable performance instances in
// the account launch using the default credit option.
// ModifyDefaultCreditSpecification is an asynchronous operation, which works at an
// AWS Region level and modifies the credit option for each Availability Zone. All
// zones in a Region are updated within five minutes. But if instances are launched
// during this operation, they might not get the new credit option until the zone
// is updated. To verify whether the update has occurred, you can call
// GetDefaultCreditSpecification and check DefaultCreditSpecification for updates.
// For more information, see Burstable performance instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ModifyDefaultCreditSpecification(ctx context.Context, params *ModifyDefaultCreditSpecificationInput, optFns ...func(*Options)) (*ModifyDefaultCreditSpecificationOutput, error) {
	stack := middleware.NewStack("ModifyDefaultCreditSpecification", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpModifyDefaultCreditSpecificationMiddlewares(stack)
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
	addOpModifyDefaultCreditSpecificationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDefaultCreditSpecification(options.Region), middleware.Before)

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
			OperationName: "ModifyDefaultCreditSpecification",
			Err:           err,
		}
	}
	out := result.(*ModifyDefaultCreditSpecificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDefaultCreditSpecificationInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The credit option for CPU usage of the instance family. Valid Values: standard |
	// unlimited
	CpuCredits *string
	// The instance family.
	InstanceFamily types.UnlimitedSupportedInstanceFamily
}

type ModifyDefaultCreditSpecificationOutput struct {
	// The default credit option for CPU usage of the instance family.
	InstanceFamilyCreditSpecification *types.InstanceFamilyCreditSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpModifyDefaultCreditSpecificationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpModifyDefaultCreditSpecification{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpModifyDefaultCreditSpecification{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDefaultCreditSpecification(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyDefaultCreditSpecification",
	}
}
