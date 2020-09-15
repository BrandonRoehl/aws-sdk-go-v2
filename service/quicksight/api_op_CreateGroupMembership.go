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

// Adds an Amazon QuickSight user to an Amazon QuickSight group.
func (c *Client) CreateGroupMembership(ctx context.Context, params *CreateGroupMembershipInput, optFns ...func(*Options)) (*CreateGroupMembershipOutput, error) {
	stack := middleware.NewStack("CreateGroupMembership", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateGroupMembershipMiddlewares(stack)
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
	addOpCreateGroupMembershipValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGroupMembership(options.Region), middleware.Before)

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
			OperationName: "CreateGroupMembership",
			Err:           err,
		}
	}
	out := result.(*CreateGroupMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGroupMembershipInput struct {
	// The name of the group that you want to add the user to.
	GroupName *string
	// The name of the user that you want to add to the group membership.
	MemberName *string
	// The namespace. Currently, you should set this to default.
	Namespace *string
	// The ID for the AWS account that the group is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string
}

type CreateGroupMembershipOutput struct {
	// The AWS request ID for this operation.
	RequestId *string
	// The group member.
	GroupMember *types.GroupMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateGroupMembershipMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateGroupMembership{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGroupMembership{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGroupMembership(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateGroupMembership",
	}
}
