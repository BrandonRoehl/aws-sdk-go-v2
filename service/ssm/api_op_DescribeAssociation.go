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

// Describes the association for the specified target or instance. If you created
// the association by using the Targets parameter, then you must retrieve the
// association by using the association ID. If you created the association by
// specifying an instance ID and a Systems Manager document, then you retrieve the
// association by specifying the document name and the instance ID.
func (c *Client) DescribeAssociation(ctx context.Context, params *DescribeAssociationInput, optFns ...func(*Options)) (*DescribeAssociationOutput, error) {
	stack := middleware.NewStack("DescribeAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAssociationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssociation(options.Region), middleware.Before)

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
			OperationName: "DescribeAssociation",
			Err:           err,
		}
	}
	out := result.(*DescribeAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssociationInput struct {
	// The association ID for which you want information.
	AssociationId *string
	// Specify the association version to retrieve. To view the latest version, either
	// specify $LATEST for this parameter, or omit this parameter. To view a list of
	// all associations for an instance, use ListAssociations (). To get a list of
	// versions for a specific association, use ListAssociationVersions ().
	AssociationVersion *string
	// The name of the Systems Manager document.
	Name *string
	// The instance ID.
	InstanceId *string
}

type DescribeAssociationOutput struct {
	// Information about the association.
	AssociationDescription *types.AssociationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeAssociation",
	}
}
