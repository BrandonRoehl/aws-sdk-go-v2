// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// (Enterprise edition only) Creates a new namespace for you to use with Amazon
// QuickSight. A namespace allows you to isolate the QuickSight users and groups
// that are registered for that namespace. Users that access the namespace can
// share assets only with other users or groups in the same namespace. They can't
// see users and groups in other namespaces. You can create a namespace after your
// AWS account is subscribed to QuickSight. The namespace must be unique within the
// AWS account. By default, there is a limit of 100 namespaces per AWS account. To
// increase your limit, create a ticket with AWS Support.
func (c *Client) CreateNamespace(ctx context.Context, params *CreateNamespaceInput, optFns ...func(*Options)) (*CreateNamespaceOutput, error) {
	stack := middleware.NewStack("CreateNamespace", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateNamespaceMiddlewares(stack)
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
	addOpCreateNamespaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNamespace(options.Region), middleware.Before)

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
			OperationName: "CreateNamespace",
			Err:           err,
		}
	}
	out := result.(*CreateNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNamespaceInput struct {
	// The ID for the AWS account that you want to create the QuickSight namespace in.
	AwsAccountId *string
	// The tags that you want to associate with the namespace that you're creating.
	Tags []*types.Tag
	// The name that you want to use to describe the new namespace.
	Namespace *string
	// Specifies the type of your user identity directory. Currently, this supports
	// users with an identity type of QUICKSIGHT.
	IdentityStore types.IdentityStore
}

type CreateNamespaceOutput struct {
	// The status of the creation of the namespace. This is an asynchronous process. A
	// status of CREATED means that your namespace is ready to use. If an error occurs,
	// it indicates if the process is retryable or non-retryable. In the case of a
	// non-retryable error, refer to the error message for follow-up actions.
	CreationStatus types.NamespaceStatus
	// The name of the new namespace that you created.
	Name *string
	// Specifies the type of your user identity directory. Currently, this supports
	// users with an identity type of QUICKSIGHT.
	IdentityStore types.IdentityStore
	// The AWS request ID for this operation.
	RequestId *string
	// The AWS Region that you want to use for the free SPICE capacity for the new
	// namespace. This is set to the region that you run CreateNamespace in.
	CapacityRegion *string
	// The ARN of the QuickSight namespace you created.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateNamespaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateNamespace{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNamespace{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateNamespace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateNamespace",
	}
}
